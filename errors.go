package jsonint

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errNilReceiver                  = erorr.Error("jsonint: nil receiver")
	errCannotMarshalNothingIntoJSON = erorr.Error("jsonint: cannot marshal jsonint.Nothing() into JSON")
)
