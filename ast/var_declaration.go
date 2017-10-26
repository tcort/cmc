package ast

import (
	"github.com/tcort/cmc/scanner"
)

type NodeVarDeclarationType string

const (
	NodeVarDeclarationTypeScalar = NodeVarDeclarationType("SCALAR")
	NodeVarDeclarationTypeArray  = NodeVarDeclarationType("ARRAY")
)

type NodeVarDeclaration struct {
	NodeType       NodeVarDeclarationType
	TypeSpecifier  *NodeTypeSpecifier
	Identifier     scanner.Token
	OpenIndex      scanner.Token
	Number         scanner.Token
	CloseIndex     scanner.Token
	EndOfStatement scanner.Token
}

func (varDeclaration *NodeVarDeclaration) String() string {
	if varDeclaration.NodeType == NodeVarDeclarationTypeScalar {
		return varDeclaration.TypeSpecifier.String() + " " + varDeclaration.Identifier.String() + varDeclaration.EndOfStatement.String()
	} else {
		return varDeclaration.TypeSpecifier.String() + " " + varDeclaration.Identifier.String() + varDeclaration.OpenIndex.String() + varDeclaration.Number.String() + varDeclaration.CloseIndex.String() + varDeclaration.EndOfStatement.String()
	}
}
