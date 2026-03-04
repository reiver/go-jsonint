package jsonint_test

import (
	"testing"

	"encoding/json"

	"github.com/reiver/go-jsonint"
)

func TestNumeric_Zero(t *testing.T) {

	var nothing jsonint.Numeric

	var  int64zero jsonint.Numeric = jsonint.NumericFromInt64(0)
	var uint64zero jsonint.Numeric = jsonint.NumericFromUint64(0)

	var jsonzero jsonint.Numeric
	{
		var data []byte = []byte("0")

		err := json.Unmarshal(data, &jsonzero)
		if nil != err {
			panic(err)
		}
	}

	var jsonquotedzero jsonint.Numeric
	{
		var data []byte = []byte(`"0"`)

		err := json.Unmarshal(data, &jsonquotedzero)
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
		t.Error("nothing ≠ jsonzero")
	}
	if nothing != jsonquotedzero {
		t.Error("nothing ≠ jsonquotedzero")
	}
}
