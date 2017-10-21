package scanner

import (
	"github.com/tcort/cmc/errors"
)

type state struct {
	filename   string
	source     string
	tokens     []Token
	position   int
	length     int
	lineNumber int
}

func Scan(filename string, source string) (tokens []Token, err error) {

	state := state{filename, source, []Token{}, 0, len(source), 1}

	for state.position = 0; state.position < state.length; state.position++ {

		if state.source[state.position] == '\n' {
			state.lineNumber++
		}

		if (state.position+1) < state.length && state.source[state.position:state.position+2] == "/*" {
			err = scan_comment(&state)
			if err != nil {
				return nil, err
			}
		} else if isLetter(state.source[state.position]) {
			scan_keyword_or_id(&state)
			if err != nil {
				return nil, err
			}
		} else if isDigit(state.source[state.position]) {
			scan_number(&state)
			if err != nil {
				return nil, err
			}
		} else if isSpecialStartChar(state.source[state.position]) {
			scan_special(&state)
			if err != nil {
				return nil, err
			}
		} else if isWhitespace(state.source[state.position]) {
			continue
		} else {
			return nil, errors.New("Unexpected token", filename, state.lineNumber)
		}
	}

	return state.tokens, nil
}

func scan_comment(state *state) (err error) {
	found := false
	for state.position = state.position + 1; state.position+1 < state.length; state.position++ {
		if state.source[state.position] == '\n' {
			state.lineNumber++
		}

		if state.source[state.position:state.position+2] == "/*" {
			return errors.New("Nested comment detected", state.filename, state.lineNumber)
		} else if state.source[state.position:state.position+2] == "*/" {
			state.position++
			found = true
			break
		}
	}
	if found {
		return nil
	} else {
		return errors.New("Unclosed comment detected", state.filename, state.lineNumber)
	}
}

func scan_keyword_or_id(state *state) (err error) {
	var begin int
	var end int
	for begin, end = state.position, state.position; end < state.length && isLetter(state.source[end]); end++ {
		continue
	}
	keywords := map[string]TokenType{
		"else":   T_ELSE,
		"if":     T_IF,
		"int":    T_INT,
		"return": T_RETURN,
		"void":   T_VOID,
		"while":  T_WHILE,
	}
	if keyword, ok := keywords[state.source[begin:end]]; ok {
		state.tokens = append(state.tokens, Token{keyword, state.source[begin:end], state.lineNumber})
	} else {
		state.tokens = append(state.tokens, Token{T_IDENTIFIER, state.source[begin:end], state.lineNumber})
	}
	state.position = end - 1
	return nil
}

func scan_number(state *state) (err error) {
	var begin int
	var end int
	for begin, end = state.position, state.position; end < state.length && isDigit(state.source[end]); end++ {
		continue
	}
	state.tokens = append(state.tokens, Token{T_NUMBER, state.source[begin:end], state.lineNumber})
	state.position = end - 1
	return nil
}

func scan_special(state *state) (err error) {

	specials := map[string]TokenType{
		"+":  T_PLUS,
		"-":  T_MINUS,
		"*":  T_TIMES,
		"/":  T_DIVIDE,
		"<":  T_LESS_THAN,
		"<=": T_LESS_THAN_OR_EQUAL_TO,
		">":  T_GREATER_THAN,
		">=": T_GREATER_THAN_OR_EQUAL_TO,
		"==": T_EQUALS,
		"!=": T_NOT_EQUALS,
		"=":  T_ASSIGN,
		";":  T_END_OF_STATEMENT,
		",":  T_LIST_SEPARATOR,
		"(":  T_OPEN_LIST,
		")":  T_CLOSE_LIST,
		"[":  T_OPEN_INDEX,
		"]":  T_CLOSE_INDEX,
		"{":  T_OPEN_BLOCK,
		"}":  T_CLOSE_BLOCK,
	}

	if state.position+1 < state.length {
		if special, ok := specials[state.source[state.position:state.position+2]]; ok {
			state.tokens = append(state.tokens, Token{special, state.source[state.position : state.position+2], state.lineNumber})
			state.position++
			return nil
		}
	}

	if special, ok := specials[state.source[state.position:state.position+1]]; ok {
		state.tokens = append(state.tokens, Token{special, state.source[state.position : state.position+1], state.lineNumber})
		return nil
	}

	return errors.New("Unexpected token", state.filename, state.lineNumber)
}
