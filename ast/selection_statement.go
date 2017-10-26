package ast

import (
	"github.com/tcort/cmc/scanner"
)

type NodeSelectionStatementType string

const (
	NodeSelectionStatementIf     = NodeSelectionStatementType("IF")
	NodeSelectionStatementIfElse = NodeSelectionStatementType("IF-ELSE")
)

type NodeSelectionStatement struct {
	NodeType      NodeSelectionStatementType
	If            scanner.Token
	OpenList      scanner.Token
	Expression    *NodeExpression
	CloseList     scanner.Token
	StatementIf   *NodeStatement
	Else          scanner.Token
	StatementElse *NodeStatement
}
