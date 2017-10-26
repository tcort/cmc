package ast

import (
	"github.com/tcort/cmc/scanner"
)

type NodeVarType string

const (
	NodeVarTypeScalar = NodeVarType("SCALAR")
	NodeVarTypeArray  = NodeVarType("ARRAY")
)

type NodeVar struct {
	NodeType   NodeVarType
	Identifier scanner.Token
	OpenIndex  scanner.Token
	Expression *NodeExpression
	CloseIndex scanner.Token
}
