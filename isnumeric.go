package jsonint

import (
	"unsafe"
)

func IsNumericBytes(data []byte) bool {
	var str string = unsafe.String(unsafe.SliceData(data), len(data))
	return IsNumericString(str)
}

func IsNumericString(data string) bool {
	if "" == data {
		return false
	}

	// If it is a JSON string, then remove the beginning and closing quotation-marks.
	if '"' == data[0] {
		if len(data) < 3 {
			return false
		}
		if '"' != data[len(data)-1] {
			return false
		}

		data = data[1 : len(data)-1]
	}

	return IsIntegerString(data)
}
