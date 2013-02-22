package w3u

import "fmt"
import "bytes"
import "io/ioutil"
import "os"

import "github.com/mewkiz/pkg/errorsutil"

import "github.com/karlek/sensou/unit"
import "github.com/karlek/sensou/field"
import "github.com/karlek/sensou/temp"

// The entire file
type W3u struct {
	Header []byte
	Units  []unit.Unit
}

// The (only?) valid w3u header
var Header = []byte{02, 00, 00, 00, 00, 00, 00, 00}

// The bytes that seperate field attribute and value
var ValueSeperator = []byte{00, 00, 00}
var FieldSeperator = []byte{00, 00, 00, 00, 00}

func New() *W3u {
	return &W3u{Header: Header}
}

// Adds a unit to the unit array in the W3u struct
// It also increases number of unit count which is used when writing to file
func (w3 *W3u) AddUnit(u *unit.Unit) {
	w3.Units = append(w3.Units, *u)
}

// Validate the w3u header with a valid header
func (w3 *W3u) IsHeader() (isHeader bool) {
	return bytes.Equal(w3.Header, Header)
}

// Prints all units and their fields
func (w3 *W3u) PrintUnits() {
	for _, u := range w3.Units {
		for _, f := range u.Fields {
			fmt.Printf("\t%s = %v\n", f.GetDescriber(), f.Value)
		}
	}
}

// Write W3u struct to file
func (w3 *W3u) Write(fileName string) (err error) {

	buffer := new(bytes.Buffer)

	buffer.Write(w3.Header)
	buffer.Write([]byte{byte(len(w3.Units))})
	buffer.Write(ValueSeperator)
	for _, u := range w3.Units {
		buffer.Write([]byte(u.BaseRawCode))
		buffer.Write([]byte(u.CustomRawCode))
		buffer.Write([]byte{byte(len(u.Fields))})
		buffer.Write(ValueSeperator)
		for _, f := range u.Fields {
			buffer.Write([]byte(f.Name))
			buffer.Write([]byte{byte(f.FieldType())})
			buffer.Write(ValueSeperator)
			buffer.Write(f.ValueToBuf())
			buffer.Write(FieldSeperator)
		}
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(buffer.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func Open(fileName string) (w3 *W3u, err error) {

	w3 = New()

	// Read w3u file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	// Read all bytes from the file pointer
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Make the bytes into a Buffer
	buffer := bytes.NewBuffer(buf)

	// Check header
	w3.Header = buffer.Next(8)
	if !w3.IsHeader() {
		return nil, errorsutil.Errorf("Invalid w3u header! Correct header: % X, your header: % X", Header, w3.Header)
	}

	// Number of units
	numUnits := uint(buffer.Next(1)[0])
	if numUnits < 0 {
		return nil, errorsutil.Errorf("Invalid number of units: %d\n", numUnits)
	}

	// Skip seperator
	buffer.Next(3)

	for i := 0; i < int(numUnits); i++ {

		// Base unit rawcode
		baseUnit, err := GetRawCode(buffer)
		if err != nil {
			return nil, err
		}

		// Custom unit rawcode
		customUnit, err := GetRawCode(buffer)
		if err != nil {
			return nil, err
		}

		u, err := unit.NewUnit(baseUnit, customUnit)
		if err != nil {
			return nil, err
		}

		// Number of fields
		numFields := uint(buffer.Next(1)[0])
		if numFields < 0 {
			return nil, errorsutil.Errorf("Invalid number of fields: %d\n", numFields)
		}

		// Skip seperator
		buffer.Next(3)

		for j := 0; j < int(numFields); j++ {

			f := new(field.Field)

			// Field name in raw data
			fieldName, err := GetRawCode(buffer)
			if err != nil {
				return nil, err
			}
			f.Name = fieldName

			// Type of field, 0:int(and bool), 1:fixed decimal point, 2:float, 3:string
			fieldType := int(buffer.Next(1)[0])

			// 0, 1, 2, 3 are valid field types
			if fieldType < 0 || fieldType > 3 {
				return nil, errorsutil.Errorf("Invalid field type: %d\n", fieldType)
			}

			// Skip seperator
			buffer.Next(3)

			// ###Add fixed
			switch fieldType {
			case 0:
				f.Value, err = temp.FromBuf(buffer.Next(3))
				if err != nil {
					return nil, err
				}
				// f.Value = temp.BigToLittle(buffer.Next(3))

			case 2:
				// Float value
				f.Value = buffer.Next(3)

			case 3:
				// Get next seperator, that's were the string ends
				strLen := bytes.Index(buffer.Bytes(), []byte{00, 00, 00})

				// String value
				f.Value = string(buffer.Next(strLen))
			}

			// Skip field seperator
			buffer.Next(5)
			u.Fields = append(u.Fields, *f)
		}
		w3.Units = append(w3.Units, *u)
	}

	return w3, nil
}

func GetRawCode(buffer *bytes.Buffer) (rawCode string, err error) {
	rawCode = string(buffer.Next(4))
	if !temp.IsRawCode(rawCode) {
		return "", errorsutil.Errorf("Illegal raw code: [% X]", rawCode)
	}

	return rawCode, nil
}
