package ast

import (
	"github.com/tcort/cmc/scanner"
)

type NodeIterationStatement struct {
	While      scanner.Token
	OpenList   scanner.Token
	Expression *NodeExpression
	CloseList  scanner.Token
	Statement  *NodeStatement
}
