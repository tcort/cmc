package parser

import (
	"github.com/tcort/cmc/ast"
	"github.com/tcort/cmc/errors"
	"github.com/tcort/cmc/scanner"
)

func parseFunDeclaration(state *state) (funDeclaration *ast.NodeFunDeclaration, err error) {
	begin := state.index
	funDeclaration = new(ast.NodeFunDeclaration)

	funDeclaration.TypeSpecifier, err = parseTypeSpecifier(state)
	if err != nil {
		state.index = begin
		return nil, err
	}

	if state.tokens[state.index].Type != scanner.TokenIdentifier {
		state.index = begin
		return nil, errors.New("unexpected token", state.filename, state.tokens[state.index].LineNumber)
	}

	funDeclaration.Identifier = state.tokens[state.index]
	state.index++

	if state.tokens[state.index].Type != scanner.TokenOpenList {
		state.index = begin
		return nil, errors.New("unexpected token", state.filename, state.tokens[state.index].LineNumber)
	}

	funDeclaration.OpenList = state.tokens[state.index]
	state.index++

	// TODO parseParams

	if state.tokens[state.index].Type != scanner.TokenCloseList {
		state.index = begin
		return nil, errors.New("unexpected token", state.filename, state.tokens[state.index].LineNumber)
	}

	funDeclaration.CloseList = state.tokens[state.index]
	state.index++

	// TODO parseCompoundStatement

	return funDeclaration, nil
}
