package ast

import (
	"github.com/tcort/cmc/scanner"
)

type NodeExpressionStatementType string

const (
	NodeExpressionStatementTypeEmpty    = NodeExpressionStatementType("EMPTY")
	NodeExpressionStatementTypeNonEmpty = NodeExpressionStatementType("NON-EMPTY")
)

type NodeExpressionStatement struct {
	NodeType       NodeExpressionStatementType
	Expression     *NodeExpression
	EndOfStatement scanner.Token
}
