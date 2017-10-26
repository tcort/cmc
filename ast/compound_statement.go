package ast

import (
	"github.com/tcort/cmc/scanner"
)

type NodeCompoundStatementType string

const (
	NodeCompoundStatementTypeEmpty    = NodeCompoundStatementType("EMPTY")
	NodeCompoundStatementTypeNonEmpty = NodeCompoundStatementType("NON-EMPTY")
)

type NodeCompoundStatement struct {
	NodeType          NodeCompoundStatementType
	OpenBlock         scanner.Token
	LocalDeclarations *NodeLocalDeclarations
	StatementList     *NodeStatementList
	CloseBlock        scanner.Token
}
