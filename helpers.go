package json

func isDigit(b byte) bool {
	// ASCII [0-9]
	condition := (b >= 48 && b <= 57)
	return condition
}

func isAlpha(b byte) bool {
	// ASCII [A-Z] || [a-z]
	condition := (b >= 65 && b <= 90) || (b >= 97 && b <= 122) || (b >= 32 && b <= 38 && b != 34)
	return condition
}
