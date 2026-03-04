package jsonint_test

import (
	"testing"

	"github.com/reiver/go-jsonint"
)

func TestInt_String(t *testing.T) {

	tests := []struct{
		Int jsonint.Int
		Expected string
	}{
		// zero value
		{
			Int: jsonint.Int{},
			Expected: "0",
		},

		// from IntFromInt64
		{
			Int: jsonint.IntFromInt64(0),
			Expected: "0",
		},
		{
			Int: jsonint.IntFromInt64(1),
			Expected: "1",
		},
		{
			Int: jsonint.IntFromInt64(-1),
			Expected: "-1",
		},
		{
			Int: jsonint.IntFromInt64(123),
			Expected: "123",
		},
		{
			Int: jsonint.IntFromInt64(-123),
			Expected: "-123",
		},
		{
			Int: jsonint.IntFromInt64(-65537),
			Expected: "-65537",
		},

		// from Uint64
		{
			Int: jsonint.IntFromUint64(0),
			Expected: "0",
		},
		{
			Int: jsonint.IntFromUint64(1),
			Expected: "1",
		},
		{
			Int: jsonint.IntFromUint64(456),
			Expected: "456",
		},
		{
			Int: jsonint.IntFromUint64(65535),
			Expected: "65535",
		},

		// from MustIntFromString — values are normalized on input
		{
			Int: jsonint.MustIntFromString("+1"),
			Expected: "1",
		},
		{
			Int: jsonint.MustIntFromString("+123"),
			Expected: "123",
		},
		{
			Int: jsonint.MustIntFromString("+0"),
			Expected: "0",
		},
		{
			Int: jsonint.MustIntFromString("-0"),
			Expected: "0",
		},
		{
			Int: jsonint.MustIntFromString("007"),
			Expected: "7",
		},
		{
			Int: jsonint.MustIntFromString("0123"),
			Expected: "123",
		},
		{
			Int: jsonint.MustIntFromString("-007"),
			Expected: "-7",
		},
		{
			Int: jsonint.MustIntFromString("-0123"),
			Expected: "-123",
		},
		{
			Int: jsonint.MustIntFromString("+007"),
			Expected: "7",
		},
		{
			Int: jsonint.MustIntFromString("+0123"),
			Expected: "123",
		},
		{
			Int: jsonint.MustIntFromString("-000"),
			Expected: "0",
		},
		{
			Int: jsonint.MustIntFromString("+000"),
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
