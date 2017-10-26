package ast

import (
	"github.com/tcort/cmc/scanner"
)

type NodeExpressionType string

const (
	NodeExpressionTypeAssignment = NodeReturnStatementType("ASSIGN")
	NodeExpressionTypeSimple     = NodeReturnStatementType("SIMPLE")
)

type NodeExpression struct {
	NodeType         NodeExpressionType
	Var              *NodeVar
	Assign           scanner.Token
	Expression       *NodeExpression
	SimpleExpression *NodeSimpleExpression
}
