package jsonint_test

import (
	"testing"

	"github.com/reiver/go-jsonint"
)

func TestNumeric_MarshalJSON(t *testing.T) {

	tests := []struct{
		Numeric jsonint.Numeric
		Expected string
	}{
		{
			Numeric: jsonint.Numeric{},
			Expected: "0",
		},



		{
			Numeric: jsonint.NumericFromInt64(-65537),
			Expected: "-65537",
		},
		{
			Numeric: jsonint.NumericFromInt64(-65536),
			Expected: "-65536",
		},
		{
			Numeric: jsonint.NumericFromInt64(-65535),
			Expected: "-65535",
		},

		{
			Numeric: jsonint.NumericFromInt64(-255),
			Expected: "-255",
		},

		{
			Numeric: jsonint.NumericFromInt64(-5),
			Expected: "-5",
		},
		{
			Numeric: jsonint.NumericFromInt64(-4),
			Expected: "-4",
		},
		{
			Numeric: jsonint.NumericFromInt64(-3),
			Expected: "-3",
		},
		{
			Numeric: jsonint.NumericFromInt64(-2),
			Expected: "-2",
		},
		{
			Numeric: jsonint.NumericFromInt64(-1),
			Expected: "-1",
		},
		{
			Numeric: jsonint.NumericFromInt64(0),
			Expected: "0",
		},
		{
			Numeric: jsonint.NumericFromInt64(1),
			Expected: "1",
		},
		{
			Numeric: jsonint.NumericFromInt64(2),
			Expected: "2",
		},
		{
			Numeric: jsonint.NumericFromInt64(3),
			Expected: "3",
		},
		{
			Numeric: jsonint.NumericFromInt64(4),
			Expected: "4",
		},
		{
			Numeric: jsonint.NumericFromInt64(5),
			Expected: "5",
		},

		{
			Numeric: jsonint.NumericFromInt64(255),
			Expected: "255",
		},

		{
			Numeric: jsonint.NumericFromInt64(65535),
			Expected: "65535",
		},
		{
			Numeric: jsonint.NumericFromInt64(65536),
			Expected: "65536",
		},
		{
			Numeric: jsonint.NumericFromInt64(65537),
			Expected: "65537",
		},



		{
			Numeric: jsonint.NumericFromUint64(0),
			Expected: "0",
		},
		{
			Numeric: jsonint.NumericFromUint64(1),
			Expected: "1",
		},
		{
			Numeric: jsonint.NumericFromUint64(2),
			Expected: "2",
		},
		{
			Numeric: jsonint.NumericFromUint64(3),
			Expected: "3",
		},
		{
			Numeric: jsonint.NumericFromUint64(4),
			Expected: "4",
		},
		{
			Numeric: jsonint.NumericFromUint64(5),
			Expected: "5",
		},

		{
			Numeric: jsonint.NumericFromUint64(255),
			Expected: "255",
		},
		{
			Numeric: jsonint.NumericFromUint64(256),
			Expected: "256",
		},
		{
			Numeric: jsonint.NumericFromUint64(257),
			Expected: "257",
		},

		{
			Numeric: jsonint.NumericFromUint64(65535),
			Expected: "65535",
		},
		{
			Numeric: jsonint.NumericFromUint64(65536),
			Expected: "65536",
		},
		{
			Numeric: jsonint.NumericFromUint64(65537),
			Expected: "65537",
		},

		{
			Numeric: jsonint.NumericFromUint64(4294967295),
			Expected: "4294967295",
		},
		{
			Numeric: jsonint.NumericFromUint64(4294967296),
			Expected: "4294967296",
		},
		{
			Numeric: jsonint.NumericFromUint64(4294967297),
			Expected: "4294967297",
		},

		{
			Numeric: jsonint.NumericFromUint64(18446744073709551614),
			Expected: "18446744073709551614",
		},
		{
			Numeric: jsonint.NumericFromUint64(18446744073709551615),
			Expected: "18446744073709551615",
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := test.Numeric.MarshalJSON()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Errorf("ERROR: (%T) %s", err, err)
			continue
		}

		{
			actual := string(actualBytes)
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual result from MarshalJSON() is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				continue
			}
		}
	}
}
