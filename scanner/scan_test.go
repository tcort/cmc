package scanner

import (
	"testing"
)

func TestScan(t *testing.T) {
	var tokens []Token

	tokens, _ = Scan("test", "{")
	if len(tokens) != 1 || tokens[0].Type != T_OPEN_BLOCK || tokens[0].Text != "{" {
		t.Errorf("Expected { token")
	}

	tokens, _ = Scan("test", ">=")
	if len(tokens) != 1 || tokens[0].Type != T_GREATER_THAN_OR_EQUAL_TO || tokens[0].Text != ">=" {
		t.Errorf("Expected >= token")
	}

	tokens, _ = Scan("test", "void")
	if len(tokens) != 1 || tokens[0].Type != T_VOID || tokens[0].Text != "void" {
		t.Errorf("Expected void token")
	}

	tokens, _ = Scan("test", "integer")
	if len(tokens) != 1 || tokens[0].Type != T_IDENTIFIER || tokens[0].Text != "integer" {
		t.Errorf("Expected integer token")
	}

	tokens, _ = Scan("test", "42")
	if len(tokens) != 1 || tokens[0].Type != T_NUMBER || tokens[0].Text != "42" {
		t.Errorf("Expected 42 token")
	}

	tokens, _ = Scan("test", "/**/42")
	if len(tokens) != 1 || tokens[0].Type != T_NUMBER || tokens[0].Text != "42" {
		t.Errorf("Expected 42 token")
	}

	tokens, _ = Scan("test", "42/**/")
	if len(tokens) != 1 || tokens[0].Type != T_NUMBER || tokens[0].Text != "42" {
		t.Errorf("Expected 42 token")
	}

	tokens, _ = Scan("test", "4 2")
	if len(tokens) != 2 ||
		tokens[0].Type != T_NUMBER || tokens[0].Text != "4" ||
		tokens[1].Type != T_NUMBER || tokens[1].Text != "2" {
		t.Errorf("Expected 4 2 tokens")
	}

	tokens, _ = Scan("test", "x=4")
	if len(tokens) != 3 ||
		tokens[0].Type != T_IDENTIFIER || tokens[0].Text != "x" ||
		tokens[1].Type != T_ASSIGN || tokens[1].Text != "=" ||
		tokens[2].Type != T_NUMBER || tokens[2].Text != "4" {
		t.Errorf("Expected x = 4 tokens")
	}
}
