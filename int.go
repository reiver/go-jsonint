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

// Int64 returns an [Int] with the same value of the provided int64.
func Int64(value int64) Int {
	if 0 == value {
		var zero Int
		return zero
	}

	return Int{
		value: strconv.FormatInt(value, 10),
	}
}

// MustBytes is similr to [Bytes] but it panic()s if the []byte does not represent an integer.
//
// See also [Bytes].
func MustBytes(value []byte) Int {
	result, ok := Bytes(value)
	if !ok {
		panic( fmt.Sprintf("jsonint: %q is not an integer", string(value)) )
	}

	return result
}

// MustString is similr to [String] but it panic()s if the string does not represent an integer.
//
// See also [String].
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

// Uint64 returns an [Int] with the same value of the provided uint64.
func Uint64(value uint64) Int {
	if 0 == value {
		return Int{}
	}

	return Int{
		value: strconv.FormatUint(value, 10),
	}
}

// MarshalJSON makes [Int] fit [json.Marshaler].
func (receiver Int) MarshalJSON() ([]byte, error) {
	return []byte(receiver.Normalize()), nil
}

// Normalize returns [Int] as a valid JSON integer.
//
// [Int] will accept (without error) strings that are valid integers but invalid JSON integers.
// (For exampple: "+5".)
// Normalize returns the valid JSON integer form.
// (For example: "5" rather than "+5".)
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

// String makes [Int] fit [fmt.Stringer].
//
// See also [Normalize].
func (receiver Int) String() string {
	if "" == receiver.value {
		return "0"
	}

	return receiver.value
}

// UnmarshalJSON makes [Int] fit [json.Unmarshaler].
func (receiver *Int) UnmarshalJSON(data []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	if !IsNumericBytes(data) {
		return erorr.Errorf("jsonint: cannot unmarshal %q into value of type %T", data, Int{})
	}

	str := string(data)

	switch s {
	case "0","-0","+0":
		receiver.value = ""
		return nil
	}

	receiver.value = str
	return nil
}
