package jsonint_test

import (
	"testing"

	"github.com/reiver/go-jsonint"
)

func TestIsNumericString(t *testing.T) {

	tests := []struct{
		Value string
		Expected bool
	}{
		{
			Value: "",
			Expected: false,
		},



		{
			Value: " ",
			Expected: false,
		},
		{
			Value: "  ",
			Expected: false,
		},
		{
			Value: "   ",
			Expected: false,
		},



		// bare integers
		{
			Value: "0",
			Expected: true,
		},
		{
			Value: "1",
			Expected: true,
		},
		{
			Value: "5",
			Expected: true,
		},
		{
			Value: "123",
			Expected: true,
		},
		{
			Value: "-1",
			Expected: true,
		},
		{
			Value: "-123",
			Expected: true,
		},
		{
			Value: "+1",
			Expected: true,
		},
		{
			Value: "+123",
			Expected: true,
		},
		{
			Value: "007",
			Expected: true,
		},



		// quoted integers
		{
			Value: `"0"`,
			Expected: true,
		},
		{
			Value: `"1"`,
			Expected: true,
		},
		{
			Value: `"5"`,
			Expected: true,
		},
		{
			Value: `"123"`,
			Expected: true,
		},
		{
			Value: `"-1"`,
			Expected: true,
		},
		{
			Value: `"-123"`,
			Expected: true,
		},
		{
			Value: `"+1"`,
			Expected: true,
		},
		{
			Value: `"+123"`,
			Expected: true,
		},
		{
			Value: `"007"`,
			Expected: true,
		},



		// quoted non-integers
		{
			Value: `"apple"`,
			Expected: false,
		},
		{
			Value: `"banana"`,
			Expected: false,
		},
		{
			Value: `"cherry"`,
			Expected: false,
		},
		{
			Value: `""`,
			Expected: false,
		},
		{
			Value: `" "`,
			Expected: false,
		},
		{
			Value: `"hello world"`,
			Expected: false,
		},



		// non-integer, non-quoted
		{
			Value: "apple",
			Expected: false,
		},
		{
			Value: "banana",
			Expected: false,
		},
		{
			Value: "cherry",
			Expected: false,
		},
		{
			Value: "ONE",
			Expected: false,
		},
		{
			Value: "TWO",
			Expected: false,
		},
		{
			Value: "THREE",
			Expected: false,
		},



		// malformed quotes
		{
			Value: `"123`,
			Expected: false,
		},
		{
			Value: `123"`,
			Expected: false,
		},
		{
			Value: `"`,
			Expected: false,
		},



		{
			Value: `"0"`,
			Expected: true,
		},
		{
			Value: `"+0"`,
			Expected: true,
		},
		{
			Value: `"-0"`,
			Expected: true,
		},
		{
			Value: `"1"`,
			Expected: true,
		},
		{
			Value: `"+1"`,
			Expected: true,
		},
		{
			Value: `"-1"`,
			Expected: true,
		},
		{
			Value: `"4567"`,
			Expected: true,
		},
		{
			Value: `"+4567"`,
			Expected: true,
		},
		{
			Value: `"-4567"`,
			Expected: true,
		},
		{
			Value: `"00000007"`,
			Expected: true,
		},
		{
			Value: `"+00000007"`,
			Expected: true,
		},
		{
			Value: `"-00000007"`,
			Expected: true,
		},
	}

	for testNumber, test := range tests {

		actual := jsonint.IsNumericString(test.Value)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual result from IsNumericString() is not what was expected.", testNumber)
			t.Logf("VALUE:    %q", test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			continue
		}
	}
}
