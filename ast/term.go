package ast

type NodeTermType string

const (
	NodeTermTypeAdd    = NodeTermType("MUL")
	NodeTermTypeSimple = NodeTermType("SIMPLE")
)

type NodeTerm struct {
	NodeType NodeTermType
	Term     *NodeTerm
	Mulop    *NodeMulop
	Factor   *NodeFactor
}
