package ast

type NodeStatementType string

const (
	NodeStatementTypeExpression = NodeStatementType("EXPRESSION")
	NodeStatementTypeCompound   = NodeStatementType("COMPOUND")
	NodeStatementTypeSelection  = NodeStatementType("SELECTION")
	NodeStatementTypeIteration  = NodeStatementType("ITERATION")
	NodeStatementTypeReturn     = NodeStatementType("RETURN")
)

type NodeStatement struct {
	NodeType            NodeStatementType
	ExpressionStatement NodeExpressionStatement
	CompoundStatement   NodeCompoundStatement
	SelectionStatement  NodeSelectionStatement
	IterationStatement  NodeIterationStatement
	ReturnStatement     NodeReturnStatement
}
