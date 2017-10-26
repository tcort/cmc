package ast

import (
	"github.com/tcort/cmc/scanner"
)

type NodeFunDeclaration struct {
	TypeSpecifier     *NodeTypeSpecifier
	Identifier        scanner.Token
	OpenList          scanner.Token
	Params            *NodeParams
	CloseList         scanner.Token
	CompoundStatement *NodeCompoundStatement
}

func (funDeclaration *NodeFunDeclaration) String() string {
	return "Fun"
}
