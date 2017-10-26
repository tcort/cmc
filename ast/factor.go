package ast

import (
	"github.com/tcort/cmc/scanner"
)

type NodeFactorType string

const (
	NodeFactorTypeExpression = NodeFactorType("EXPRESSION")
	NodeFactorTypeVar        = NodeFactorType("VAR")
	NodeFactorTypeCall       = NodeFactorType("CALL")
	NodeFactorTypeNumber     = NodeFactorType("Number")
)

type NodeFactor struct {
	NodeType   NodeFactorType
	OpenList   scanner.Token
	Expression *NodeExpression
	CloseList  scanner.Token
	Var        *NodeVar
	Call       *NodeCall
	Number     scanner.Token
}
