package parser

import (
	"github.com/tcort/cmc/ast"
)

func parseProgram(state *state) (program *ast.NodeProgram, err error) {
	begin := state.index
	program = new(ast.NodeProgram)

	program.DeclarationList, err = parseDeclarationList(state)
	if err != nil {
		state.index = begin
		return nil, err
	}

	return program, nil
}
