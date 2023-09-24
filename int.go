package jsonint

import (
	"encoding/json"
	"strconv"

	"sourcecode.social/reiver/go-erorr"
	"sourcecode.social/reiver/go-opt"
)

var _ json.Marshaler = Int{}
var _ json.Unmarshaler = new(Int)

type Int struct {
	value opt.Optional[string]
}

func Nothing() Int {
	return Int{}
}

func Int64(value int64) Int {
	return Int{
		value: opt.Something(strconv.FormatInt(value, 10)),
	}
}

func Uint64(value uint64) Int {
	return Int{
		value: opt.Something(strconv.FormatUint(value, 10)),
	}
}

func (receiver Int) Get() (string, bool) {
	return receiver.value.Get()
}

func (receiver Int) MarshalJSON() ([]byte, error) {
	value, found := receiver.value.Get()
	if !found {
		return nil, errCannotMarshalNothingIntoJSON
	}

	return []byte(value), nil
}

func (receiver *Int) UnmarshalJSON(data []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	if !isNumeric(data) {
		return erorr.Errorf("jsonint: cannot unmarshal %q into value of type %T", data, Int{})
	}

	receiver.value = opt.Something(string(data))

	return nil
}
