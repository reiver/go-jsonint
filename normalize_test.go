package jsonint

import (
	"testing"
)

func TestNormalize(t *testing.T) {

	tests := []struct{
		String string
		Expected string
	}{
		// zero value
		{
			String:   "",
			Expected: "0",
		},



		{
			String:   "0",
			Expected: "0",
		},
		{
			String:  "+0",
			Expected: "0",
		},
		{
			String:  "-0",
			Expected: "0",
		},



		{
			String:   "1",
			Expected: "1",
		},
		{
			String:  "+1",
			Expected: "1",
		},
		{
			String:   "-1",
			Expected: "-1",
		},
		{
			String:   "123",
			Expected: "123",
		},
		{
			String:  "+123",
			Expected: "123",
		},
		{
			String:   "-123",
			Expected: "-123",
		},
		{
			String:   "4567",
			Expected: "4567",
		},
		{
			String:  "+4567",
			Expected: "4567",
		},
		{
			String:   "-4567",
			Expected: "-4567",
		},



		// leading zeros
		{
			String:  "00",
			Expected: "0",
		},
		{
			String: "000",
			Expected: "0",
		},
		{
			String: "0000",
			Expected:  "0",
		},
		{
			String:  "07",
			Expected: "7",
		},
		{
			String: "007",
			Expected: "7",
		},
		{
			String: "000000000000007",
			Expected:             "7",
		},
		{
			String:  "0123",
			Expected: "123",
		},
		{
			String: "00123",
			Expected: "123",
		},
		{
			String: "000123",
			Expected:  "123",
		},
		{
			String: "0000123",
			Expected:   "123",
		},



		// negative with leading zeros
		{
			String:   "-0",
			Expected:  "0",
		},
		{
			String:  "-00",
			Expected:  "0",
		},
		{
			String: "-000",
			Expected:  "0",
		},
		{
			String: "+0000",
			Expected:   "0",
		},
		{
			String:  "-07",
			Expected: "-7",
		},
		{
			String: "-007",
			Expected: "-7",
		},
		{
			String: "-000000000000007",
			Expected:             "-7",
		},
		{
			String:  "-0123",
			Expected: "-123",
		},
		{
			String: "-00123",
			Expected: "-123",
		},
		{
			String: "-000123",
			Expected:  "-123",
		},
		{
			String: "-0000123",
			Expected:   "-123",
		},



		// positive sign with leading zeros
		{
			String:   "+0",
			Expected:  "0",
		},
		{
			String:  "+00",
			Expected:  "0",
		},
		{
			String: "+000",
			Expected:  "0",
		},
		{
			String: "+0000",
			Expected:   "0",
		},
		{
			String:  "+07",
			Expected:  "7",
		},
		{
			String: "+007",
			Expected:  "7",
		},
		{
			String: "+000000000000007",
			Expected:              "7",
		},
		{
			String:  "+0123",
			Expected:  "123",
		},
		{
			String: "+00123",
			Expected:  "123",
		},
		{
			String: "+000123",
			Expected:   "123",
		},
		{
			String: "+0000123",
			Expected:    "123",
		},
	}

	for testNumber, test := range tests {

		actual := normalize(test.String)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual result from Normalize() is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
