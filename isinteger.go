package jsonint

import (
	"unsafe"
)

func IsIntegerBytes(data []byte) bool {
	var str string = unsafe.String(unsafe.SliceData(data), len(data))
	return IsIntegerString(str)
}

func IsIntegerString(data string) bool {
	if len(data) <= 0 {
		return false
	}

	{
		datum := data[0]

		if '-' == datum || '+' == datum {
			data = data[1:]

			if len(data) <= 0 {
				return false
			}
		}
	}

	for _, datum := range data {

		if datum < '0' || '9' < datum {
			return false
		}
	}

	return true
}
