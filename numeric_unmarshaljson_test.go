package jsonint_test

import (
	"testing"

	"encoding/json"

	"github.com/reiver/go-jsonint"
)

func TestNumeric_UnmarshalJSON(t *testing.T) {

	tests := []struct {
		JSON []byte
		Expected string
	}{
		// bare integers
		{
			JSON: []byte("-257"),
			Expected: "-257",
		},
		{
			JSON: []byte("-256"),
			Expected: "-256",
		},
		{
			JSON: []byte("-255"),
			Expected: "-255",
		},
		{
			JSON: []byte("-254"),
			Expected: "-254",
		},

		{
			JSON: []byte("-5"),
			Expected: "-5",
		},
		{
			JSON: []byte("-4"),
			Expected: "-4",
		},
		{
			JSON: []byte("-3"),
			Expected: "-3",
		},
		{
			JSON: []byte("-2"),
			Expected: "-2",
		},
		{
			JSON: []byte("-1"),
			Expected: "-1",
		},



		{
			JSON: []byte("0"),
			Expected: "0",
		},



		{
			JSON: []byte("1"),
			Expected: "1",
		},
		{
			JSON: []byte("+1"),
			Expected: "1",
		},

		{
			JSON: []byte("2"),
			Expected: "2",
		},
		{
			JSON: []byte("+2"),
			Expected: "2",
		},

		{
			JSON: []byte("3"),
			Expected: "3",
		},
		{
			JSON: []byte("+3"),
			Expected: "3",
		},

		{
			JSON: []byte("4"),
			Expected: "4",
		},
		{
			JSON: []byte("+4"),
			Expected: "4",
		},

		{
			JSON: []byte("5"),
			Expected: "5",
		},
		{
			JSON: []byte("+5"),
			Expected: "5",
		},



		{
			JSON: []byte("254"),
			Expected: "254",
		},
		{
			JSON: []byte("+254"),
			Expected: "254",
		},

		{
			JSON: []byte("255"),
			Expected: "255",
		},
		{
			JSON: []byte("+255"),
			Expected: "255",
		},

		{
			JSON: []byte("256"),
			Expected: "256",
		},
		{
			JSON: []byte("+256"),
			Expected: "256",
		},

		{
			JSON: []byte("257"),
			Expected: "257",
		},
		{
			JSON: []byte("+257"),
			Expected: "257",
		},



		// quoted integers
		{
			JSON: []byte(`"0"`),
			Expected: "0",
		},
		{
			JSON: []byte(`"-0"`),
			Expected: "0",
		},
		{
			JSON: []byte(`"+0"`),
			Expected: "0",
		},

		{
			JSON: []byte(`"1"`),
			Expected: "1",
		},
		{
			JSON: []byte(`"+1"`),
			Expected: "1",
		},
		{
			JSON: []byte(`"-1"`),
			Expected: "-1",
		},

		{
			JSON: []byte(`"5"`),
			Expected: "5",
		},
		{
			JSON: []byte(`"+5"`),
			Expected: "5",
		},
		{
			JSON: []byte(`"-5"`),
			Expected: "-5",
		},

		{
			JSON: []byte(`"123"`),
			Expected: "123",
		},
		{
			JSON: []byte(`"+123"`),
			Expected: "123",
		},
		{
			JSON: []byte(`"-123"`),
			Expected: "-123",
		},

		{
			JSON: []byte(`"255"`),
			Expected: "255",
		},
		{
			JSON: []byte(`"-255"`),
			Expected: "-255",
		},

		{
			JSON: []byte(`"256"`),
			Expected: "256",
		},
		{
			JSON: []byte(`"-256"`),
			Expected: "-256",
		},

		{
			JSON: []byte(`"257"`),
			Expected: "257",
		},
		{
			JSON: []byte(`"-257"`),
			Expected: "-257",
		},



		// quoted with leading zeros
		{
			JSON: []byte(`"007"`),
			Expected: "7",
		},
		{
			JSON: []byte(`"+007"`),
			Expected: "7",
		},
		{
			JSON: []byte(`"-007"`),
			Expected: "-7",
		},
	}

	for testNumber, test := range tests {

		var actual jsonint.Numeric

		err := actual.UnmarshalJSON(test.JSON)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("JSON: %#v", test.JSON)
			t.Logf("JSON: %q", test.JSON)
			t.Logf("EXPECTED: %q", test.Expected)
			continue
		}

		{
			actualStr := actual.String()
			expected := test.Expected

			if expected != actualStr {
				t.Errorf("For test #%d, the actual result from UnmarshalJSON() is not what was expected.", testNumber)
				t.Logf("JSON:     %q", test.JSON)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actualStr)
				continue
			}
		}
	}
}

func TestNumeric_UnmarshalJSON_fail(t *testing.T) {

	tests := []struct{
		JSON []byte
	}{
		{
			JSON: nil,
		},



		{
			JSON: []byte{},
		},



		{
			JSON: []byte(""),
		},



		{
			JSON: []byte(" "),
		},
		{
			JSON: []byte("  "),
		},
		{
			JSON: []byte("   "),
		},



		{
			JSON: []byte("apple"),
		},
		{
			JSON: []byte("banana"),
		},
		{
			JSON: []byte("cherry"),
		},



		{
			JSON: []byte("ONE"),
		},
		{
			JSON: []byte("TWO"),
		},
		{
			JSON: []byte("THREE"),
		},
		{
			JSON: []byte("FOUR"),
		},
		{
			JSON: []byte("FIVE"),
		},



		// quoted non-integers
		{
			JSON: []byte(`"apple"`),
		},
		{
			JSON: []byte(`"banana"`),
		},
		{
			JSON: []byte(`"cherry"`),
		},
		{
			JSON: []byte(`""`),
		},
		{
			JSON: []byte(`" "`),
		},
		{
			JSON: []byte(`"hello world"`),
		},
	}

	for testNumber, test := range tests {

		var actual jsonint.Numeric

		err := json.Unmarshal(test.JSON, &actual)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("JSON: %q", test.JSON)
			continue
		}
	}
}
