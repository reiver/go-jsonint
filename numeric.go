package jsonint

import (
	"encoding/json"
	"fmt"

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

func NumericFromBytes(value []byte) (Numeric, bool) {
	var str string = string(value)
	return NumericFromString(str)
}

// NumericFromInt64 returns a [Numeric] with the same value of the provided int64.
func NumericFromInt64(value int64) Numeric {
	return Numeric{
		integer: IntegerFromInt64(value),
	}
}

// NumericFromUint64 returns a [Numeric] with the same value of the provided uint64.
func NumericFromUint64(value uint64) Numeric {
	return Numeric{
		integer: IntegerFromUint64(value),
	}
}

func NumericFromString(value string) (Numeric, bool) {
	if !IsNumericString(value) {
		var nada Numeric
		return nada, false
	}

	// Strip surrounding quotes if present.
	if len(value) >= 3 && value[0] == '"' && value[len(value)-1] == '"' {
		value = value[1 : len(value)-1]
	}

	integer, ok := IntegerFromString(value)
	if !ok {
		var nada Numeric
		return nada, false
	}

	return Numeric{
		integer: integer,
	}, true
}

// MustNumericFromBytes is similar to [NumericFromBytes] but it panic()s if the []byte does not represent a numeric value.
//
// See also [NumericFromBytes].
func MustNumericFromBytes(value []byte) Numeric {
	result, ok := NumericFromBytes(value)
	if !ok {
		panic( fmt.Sprintf("jsonint: %q is not numeric", string(value)) )
	}

	return result
}

// MustNumericFromString is similar to [NumericFromString] but it panic()s if the string does not represent a numeric value.
//
// See also [NumericFromString].
func MustNumericFromString(value string) Numeric {
	result, ok := NumericFromString(value)
	if !ok {
		panic( fmt.Sprintf("jsonint: %q is not numeric", value) )
	}

	return result
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
