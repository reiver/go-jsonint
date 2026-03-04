package jsonint_test

import (
	"testing"

	"github.com/reiver/go-jsonint"
)

func TestInt_MarshalJSON(t *testing.T) {

	tests := []struct{
		Int jsonint.Integer
		Expected string
	}{
		{
			Int: jsonint.Integer{},
			Expected: "0",
		},



		{
			Int: jsonint.IntegerFromInt64(-65537),
			Expected: "-65537",
		},
		{
			Int: jsonint.IntegerFromInt64(-65536),
			Expected: "-65536",
		},
		{
			Int: jsonint.IntegerFromInt64(-65535),
			Expected: "-65535",
		},

		{
			Int: jsonint.IntegerFromInt64(-255),
			Expected: "-255",
		},

		{
			Int: jsonint.IntegerFromInt64(-5),
			Expected: "-5",
		},
		{
			Int: jsonint.IntegerFromInt64(-4),
			Expected: "-4",
		},
		{
			Int: jsonint.IntegerFromInt64(-3),
			Expected: "-3",
		},
		{
			Int: jsonint.IntegerFromInt64(-2),
			Expected: "-2",
		},
		{
			Int: jsonint.IntegerFromInt64(-1),
			Expected: "-1",
		},
		{
			Int: jsonint.IntegerFromInt64(0),
			Expected: "0",
		},
		{
			Int: jsonint.IntegerFromInt64(1),
			Expected: "1",
		},
		{
			Int: jsonint.IntegerFromInt64(2),
			Expected: "2",
		},
		{
			Int: jsonint.IntegerFromInt64(3),
			Expected: "3",
		},
		{
			Int: jsonint.IntegerFromInt64(4),
			Expected: "4",
		},
		{
			Int: jsonint.IntegerFromInt64(5),
			Expected: "5",
		},

		{
			Int: jsonint.IntegerFromInt64(255),
			Expected: "255",
		},

		{
			Int: jsonint.IntegerFromInt64(65535),
			Expected: "65535",
		},
		{
			Int: jsonint.IntegerFromInt64(65536),
			Expected: "65536",
		},
		{
			Int: jsonint.IntegerFromInt64(65537),
			Expected: "65537",
		},



		{
			Int: jsonint.IntegerFromUint64(0),
			Expected: "0",
		},
		{
			Int: jsonint.IntegerFromUint64(1),
			Expected: "1",
		},
		{
			Int: jsonint.IntegerFromUint64(2),
			Expected: "2",
		},
		{
			Int: jsonint.IntegerFromUint64(3),
			Expected: "3",
		},
		{
			Int: jsonint.IntegerFromUint64(4),
			Expected: "4",
		},
		{
			Int: jsonint.IntegerFromUint64(5),
			Expected: "5",
		},

		{
			Int: jsonint.IntegerFromUint64(255),
			Expected: "255",
		},
		{
			Int: jsonint.IntegerFromUint64(256),
			Expected: "256",
		},
		{
			Int: jsonint.IntegerFromUint64(257),
			Expected: "257",
		},

		{
			Int: jsonint.IntegerFromUint64(65535),
			Expected: "65535",
		},
		{
			Int: jsonint.IntegerFromUint64(65536),
			Expected: "65536",
		},
		{
			Int: jsonint.IntegerFromUint64(65537),
			Expected: "65537",
		},

		{
			Int: jsonint.IntegerFromUint64(4294967295),
			Expected: "4294967295",
		},
		{
			Int: jsonint.IntegerFromUint64(4294967296),
			Expected: "4294967296",
		},
		{
			Int: jsonint.IntegerFromUint64(4294967297),
			Expected: "4294967297",
		},

		{
			Int: jsonint.IntegerFromUint64(18446744073709551614),
			Expected: "18446744073709551614",
		},
		{
			Int: jsonint.IntegerFromUint64(18446744073709551615),
			Expected: "18446744073709551615",
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := test.Int.MarshalJSON()
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
