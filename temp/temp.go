package temp

import "bytes"
import "encoding/binary"
import "github.com/mewkiz/pkg/errorsutil"
import "unicode"
import "strconv"

type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(i.Get(), 2)
}

func FromString(s string) (b Binary, err error) {
	uInt, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		return 0, err
	}

	return Binary(uInt), nil
}

func FromBuf(byteNum []byte) (b Binary, err error) {

	// buf := bytes.NewBuffer([]byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40})

	if len(byteNum)%8 != 1 {
		for {
			if len(byteNum)%8 == 0 {
				break
			}
			byteNum = append(byteNum, 0x00)
		}
	}

	buf := bytes.NewBuffer(byteNum)

	err = binary.Read(buf, binary.LittleEndian, &b)
	if err != nil {
		return 0, errorsutil.Errorf("FromBuf: %s - % x", err, byteNum)
	}

	return b, nil
}

func (i Binary) Get() uint64 {
	return uint64(i)
}

func (i Binary) Little() (little []byte) {

	buf := new(bytes.Buffer)

	/// Do I need to check for errors here?
	binary.Write(buf, binary.LittleEndian, i.Get())

	return buf.Bytes()
}

func (i Binary) Big() (big []byte) {

	buf := new(bytes.Buffer)

	/// Do I need to check for errors here?
	binary.Write(buf, binary.BigEndian, i.Get())

	return buf.Bytes()
}

func IsRawCode(rawCode string) bool {
	if len(rawCode) != 4 {
		return false
	}

	for _, r := range rawCode {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

/*

// ### Bugged, add error
func LittleToBig(littleByte []byte) (bigByte []byte, err error) {

	buf := new(bytes.Buffer)
	for i := len(littleByte) - 1; i >= 0; i-- {
		err := binary.Write(buf, binary.BigEndian, littleByte[i])

		if err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

// ### Bugged, add error
func BigToLittle(bigByte []byte) (littleByte []byte) {

	buf := new(bytes.Buffer)
	for _, aByte := range bigByte {
		_ = binary.Write(buf, binary.LittleEndian, aByte)
	}

	return buf.Bytes()
}
*/
