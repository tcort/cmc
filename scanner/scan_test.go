package scanner

import (
	"testing"
)

func TestScan(t *testing.T) {
	var tokens []Token

	tokens, _ = Scan("test", "{")
	if len(tokens) != 1 || tokens[0].Type != TokenOpenBlock || tokens[0].Text != "{" {
		t.Errorf("Expected { token")
	}

	tokens, _ = Scan("test", ">=")
	if len(tokens) != 1 || tokens[0].Type != TokenGreaterThanOrEqualTo || tokens[0].Text != ">=" {
		t.Errorf("Expected >= token")
	}

	tokens, _ = Scan("test", "void")
	if len(tokens) != 1 || tokens[0].Type != TokenVoid || tokens[0].Text != "void" {
		t.Errorf("Expected void token")
	}

	tokens, _ = Scan("test", "integer")
	if len(tokens) != 1 || tokens[0].Type != TokenIdentifier || tokens[0].Text != "integer" {
		t.Errorf("Expected integer token")
	}

	tokens, _ = Scan("test", "42")
	if len(tokens) != 1 || tokens[0].Type != TokenNumber || tokens[0].Text != "42" {
		t.Errorf("Expected 42 token")
	}

	tokens, _ = Scan("test", "/**/42")
	if len(tokens) != 1 || tokens[0].Type != TokenNumber || tokens[0].Text != "42" {
		t.Errorf("Expected 42 token")
	}

	tokens, _ = Scan("test", "42/**/")
	if len(tokens) != 1 || tokens[0].Type != TokenNumber || tokens[0].Text != "42" {
		t.Errorf("Expected 42 token")
	}

	tokens, _ = Scan("test", "4 2")
	if len(tokens) != 2 ||
		tokens[0].Type != TokenNumber || tokens[0].Text != "4" ||
		tokens[1].Type != TokenNumber || tokens[1].Text != "2" {
		t.Errorf("Expected 4 2 tokens")
	}

	tokens, _ = Scan("test", "x=4")
	if len(tokens) != 3 ||
		tokens[0].Type != TokenIdentifier || tokens[0].Text != "x" ||
		tokens[1].Type != TokenAssign || tokens[1].Text != "=" ||
		tokens[2].Type != TokenNumber || tokens[2].Text != "4" {
		t.Errorf("Expected x = 4 tokens")
	}
}
