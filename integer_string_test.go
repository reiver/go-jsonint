package jsonint_test

import (
	"testing"

	"github.com/reiver/go-jsonint"
)

func TestInt_String(t *testing.T) {

	tests := []struct{
		Int jsonint.Integer
		Expected string
	}{
		// zero value
		{
			Int: jsonint.Integer{},
			Expected: "0",
		},

		// from IntegerFromInt64
		{
			Int: jsonint.IntegerFromInt64(0),
			Expected: "0",
		},
		{
			Int: jsonint.IntegerFromInt64(1),
			Expected: "1",
		},
		{
			Int: jsonint.IntegerFromInt64(-1),
			Expected: "-1",
		},
		{
			Int: jsonint.IntegerFromInt64(123),
			Expected: "123",
		},
		{
			Int: jsonint.IntegerFromInt64(-123),
			Expected: "-123",
		},
		{
			Int: jsonint.IntegerFromInt64(-65537),
			Expected: "-65537",
		},

		// from Uint64
		{
			Int: jsonint.IntegerFromUint64(0),
			Expected: "0",
		},
		{
			Int: jsonint.IntegerFromUint64(1),
			Expected: "1",
		},
		{
			Int: jsonint.IntegerFromUint64(456),
			Expected: "456",
		},
		{
			Int: jsonint.IntegerFromUint64(65535),
			Expected: "65535",
		},

		// from MustIntegerFromString — values are normalized on input
		{
			Int: jsonint.MustIntegerFromString("+1"),
			Expected: "1",
		},
		{
			Int: jsonint.MustIntegerFromString("+123"),
			Expected: "123",
		},
		{
			Int: jsonint.MustIntegerFromString("+0"),
			Expected: "0",
		},
		{
			Int: jsonint.MustIntegerFromString("-0"),
			Expected: "0",
		},
		{
			Int: jsonint.MustIntegerFromString("007"),
			Expected: "7",
		},
		{
			Int: jsonint.MustIntegerFromString("0123"),
			Expected: "123",
		},
		{
			Int: jsonint.MustIntegerFromString("-007"),
			Expected: "-7",
		},
		{
			Int: jsonint.MustIntegerFromString("-0123"),
			Expected: "-123",
		},
		{
			Int: jsonint.MustIntegerFromString("+007"),
			Expected: "7",
		},
		{
			Int: jsonint.MustIntegerFromString("+0123"),
			Expected: "123",
		},
		{
			Int: jsonint.MustIntegerFromString("-000"),
			Expected: "0",
		},
		{
			Int: jsonint.MustIntegerFromString("+000"),
			Expected: "0",
		},
	}

	for testNumber, test := range tests {

		actual := test.Int.String()
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual result from String() is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
