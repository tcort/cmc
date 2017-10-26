package ast

type NodeStatementListType string

const (
	NodeStatementListTypeEmpty    = NodeStatementListType("EMPTY")
	NodeStatementListTypeNonEmpty = NodeStatementListType("NON-EMPTY")
)

type NodeStatementList struct {
	NodeType      NodeStatementListType
	StatementList *NodeStatementList
	Statement     *NodeStatement
}
