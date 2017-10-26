package ast

type NodeAdditiveExpressionType string

const (
	NodeAdditiveExpressionTypeAdd    = NodeAdditiveExpressionType("ADD")
	NodeAdditiveExpressionTypeSimple = NodeAdditiveExpressionType("SIMPLE")
)

type NodeAdditiveExpression struct {
	NodeType           NodeAdditiveExpressionType
	AdditiveExpression *NodeAdditiveExpression
	Addop              *NodeAddop
	Term               *NodeTerm
}
