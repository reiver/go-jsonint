package jsonint

import (
	"encoding/json"
	"fmt"
	"strconv"

	"codeberg.org/reiver/go-erorr"
)

var _ json.Marshaler = Integer{}
var _ json.Unmarshaler = new(Integer)

// Integer stores a JSON integer in precise, exact way.
//
// This is in constast to many JSON implementations that either store JSON integers in lossy, inexact ways, or ways that risk overflow or underflow.
type Integer struct {
	// Zero is always stored as an empty string.
	// Everything else is stored in normalized format that is a valid JSON integer.
	value string
}

// IntegerFromBytes returns an [Integer] with the same value of the provided []byte with a JSON integer value.
//
// See also: [IntegerFromString].
func IntegerFromBytes(value []byte) (Integer, bool) {
	var str string = string(value)
	return IntegerFromString(str)
}

// IntegerFromInt64 returns an [Integer] with the same value of the provided int64.
func IntegerFromInt64(value int64) Integer {
	if 0 == value {
		var zero Integer
		return zero
	}

	var str string = strconv.FormatInt(value, 10)

	var result Integer
	result.set(str)

	return result
}

// MustIntegerFromBytes is similr to [IntegerFromBytes] but it panic()s if the []byte does not represent an integer.
//
// See also [IntegerFromBytes].
func MustIntegerFromBytes(value []byte) Integer {
	result, ok := IntegerFromBytes(value)
	if !ok {
		panic( fmt.Sprintf("jsonint: %q is not an integer", string(value)) )
	}

	return result
}

// MustIntegerFromString is similr to [IntegerFromString] but it panic()s if the string does not represent an integer.
//
// See also [IntegerFromString].
func MustIntegerFromString(value string) Integer {
	result, ok := IntegerFromString(value)
	if !ok {
		panic( fmt.Sprintf("jsonint: %q is not an integer", value) )
	}

	return result
}

// IntegerFromString returns an [Integer] with the same value of the provided string with a JSON integer value.
//
// See also: [IntegerFromBytes].
func IntegerFromString(value string) (Integer, bool) {
	if !IsIntegerString(value) {
		var nada Integer
		return nada, false
	}

	var result Integer
	result.set(value)

	return result, true
}

// IntegerFromUint64 returns an [Integer] with the same value of the provided uint64.
func IntegerFromUint64(value uint64) Integer {
	if 0 == value {
		var zero Integer
		return zero
	}

	var str string = strconv.FormatUint(value, 10)

	var result Integer
	result.set(str)

	return result
}

func (receiver Integer) get() string {
	if "" == receiver.value {
		return "0"
	}

	return receiver.value
}

// MarshalJSON makes [Integer] fit [json.Marshaler].
func (receiver Integer) MarshalJSON() ([]byte, error) {
	return []byte(receiver.get()), nil
}

func (receiver *Integer) set(value string) {
	if nil == receiver {
		panic(errNilReceiver)
	}

	value = normalize(value)
	if "0" == value {
		value = ""
	}

	receiver.value = value
}

// String makes [Integer] fit [fmt.Stringer].
func (receiver Integer) String() string {
	return receiver.get()
}

// UnmarshalJSON makes [Integer] fit [json.Unmarshaler].
func (receiver *Integer) UnmarshalJSON(data []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	if !IsIntegerBytes(data) {
		return erorr.Errorf("jsonint: cannot unmarshal %q into value of type %T", data, Integer{})
	}

	var str string = string(data)

	receiver.set(str)
	return nil
}
