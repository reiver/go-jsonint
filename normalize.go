package jsonint

// normalize returns a valid JSON integer.
//
// [Integer] will accept (without error) strings that are valid integers but invalid JSON integers.
// (For exampple: "+5".)
// Normalize returns the valid JSON integer form.
// (For example: "5" rather than "+5".)
func normalize(str string) string {
	if "" == str {
		return "0"
	}

	negative := false

	switch str[0] {
	case '-':
		negative = true
		str = str[1:]
	case '+':
		str = str[1:]
	}

	if "" == str {
		return "0"
	}

	// strip leading zeros, keeping at least one digit
	i := 0
	for i < len(str)-1 && str[i] == '0' {
		i++
	}
	str = str[i:]

	if negative && "0" != str {
		return "-" + str
	}

	return str
}
