package parser

import (
	"github.com/tcort/cmc/ast"
	"github.com/tcort/cmc/errors"
)

func parseDeclaration(state *state) (declaration *ast.NodeDeclaration, err error) {
	begin := state.index
	declaration = new(ast.NodeDeclaration)

	declaration.NodeType = ast.NodeDeclarationTypeVar
	declaration.VarDeclaration, err = parseVarDeclaration(state)
	if err == nil {
		return declaration, nil
	}
	state.index = begin

	declaration.NodeType = ast.NodeDeclarationTypeFun
	declaration.FunDeclaration, err = parseFunDeclaration(state)
	if err == nil {
		return declaration, nil
	}
	state.index = begin

	return nil, errors.New("unexpected token", state.filename, state.tokens[state.index].LineNumber)
}
