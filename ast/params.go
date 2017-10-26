package ast

import (
	"github.com/tcort/cmc/scanner"
)

type NodeParamsType string

const (
	NodeParamsTypeVoid = NodeParamsType("VOID")
	NodeParamsTypeList = NodeParamsType("LIST")
)

type NodeParams struct {
	NodeType  NodeParamsType
	Type      scanner.Token
	ParamList *NodeParamList
}
