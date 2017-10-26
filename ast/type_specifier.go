package ast

import (
	"github.com/tcort/cmc/scanner"
)

type NodeTypeSpecifierType string

const (
	NodeTypeSpecifierTypeInt  = NodeTypeSpecifierType("INT")
	NodeTypeSpecifierTypeVoid = NodeTypeSpecifierType("VOID")
)

type NodeTypeSpecifier struct {
	NodeType NodeTypeSpecifierType
	Type     scanner.Token
}

func (typeSpecifier *NodeTypeSpecifier) String() string {
	return typeSpecifier.Type.String()
}
