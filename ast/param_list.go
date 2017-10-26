package ast

import (
	"github.com/tcort/cmc/scanner"
)

type NodeParamListType string

const (
	NodeParamListTypeList   = NodeParamType("LIST")
	NodeParamListTypeSingle = NodeParamType("SINGLE")
)

type NodeParamList struct {
	NodeType      NodeParamListType
	ParamList     *NodeParamList
	ListSeparator scanner.Token
	Param         *NodeParam
}
