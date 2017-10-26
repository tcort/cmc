package parser

import (
	"github.com/tcort/cmc/ast"
	"github.com/tcort/cmc/errors"
	"github.com/tcort/cmc/scanner"
)

func parseTypeSpecifier(state *state) (typeSpecifier *ast.NodeTypeSpecifier, err error) {
	typeSpecifier = new(ast.NodeTypeSpecifier)

	if state.index >= len(state.tokens) {
		return nil, errors.New("not enough tokens", state.filename, state.tokens[len(state.tokens)-1].LineNumber)
	} else if state.tokens[state.index].Type == scanner.TokenVoid {
		typeSpecifier.NodeType = ast.NodeTypeSpecifierTypeVoid
		typeSpecifier.Type = state.tokens[state.index]
		state.index++
		return typeSpecifier, nil
	} else if state.tokens[state.index].Type == scanner.TokenInt {
		typeSpecifier.NodeType = ast.NodeTypeSpecifierTypeInt
		typeSpecifier.Type = state.tokens[state.index]
		state.index++
		return typeSpecifier, nil
	}
	return nil, errors.New("unexpected token", state.filename, state.tokens[state.index].LineNumber)
}
