package jsonint

import (
	"encoding/json"

	"github.com/reiver/go-erorr"
)

var _ json.Marshaler = Numeric{}
var _ json.Unmarshaler = new(Numeric)

// Numeric is similar to [Integer] but also accepts JSON strings containing integer values.
//
// For example, both 5 and "5" are accepted, but "banana" is not.
type Numeric struct {
	integer Integer
}

// MarshalJSON makes [Numeric] fit [json.Marshaler].
//
// The output is always a bare JSON integer.
// It is never a quoted integer.
func (receiver Numeric) MarshalJSON() ([]byte, error) {
	return receiver.integer.MarshalJSON()
}

// String makes [Numeric] fit [fmt.Stringer].
func (receiver Numeric) String() string {
	return receiver.integer.String()
}

// UnmarshalJSON makes [Numeric] fit [json.Unmarshaler].
func (receiver *Numeric) UnmarshalJSON(data []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	if !IsNumericBytes(data) {
		return erorr.Errorf("jsonint: cannot unmarshal %q into value of type %T", data, Numeric{})
	}

	// Strip surrounding quotes if present.
	if len(data) >= 3 && data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}

	return receiver.integer.UnmarshalJSON(data)
}
