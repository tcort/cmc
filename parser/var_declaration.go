package parser

import (
	"github.com/tcort/cmc/ast"
	"github.com/tcort/cmc/errors"
	"github.com/tcort/cmc/scanner"
)

func parseVarDeclaration(state *state) (varDeclaration *ast.NodeVarDeclaration, err error) {
	begin := state.index
	varDeclaration = new(ast.NodeVarDeclaration)

	varDeclaration.TypeSpecifier, err = parseTypeSpecifier(state)
	if err != nil {
		state.index = begin
		return nil, err
	}

	if state.tokens[state.index].Type != scanner.TokenIdentifier {
		state.index = begin
		return nil, errors.New("unexpected token", state.filename, state.tokens[state.index].LineNumber)
	}

	varDeclaration.Identifier = state.tokens[state.index]
	state.index++

	if state.tokens[state.index].Type == scanner.TokenEndOfStatement {
		varDeclaration.NodeType = ast.NodeVarDeclarationTypeScalar
		varDeclaration.EndOfStatement = state.tokens[state.index]
		state.index++
		return varDeclaration, nil
	}

	varDeclaration.NodeType = ast.NodeVarDeclarationTypeArray

	if state.tokens[state.index].Type != scanner.TokenOpenIndex {
		state.index = begin
		return nil, errors.New("unexpected token", state.filename, state.tokens[state.index].LineNumber)
	}
	varDeclaration.OpenIndex = state.tokens[state.index]
	state.index++

	if state.tokens[state.index].Type != scanner.TokenNumber {
		state.index = begin
		return nil, errors.New("unexpected token", state.filename, state.tokens[state.index].LineNumber)
	}
	varDeclaration.Number = state.tokens[state.index]
	state.index++

	if state.tokens[state.index].Type != scanner.TokenCloseIndex {
		state.index = begin
		return nil, errors.New("unexpected token", state.filename, state.tokens[state.index].LineNumber)
	}
	varDeclaration.CloseIndex = state.tokens[state.index]
	state.index++

	if state.tokens[state.index].Type != scanner.TokenEndOfStatement {
		state.index = begin
		return nil, errors.New("unexpected token", state.filename, state.tokens[state.index].LineNumber)
	}
	varDeclaration.EndOfStatement = state.tokens[state.index]
	state.index++

	return varDeclaration, nil
}
