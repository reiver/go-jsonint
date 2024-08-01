package jsonint_test

import (
	"testing"

	"github.com/reiver/go-jsonint"
)

func TestInt_MarshalJSON(t *testing.T) {

	tests := []struct{
		Int jsonint.Int
		Expected string
	}{
		{
			Int: jsonint.Int{},
			Expected: "0",
		},



		{
			Int: jsonint.Int64(-65537),
			Expected: "-65537",
		},
		{
			Int: jsonint.Int64(-65536),
			Expected: "-65536",
		},
		{
			Int: jsonint.Int64(-65535),
			Expected: "-65535",
		},

		{
			Int: jsonint.Int64(-255),
			Expected: "-255",
		},

		{
			Int: jsonint.Int64(-5),
			Expected: "-5",
		},
		{
			Int: jsonint.Int64(-4),
			Expected: "-4",
		},
		{
			Int: jsonint.Int64(-3),
			Expected: "-3",
		},
		{
			Int: jsonint.Int64(-2),
			Expected: "-2",
		},
		{
			Int: jsonint.Int64(-1),
			Expected: "-1",
		},
		{
			Int: jsonint.Int64(0),
			Expected: "0",
		},
		{
			Int: jsonint.Int64(1),
			Expected: "1",
		},
		{
			Int: jsonint.Int64(2),
			Expected: "2",
		},
		{
			Int: jsonint.Int64(3),
			Expected: "3",
		},
		{
			Int: jsonint.Int64(4),
			Expected: "4",
		},
		{
			Int: jsonint.Int64(5),
			Expected: "5",
		},

		{
			Int: jsonint.Int64(255),
			Expected: "255",
		},

		{
			Int: jsonint.Int64(65535),
			Expected: "65535",
		},
		{
			Int: jsonint.Int64(65536),
			Expected: "65536",
		},
		{
			Int: jsonint.Int64(65537),
			Expected: "65537",
		},



		{
			Int: jsonint.Uint64(0),
			Expected: "0",
		},
		{
			Int: jsonint.Uint64(1),
			Expected: "1",
		},
		{
			Int: jsonint.Uint64(2),
			Expected: "2",
		},
		{
			Int: jsonint.Uint64(3),
			Expected: "3",
		},
		{
			Int: jsonint.Uint64(4),
			Expected: "4",
		},
		{
			Int: jsonint.Uint64(5),
			Expected: "5",
		},

		{
			Int: jsonint.Uint64(255),
			Expected: "255",
		},
		{
			Int: jsonint.Uint64(256),
			Expected: "256",
		},
		{
			Int: jsonint.Uint64(257),
			Expected: "257",
		},

		{
			Int: jsonint.Uint64(65535),
			Expected: "65535",
		},
		{
			Int: jsonint.Uint64(65536),
			Expected: "65536",
		},
		{
			Int: jsonint.Uint64(65537),
			Expected: "65537",
		},

		{
			Int: jsonint.Uint64(4294967295),
			Expected: "4294967295",
		},
		{
			Int: jsonint.Uint64(4294967296),
			Expected: "4294967296",
		},
		{
			Int: jsonint.Uint64(4294967297),
			Expected: "4294967297",
		},

		{
			Int: jsonint.Uint64(18446744073709551614),
			Expected: "18446744073709551614",
		},
		{
			Int: jsonint.Uint64(18446744073709551615),
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
