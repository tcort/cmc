package scanner

import (
	"testing"
)

func TestIsDigit(t *testing.T) {
	ndigits := 0
	for i := 0; i < 256; i++ {
		if isDigit(byte(i)) {
			ndigits++
		}
	}
	if ndigits != 10 {
		t.Errorf("Found %d digits in alphabet, expected 10", ndigits)
	}

	if !isDigit('0') {
		t.Errorf("Expected %c to be true", '0')
	}
	if !isDigit('5') {
		t.Errorf("Expected %c to be true", '5')
	}
	if !isDigit('9') {
		t.Errorf("Expected %c to be true", '9')
	}

	if isDigit('T') {
		t.Errorf("Expected %c to be false", 'T')
	}
}

func TestIsLetter(t *testing.T) {
	nletters := 0
	for i := 0; i < 256; i++ {
		if isLetter(byte(i)) {
			nletters++
		}
	}
	if nletters != 52 {
		t.Errorf("Found %d letters in alphabet, expected 52", nletters)
	}

	if !isLetter('T') {
		t.Errorf("Expected %c to be true", 'T')
	}
	if !isLetter('t') {
		t.Errorf("Expected %c to be true", 't')
	}
	if !isLetter('A') {
		t.Errorf("Expected %c to be true", 'A')
	}
	if !isLetter('z') {
		t.Errorf("Expected %c to be true", 'z')
	}

	if isLetter('7') {
		t.Errorf("Expected %c to be false", '7')
	}
}

func TestIsWhitespace(t *testing.T) {
	nwhitespace := 0
	for i := 0; i < 256; i++ {
		if isWhitespace(byte(i)) {
			nwhitespace++
		}
	}
	if nwhitespace != 3 {
		t.Errorf("Found %d whitespace characters in alphabet, expected 3", nwhitespace)
	}

	if !isWhitespace(' ') {
		t.Errorf("Expected ' ' to be true")
	}
	if !isWhitespace('\n') {
		t.Errorf("Expected '\\n' to be true")
	}
	if !isWhitespace('\t') {
		t.Errorf("Expected '\\t' to be true")
	}

	if isWhitespace('7') {
		t.Errorf("Expected %c to be false", '7')
	}
}

func TestIsSpecialStartChar(t *testing.T) {
	nspecials := 0
	for i := 0; i < 256; i++ {
		if isSpecialStartChar(byte(i)) {
			nspecials++
		}
	}
	if nspecials != 16 {
		t.Errorf("Found %d special characters in alphabet, expected 16", nspecials)
	}

	if !isSpecialStartChar('+') {
		t.Errorf("Expected %c to be true", '+')
	}

	if isSpecialStartChar('7') {
		t.Errorf("Expected %c to be false", '7')
	}
}
