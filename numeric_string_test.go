package jsonint_test

import (
	"testing"

	"github.com/reiver/go-jsonint"
)

func TestNumeric_String(t *testing.T) {

	tests := []struct{
		Numeric jsonint.Numeric
		Expected string
	}{
		// zero value
		{
			Numeric: jsonint.Numeric{},
			Expected: "0",
		},

		// from NumericFromInt64
		{
			Numeric: jsonint.NumericFromInt64(0),
			Expected: "0",
		},
		{
			Numeric: jsonint.NumericFromInt64(1),
			Expected: "1",
		},
		{
			Numeric: jsonint.NumericFromInt64(-1),
			Expected: "-1",
		},
		{
			Numeric: jsonint.NumericFromInt64(123),
			Expected: "123",
		},
		{
			Numeric: jsonint.NumericFromInt64(-123),
			Expected: "-123",
		},
		{
			Numeric: jsonint.NumericFromInt64(-65537),
			Expected: "-65537",
		},

		// from NumericFromUint64
		{
			Numeric: jsonint.NumericFromUint64(0),
			Expected: "0",
		},
		{
			Numeric: jsonint.NumericFromUint64(1),
			Expected: "1",
		},
		{
			Numeric: jsonint.NumericFromUint64(456),
			Expected: "456",
		},
		{
			Numeric: jsonint.NumericFromUint64(65535),
			Expected: "65535",
		},

		// from MustNumericFromString — bare integers, values are normalized on input
		{
			Numeric: jsonint.MustNumericFromString("+1"),
			Expected: "1",
		},
		{
			Numeric: jsonint.MustNumericFromString("+123"),
			Expected: "123",
		},
		{
			Numeric: jsonint.MustNumericFromString("+0"),
			Expected: "0",
		},
		{
			Numeric: jsonint.MustNumericFromString("-0"),
			Expected: "0",
		},
		{
			Numeric: jsonint.MustNumericFromString("007"),
			Expected: "7",
		},
		{
			Numeric: jsonint.MustNumericFromString("0123"),
			Expected: "123",
		},
		{
			Numeric: jsonint.MustNumericFromString("-007"),
			Expected: "-7",
		},
		{
			Numeric: jsonint.MustNumericFromString("-0123"),
			Expected: "-123",
		},
		{
			Numeric: jsonint.MustNumericFromString("+007"),
			Expected: "7",
		},
		{
			Numeric: jsonint.MustNumericFromString("+0123"),
			Expected: "123",
		},
		{
			Numeric: jsonint.MustNumericFromString("-000"),
			Expected: "0",
		},
		{
			Numeric: jsonint.MustNumericFromString("+000"),
			Expected: "0",
		},

		// from MustNumericFromString — quoted integers
		{
			Numeric: jsonint.MustNumericFromString(`"0"`),
			Expected: "0",
		},
		{
			Numeric: jsonint.MustNumericFromString(`"1"`),
			Expected: "1",
		},
		{
			Numeric: jsonint.MustNumericFromString(`"-1"`),
			Expected: "-1",
		},
		{
			Numeric: jsonint.MustNumericFromString(`"123"`),
			Expected: "123",
		},
		{
			Numeric: jsonint.MustNumericFromString(`"+123"`),
			Expected: "123",
		},
		{
			Numeric: jsonint.MustNumericFromString(`"-123"`),
			Expected: "-123",
		},
		{
			Numeric: jsonint.MustNumericFromString(`"007"`),
			Expected: "7",
		},
		{
			Numeric: jsonint.MustNumericFromString(`"+007"`),
			Expected: "7",
		},
		{
			Numeric: jsonint.MustNumericFromString(`"-007"`),
			Expected: "-7",
		},
	}

	for testNumber, test := range tests {

		actual := test.Numeric.String()
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual result from String() is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
