package jsonint_test

import (
	"testing"

	"encoding/json"

	"sourcecode.social/reiver/go-jsonint"
)

func TestZero(t *testing.T) {

	var nothing jsonint.Int

	var  int64zero jsonint.Int = jsonint.Int64(0)
	var uint64zero jsonint.Int = jsonint.Uint64(0)

	var jsonzero jsonint.Int
	{
		var data []byte = []byte("0")

		err := json.Unmarshal(data, &jsonzero)
		if nil != err {
			panic(err)
		}
	}

	if nothing != int64zero {
		t.Error("nothing ≠ int64zero")
	}
	if nothing != uint64zero {
		t.Error("nothing ≠ uint64zero")
	}
	if nothing != jsonzero {
		t.Error("nothing ≠ json64zero")
	}
}
