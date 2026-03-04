package jsonint_test

import (
	"testing"

	"github.com/reiver/go-jsonint"
)

func TestInt_Normalize(t *testing.T) {

	tests := []struct{
		Int jsonint.Int
		Expected string
	}{
		// zero value
		{
			Int: jsonint.Int{},
			Expected: "0",
		},

		// from Int64
		{
			Int: jsonint.Int64(0),
			Expected: "0",
		},
		{
			Int: jsonint.Int64(1),
			Expected: "1",
		},
		{
			Int: jsonint.Int64(-1),
			Expected: "-1",
		},
		{
			Int: jsonint.Int64(123),
			Expected: "123",
		},
		{
			Int: jsonint.Int64(-123),
			Expected: "-123",
		},

		// from Uint64
		{
			Int: jsonint.Uint64(0),
			Expected: "0",
		},
		{
			Int: jsonint.Uint64(456),
			Expected: "456",
		},

		// leading +
		{
			Int: jsonint.MustString("+0"),
			Expected: "0",
		},
		{
			Int: jsonint.MustString("+1"),
			Expected: "1",
		},
		{
			Int: jsonint.MustString("+5"),
			Expected: "5",
		},
		{
			Int: jsonint.MustString("+123"),
			Expected: "123",
		},

		// leading zeros
		{
			Int: jsonint.MustString("00"),
			Expected: "0",
		},
		{
			Int: jsonint.MustString("000"),
			Expected: "0",
		},
		{
			Int: jsonint.MustString("007"),
			Expected: "7",
		},
		{
			Int: jsonint.MustString("0123"),
			Expected: "123",
		},

		// negative with leading zeros
		{
			Int: jsonint.MustString("-0"),
			Expected: "0",
		},
		{
			Int: jsonint.MustString("-00"),
			Expected: "0",
		},
		{
			Int: jsonint.MustString("-000"),
			Expected: "0",
		},
		{
			Int: jsonint.MustString("-007"),
			Expected: "-7",
		},
		{
			Int: jsonint.MustString("-0123"),
			Expected: "-123",
		},

		// positive sign with leading zeros
		{
			Int: jsonint.MustString("+00"),
			Expected: "0",
		},
		{
			Int: jsonint.MustString("+000"),
			Expected: "0",
		},
		{
			Int: jsonint.MustString("+007"),
			Expected: "7",
		},
		{
			Int: jsonint.MustString("+0123"),
			Expected: "123",
		},
	}

	for testNumber, test := range tests {

		actual := test.Int.Normalize()
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual result from Normalize() is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
