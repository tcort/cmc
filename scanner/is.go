package scanner

func isDigit(b byte) bool {
	return (b >= '0' && b <= '9')
}

func isLetter(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z')
}

func isWhitespace(b byte) bool {
	return (b == ' ' || b == '\n' || b == '\t')
}

func isSpecialStartChar(b byte) bool {
	return (b == '+' ||
		b == '-' ||
		b == '*' ||
		b == '/' ||
		b == '<' ||
		b == '>' ||
		b == '!' ||
		b == '=' ||
		b == ';' ||
		b == ',' ||
		b == '(' ||
		b == ')' ||
		b == '{' ||
		b == '}' ||
		b == '[' ||
		b == ']')
}
