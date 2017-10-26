package ast

import (
	"github.com/tcort/cmc/scanner"
)

type NodeParamType string

const (
	NodeParamTypeScalar = NodeParamType("SCALAR")
	NodeParamTypeArray  = NodeParamType("ARRAY")
)

type NodeParam struct {
	NodeType      NodeParamType
	TypeSpecifier *NodeTypeSpecifier
	Identifier    scanner.Token
	OpenIndex     scanner.Token
	CloseIndex    scanner.Token
}
