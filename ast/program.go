package ast

type NodeProgram struct {
	DeclarationList *NodeDeclarationList
}

func (program NodeProgram) String() string {
	return program.DeclarationList.String()
}
