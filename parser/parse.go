package parser

import (
	"github.com/tcort/cmc/ast"
	"github.com/tcort/cmc/errors"
	"github.com/tcort/cmc/scanner"
)

type state struct {
	filename string
	tokens   []scanner.Token
	index    int
}

func Parse(filename string, tokens []scanner.Token) (program *ast.NodeProgram, err error) {
	if len(tokens) == 0 {
		return nil, errors.New("no tokens found", filename, 1)
	}
	return parseProgram(&state{filename, tokens, 0})
}
