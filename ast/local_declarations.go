package ast

type NodeLocalDeclarationsType string

const (
	NodeLocalDeclarationsEmpty    = NodeLocalDeclarationsType("EMPTY")
	NodeLocalDeclarationsNonEmpty = NodeLocalDeclarationsType("NON-EMPTY")
)

type NodeLocalDeclarations struct {
	NodeType          NodeLocalDeclarationsType
	LocalDeclarations *NodeLocalDeclarations
	VarDeclaration    *NodeVarDeclaration
}
