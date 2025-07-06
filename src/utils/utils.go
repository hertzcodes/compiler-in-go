package utils

func IsLetter(ch byte) bool { // TODO: allow numbers in identifiers (e.g: num1, variab5le9)
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func IsDigit(ch byte) bool { // TODO: parse floats
	return '0' <= ch && ch <= '9'
}
