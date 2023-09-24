package jsonint_test

import (
	"testing"

	"encoding/json"

	"sourcecode.social/reiver/go-jsonint"
)

func TestInt_UnmarshalJSON(t *testing.T) {

	tests := []struct {
		JSON []byte
		Expected jsonint.Int
	}{
		{
			JSON: []byte("-257"),
			Expected: jsonint.Int64(-257),
		},
		{
			JSON: []byte("-256"),
			Expected: jsonint.Int64(-256),
		},
		{
			JSON: []byte("-255"),
			Expected: jsonint.Int64(-255),
		},
		{
			JSON: []byte("-254"),
			Expected: jsonint.Int64(-254),
		},

		{
			JSON: []byte("-5"),
			Expected: jsonint.Int64(-5),
		},
		{
			JSON: []byte("-4"),
			Expected: jsonint.Int64(-4),
		},
		{
			JSON: []byte("-3"),
			Expected: jsonint.Int64(-3),
		},
		{
			JSON: []byte("-2"),
			Expected: jsonint.Int64(-2),
		},
		{
			JSON: []byte("-1"),
			Expected: jsonint.Int64(-1),
		},



		{
			JSON: []byte("0"),
			Expected: jsonint.Int64(0),
		},
		{
			JSON: []byte("0"),
			Expected: jsonint.Uint64(0),
		},



		{
			JSON: []byte("1"),
			Expected: jsonint.Int64(1),
		},
		{
			JSON: []byte("1"),
			Expected: jsonint.Uint64(1),
		},
		{
			JSON: []byte("+1"),
			Expected: jsonint.Int64(1),
		},
		{
			JSON: []byte("+1"),
			Expected: jsonint.Uint64(1),
		},



		{
			JSON: []byte("2"),
			Expected: jsonint.Int64(2),
		},
		{
			JSON: []byte("2"),
			Expected: jsonint.Uint64(2),
		},
		{
			JSON: []byte("+2"),
			Expected: jsonint.Int64(2),
		},
		{
			JSON: []byte("+2"),
			Expected: jsonint.Uint64(2),
		},



		{
			JSON: []byte("3"),
			Expected: jsonint.Int64(3),
		},
		{
			JSON: []byte("3"),
			Expected: jsonint.Uint64(3),
		},
		{
			JSON: []byte("+3"),
			Expected: jsonint.Int64(3),
		},
		{
			JSON: []byte("+3"),
			Expected: jsonint.Uint64(3),
		},



		{
			JSON: []byte("4"),
			Expected: jsonint.Int64(4),
		},
		{
			JSON: []byte("4"),
			Expected: jsonint.Uint64(4),
		},
		{
			JSON: []byte("+4"),
			Expected: jsonint.Int64(4),
		},
		{
			JSON: []byte("+4"),
			Expected: jsonint.Uint64(4),
		},



		{
			JSON: []byte("5"),
			Expected: jsonint.Int64(5),
		},
		{
			JSON: []byte("5"),
			Expected: jsonint.Uint64(5),
		},
		{
			JSON: []byte("5"),
			Expected: jsonint.Int64(+5),
		},
		{
			JSON: []byte("5"),
			Expected: jsonint.Uint64(+5),
		},






		{
			JSON: []byte("254"),
			Expected: jsonint.Int64(254),
		},
		{
			JSON: []byte("254"),
			Expected: jsonint.Uint64(254),
		},
		{
			JSON: []byte("+254"),
			Expected: jsonint.Int64(254),
		},
		{
			JSON: []byte("+254"),
			Expected: jsonint.Uint64(254),
		},



		{
			JSON: []byte("255"),
			Expected: jsonint.Int64(255),
		},
		{
			JSON: []byte("255"),
			Expected: jsonint.Uint64(255),
		},
		{
			JSON: []byte("+255"),
			Expected: jsonint.Int64(255),
		},
		{
			JSON: []byte("+255"),
			Expected: jsonint.Uint64(255),
		},



		{
			JSON: []byte("256"),
			Expected: jsonint.Int64(256),
		},
		{
			JSON: []byte("256"),
			Expected: jsonint.Uint64(256),
		},
		{
			JSON: []byte("+256"),
			Expected: jsonint.Int64(256),
		},
		{
			JSON: []byte("+256"),
			Expected: jsonint.Uint64(256),
		},



		{
			JSON: []byte("257"),
			Expected: jsonint.Int64(257),
		},
		{
			JSON: []byte("257"),
			Expected: jsonint.Uint64(257),
		},
		{
			JSON: []byte("+257"),
			Expected: jsonint.Int64(257),
		},
		{
			JSON: []byte("+257"),
			Expected: jsonint.Uint64(257),
		},
	}

	for testNumber, test := range tests {

		var actualInt jsonint.Int

		err := actualInt.UnmarshalJSON(test.JSON)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("JSON: %#v", test.JSON)
			t.Logf("JSON: %q", test.JSON)
			t.Logf("EXPECTED: %#v", test.Expected)
			continue
		}
	}
}

func TestInt_UnmarshalJSON_fail(t *testing.T) {

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
	}

	for testNumber, test := range tests {

		var actual jsonint.Int

		err := json.Unmarshal(test.JSON, &actual)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			continue
		}
	}
}
