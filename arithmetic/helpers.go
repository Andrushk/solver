package arithmetic

func IsDigit(symbol rune) bool {
	switch symbol {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	}
	return false
}

func IsSeparator(symbol rune) bool {
	if symbol == '.' {
		return true
	}
	return false
}

func IsSign(symbol rune) bool {
	switch symbol {
	case '+', '-': return true
	}
	return false
}
