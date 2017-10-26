package ast

import (
	"github.com/tcort/cmc/scanner"
)

type NodeArgsListType string

const (
	NodeArgsListTypeList   = NodeArgsListType("LIST")
	NodeArgsListTypeSingle = NodeArgsListType("SINGLE")
)

type NodeArgsList struct {
	NodeType      NodeArgsListType
	ArgsList      *NodeArgsList
	ListSeparator scanner.Token
	Expression    *NodeExpression
}
