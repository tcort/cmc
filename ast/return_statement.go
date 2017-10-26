package ast

import (
	"github.com/tcort/cmc/scanner"
)

type NodeReturnStatementType string

const (
	NodeReturnStatementTypeInt  = NodeReturnStatementType("INT")
	NodeReturnStatementTypeVoid = NodeReturnStatementType("VOID")
)

type NodeReturnStatement struct {
	TokenType      NodeReturnStatementType
	Return         scanner.Token
	Expression     *NodeExpression
	EndOfStatement scanner.Token
}
