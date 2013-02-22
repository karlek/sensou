package unit

import "github.com/karlek/sensou/field"
import "github.com/karlek/sensou/temp"
import "github.com/mewkiz/pkg/errorsutil"

// A single warcraft 3 unit
type Unit struct {
	BaseRawCode   string
	CustomRawCode string
	Fields        []field.Field
}

// Adds a new field to the unit
func (unit *Unit) AddField(fieldRawCode string, value interface{}) (err error) {
	if temp.IsRawCode(fieldRawCode) == false {
		return errorsutil.Errorf("Invalid field rawcode: %s - must be upper or lowercase alpha numeric and 4 characters long", fieldRawCode)
	}

	unit.Fields = append(unit.Fields, field.Field{Name: fieldRawCode, Value: value})
	return nil
}

// Returns a unit struct initiated with the raw codes
func NewUnit(baseRawCode, customRawCode string) (unit *Unit, err error) {
	if temp.IsRawCode(baseRawCode) == false {
		return nil, errorsutil.Errorf("Invalid base unit rawcode: %s - must be upper or lowercase alpha numeric and 4 characters long", baseRawCode)
	} else if temp.IsRawCode(customRawCode) == false {
		return nil, errorsutil.Errorf("Invalid custom unit rawcode: %s - must be upper or lowercase alpha numeric and 4 characters long", customRawCode)
	}

	return &Unit{BaseRawCode: baseRawCode, CustomRawCode: customRawCode}, nil
}
