package ast

type NodeSimpleExpressionType string

const (
	NodeStatementListTypeRel    = NodeSimpleExpressionType("REL")
	NodeStatementListTypeSimple = NodeSimpleExpressionType("SIMPLE")
)

type NodeSimpleExpression struct {
	NodeType              NodeSimpleExpressionType
	AdditiveExpression    *NodeAdditiveExpression
	Relop                 *NodeRelop
	AdditiveExpressionRel *NodeAdditiveExpression
}
