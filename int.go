package jsonint

import (
	"encoding/json"
	"strconv"

	"sourcecode.social/reiver/go-erorr"
)

var _ json.Marshaler = Int{}
var _ json.Unmarshaler = new(Int)

type Int struct {
	value string
}

func Int64(value int64) Int {
	if 0 == value {
		return Int{}
	}

	return Int{
		value: strconv.FormatInt(value, 10),
	}
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
	return []byte(receiver.String()), nil
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

	if !isNumeric(data) {
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
