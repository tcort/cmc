package ast

type NodeDeclarationType string

const (
	NodeDeclarationTypeVar = NodeDeclarationType("VAR")
	NodeDeclarationTypeFun = NodeDeclarationType("FUN")
)

type NodeDeclaration struct {
	NodeType       NodeDeclarationType
	VarDeclaration *NodeVarDeclaration
	FunDeclaration *NodeFunDeclaration
}

func (declaration *NodeDeclaration) String() string {
	if declaration.NodeType == NodeDeclarationTypeVar {
		return declaration.VarDeclaration.String()
	} else {
		return declaration.FunDeclaration.String()
	}
}
