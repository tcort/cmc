package ast

type NodeArgsType string

const (
	NodeArgsTypeEmpty    = NodeArgsType("EMPTY")
	NodeArgsTypeNonEmpty = NodeArgsType("NON-EMPTY")
)

type NodeArgs struct {
	NodeType NodeArgsType
	ArgsList *NodeArgsList
}
