package jsonint_test

import (
	"testing"

	"github.com/reiver/go-jsonint"

)

func TestIsNumericBytes(t *testing.T) {

	tests := []struct{
		Value []byte
		Expected bool
	}{
		{
			Value: nil,
			Expected: false,
		},



		{
			Value: []byte(""),
			Expected: false,
		},



		{
			Value: []byte(" "),
			Expected: false,
		},
		{
			Value: []byte("  "),
			Expected: false,
		},
		{
			Value: []byte("   "),
			Expected: false,
		},
		{
			Value: []byte("    "),
			Expected: false,
		},
		{
			Value: []byte("     "),
			Expected: false,
		},



		{
			Value: []byte("apple"),
			Expected: false,
		},
		{
			Value: []byte("banana"),
			Expected: false,
		},
		{
			Value: []byte("cherry"),
			Expected: false,
		},



		{
			Value: []byte("ONCE"),
			Expected: false,
		},
		{
			Value: []byte("TWICE"),
			Expected: false,
		},
		{
			Value: []byte("THRICE"),
			Expected: false,
		},
		{
			Value: []byte("FOURCE"),
			Expected: false,
		},



		{
			Value: []byte("one"),
			Expected: false,
		},
		{
			Value: []byte("two"),
			Expected: false,
		},
		{
			Value: []byte("three"),
			Expected: false,
		},
		{
			Value: []byte("four"),
			Expected: false,
		},
		{
			Value: []byte("five"),
			Expected: false,
		},



		{
			Value: []byte(`"-255"`),
			Expected: false,
		},

		{
			Value: []byte(`"-5"`),
			Expected: false,
		},
		{
			Value: []byte(`"-4"`),
			Expected: false,
		},
		{
			Value: []byte(`"-3"`),
			Expected: false,
		},
		{
			Value: []byte(`"-2"`),
			Expected: false,
		},
		{
			Value: []byte(`"-1"`),
			Expected: false,
		},
		{
			Value: []byte(`"0"`),
			Expected: false,
		},
		{
			Value: []byte(`"1"`),
			Expected: false,
		},
		{
			Value: []byte(`"2"`),
			Expected: false,
		},
		{
			Value: []byte(`"3"`),
			Expected: false,
		},
		{
			Value: []byte(`"4"`),
			Expected: false,
		},
		{
			Value: []byte(`"5"`),
			Expected: false,
		},

		{
			Value: []byte(`"255"`),
			Expected: false,
		},



		{
			Value: []byte("-255"),
			Expected: true,
		},

		{
			Value: []byte("-5"),
			Expected: true,
		},
		{
			Value: []byte("-4"),
			Expected: true,
		},
		{
			Value: []byte("-3"),
			Expected: true,
		},
		{
			Value: []byte("-2"),
			Expected: true,
		},
		{
			Value: []byte("-1"),
			Expected: true,
		},
		{
			Value: []byte("0"),
			Expected: true,
		},
		{
			Value: []byte("1"),
			Expected: true,
		},
		{
			Value: []byte("2"),
			Expected: true,
		},
		{
			Value: []byte("3"),
			Expected: true,
		},
		{
			Value: []byte("4"),
			Expected: true,
		},
		{
			Value: []byte("5"),
			Expected: true,
		},

		{
			Value: []byte("255"),
			Expected: true,
		},



		{
			Value: []byte("-0"),
			Expected: true,
		},
		{
			Value: []byte("+0"),
			Expected: true,
		},



		{
			Value: []byte("+0"),
			Expected: true,
		},
		{
			Value: []byte("+1"),
			Expected: true,
		},
		{
			Value: []byte("+2"),
			Expected: true,
		},
		{
			Value: []byte("+3"),
			Expected: true,
		},
		{
			Value: []byte("+4"),
			Expected: true,
		},
		{
			Value: []byte("+5"),
			Expected: true,
		},
		{
			Value: []byte("+255"),
			Expected: true,
		},



		{
			Value: []byte("-"),
			Expected: false,
		},
		{
			Value: []byte("+"),
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual := jsonint.IsNumericBytes(test.Value)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value return from isNumeric() is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: %#v", test.Value)
			t.Logf("VALUE: %q", test.Value)
			continue
		}

	}
}
