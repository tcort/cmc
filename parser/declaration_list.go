package parser

import "github.com/tcort/cmc/ast"

func parseDeclarationList(state *state) (declarationList *ast.NodeDeclarationList, err error) {
	begin := state.index
	declarationList = new(ast.NodeDeclarationList)

	declarationList.Declaration, err = parseDeclaration(state)
	if err != nil {
		state.index = begin
		return nil, err
	}

	if state.index < len(state.tokens) {
		declarationList.DeclarationList, err = parseDeclarationList(state)
		if err != nil {
			state.index = begin
			return nil, err
		}
	}

	return declarationList, nil
}
