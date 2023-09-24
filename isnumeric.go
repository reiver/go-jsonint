package jsonint

func isNumeric(data []byte) bool {
	if len(data) <= 0 {
		return false
	}

	{
		datum := data[0]

		if '-' == datum || '+' == datum {
			data = data[1:]

			if len(data) <= 0 {
				return false
			}
		}
	}

	// This is safe to do even though it is UTF-8 encoded Unicode.
	for _, datum := range data {

		if datum < '0' || '9' < datum {
			return false
		}
	}

	return true
}
