package ast

import (
	"github.com/tcort/cmc/scanner"
)

type NodeCall struct {
	Identifier scanner.Token
	OpenList   scanner.Token
	Args       *NodeArgs
	CloseList  scanner.Token
}
