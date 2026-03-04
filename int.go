package jsonint

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/reiver/go-erorr"
)

var _ json.Marshaler = Int{}
var _ json.Unmarshaler = new(Int)

type Int struct {
	value string
}

func Bytes(value []byte) (Int, bool) {
	if !IsNumericBytes(value) {
		var nada Int
		return nada, false
	}

	return Int{
		value: string(value),
	}, true
}

func Int64(value int64) Int {
	if 0 == value {
		var zero Int
		return zero
	}

	return Int{
		value: strconv.FormatInt(value, 10),
	}
}

func MustBytes(value []byte) Int {
	result, ok := Bytes(value)
	if !ok {
		panic( fmt.Sprintf("jsonint: %q is not an integer", string(value)) )
	}

	return result
}

func MustString(value string) Int {
	result, ok := String(value)
	if !ok {
		panic( fmt.Sprintf("jsonint: %q is not an integer", value) )
	}

	return result
}

func String(value string) (Int, bool) {
	if !IsNumericString(value) {
		var nada Int
		return nada, false
	}

	return Int{
		value: value,
	}, true
}

func Uint64(value uint64) Int {
	if 0 == value {
		return Int{}
	}

	return Int{
		value: strconv.FormatUint(value, 10),
	}
}

func (receiver Int) MarshalJSON() ([]byte, error) {
	return []byte(receiver.Normalize()), nil
}

func (receiver Int) Normalize() string {
	str := receiver.String()

	if "" == str {
		return "0"
	}

	negative := false

	switch str[0] {
	case '-':
		negative = true
		str = str[1:]
	case '+':
		str = str[1:]
	}

	if "" == str {
		return "0"
	}

	// strip leading zeros, keeping at least one digit
	i := 0
	for i < len(str)-1 && str[i] == '0' {
		i++
	}
	str = str[i:]

	if negative && "0" != str {
		return "-" + str
	}

	return str
}

func (receiver Int) String() string {
	if "" == receiver.value {
		return "0"
	}

	return receiver.value
}

func (receiver *Int) UnmarshalJSON(data []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	if !IsNumericBytes(data) {
		return erorr.Errorf("jsonint: cannot unmarshal %q into value of type %T", data, Int{})
	}

	s := string(data)

	switch s {
	case "0","-0","+0":
		receiver.value = ""
		return nil
	}

	receiver.value = string(data)
	return nil
}
