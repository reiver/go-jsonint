package jsonint

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/reiver/go-erorr"
)

var _ json.Marshaler = Int{}
var _ json.Unmarshaler = new(Int)

// Int stores a JSON integer in precise, exact way.
//
// This is in constast to many JSON implementations that either store JSON integers in lossy, inexact ways, or ways that risk overflow or underflow.
type Int struct {
	// Zero is always stored as an empty string.
	// Everything else is stored in normalized format that is a valid JSON integer.
	value string
}

func IntFromBytes(value []byte) (Int, bool) {
	var str string = string(value)
	return IntFromString(str)
}

// IntFromInt64 returns an [Int] with the same value of the provided int64.
func IntFromInt64(value int64) Int {
	if 0 == value {
		var zero Int
		return zero
	}

	var str string = strconv.FormatInt(value, 10)

	var result Int
	result.set(str)

	return result
}

// MustIntFromBytes is similr to [IntFromBytes] but it panic()s if the []byte does not represent an integer.
//
// See also [IntFromBytes].
func MustIntFromBytes(value []byte) Int {
	result, ok := IntFromBytes(value)
	if !ok {
		panic( fmt.Sprintf("jsonint: %q is not an integer", string(value)) )
	}

	return result
}

// MustIntFromString is similr to [IntFromString] but it panic()s if the string does not represent an integer.
//
// See also [IntFromString].
func MustIntFromString(value string) Int {
	result, ok := IntFromString(value)
	if !ok {
		panic( fmt.Sprintf("jsonint: %q is not an integer", value) )
	}

	return result
}

func IntFromString(value string) (Int, bool) {
	if !IsIntegerString(value) {
		var nada Int
		return nada, false
	}

	var result Int
	result.set(value)

	return result, true
}

// IntFromUint64 returns an [Int] with the same value of the provided uint64.
func IntFromUint64(value uint64) Int {
	if 0 == value {
		var zero Int
		return zero
	}

	var str string = strconv.FormatUint(value, 10)

	var result Int
	result.set(str)

	return result
}

func (receiver Int) get() string {
	if "" == receiver.value {
		return "0"
	}

	return receiver.value
}

// MarshalJSON makes [Int] fit [json.Marshaler].
func (receiver Int) MarshalJSON() ([]byte, error) {
	return []byte(receiver.get()), nil
}

func (receiver *Int) set(value string) {
	if nil == receiver {
		panic(errNilReceiver)
	}

	value = normalize(value)
	if "0" == value {
		value = ""
	}

	receiver.value = value
}

// String makes [Int] fit [fmt.Stringer].
func (receiver Int) String() string {
	return receiver.get()
}

// UnmarshalJSON makes [Int] fit [json.Unmarshaler].
func (receiver *Int) UnmarshalJSON(data []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	if !IsIntegerBytes(data) {
		return erorr.Errorf("jsonint: cannot unmarshal %q into value of type %T", data, Int{})
	}

	var str string = string(data)

	receiver.set(str)
	return nil
}
