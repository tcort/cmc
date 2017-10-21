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
			err = scanComment(&state)
			if err != nil {
				return nil, err
			}
		} else if isLetter(state.source[state.position]) {
			scanKeywordOrID(&state)
			if err != nil {
				return nil, err
			}
		} else if isDigit(state.source[state.position]) {
			scanNumber(&state)
			if err != nil {
				return nil, err
			}
		} else if isSpecialStartChar(state.source[state.position]) {
			scanSpecial(&state)
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

func scanComment(state *state) (err error) {
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
	if !found {
		return errors.New("Unclosed comment detected", state.filename, state.lineNumber)
	}

	return nil
}

func scanKeywordOrID(state *state) (err error) {
	var begin int
	var end int
	for begin, end = state.position, state.position; end < state.length && isLetter(state.source[end]); end++ {
		continue
	}
	keywords := map[string]TokenType{
		"else":   TokenElse,
		"if":     TokenIf,
		"int":    TokenInt,
		"return": TokenReturn,
		"void":   TokenVoid,
		"while":  TokenWhile,
	}
	if keyword, ok := keywords[state.source[begin:end]]; ok {
		state.tokens = append(state.tokens, Token{keyword, state.source[begin:end], state.lineNumber})
	} else {
		state.tokens = append(state.tokens, Token{TokenIdentifier, state.source[begin:end], state.lineNumber})
	}
	state.position = end - 1
	return nil
}

func scanNumber(state *state) (err error) {
	var begin int
	var end int
	for begin, end = state.position, state.position; end < state.length && isDigit(state.source[end]); end++ {
		continue
	}
	state.tokens = append(state.tokens, Token{TokenNumber, state.source[begin:end], state.lineNumber})
	state.position = end - 1
	return nil
}

func scanSpecial(state *state) (err error) {

	specials := map[string]TokenType{
		"+":  TokenPlus,
		"-":  TokenMinus,
		"*":  TokenTimes,
		"/":  TokenDivide,
		"<":  TokenLessThan,
		"<=": TokenLessThanOrEqualTo,
		">":  TokenGreaterThan,
		">=": TokenGreaterThanOrEqualTo,
		"==": TokenEquals,
		"!=": TokenNotEquals,
		"=":  TokenAssign,
		";":  TokenEndOfStatement,
		",":  TokenListSeparator,
		"(":  TokenOpenList,
		")":  TokenCloseList,
		"[":  TokenOpenIndex,
		"]":  TokenCloseIndex,
		"{":  TokenOpenBlock,
		"}":  TokenCloseBlock,
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
