package ast

type NodeDeclarationListType string

const (
	NodeDeclarationListTypeList   = NodeDeclarationListType("LIST")
	NodeDeclarationListTypeSingle = NodeDeclarationListType("SINGLE")
)

type NodeDeclarationList struct {
	NodeType        NodeDeclarationListType
	DeclarationList *NodeDeclarationList
	Declaration     *NodeDeclaration
}

func (declarationList *NodeDeclarationList) String() string {

	var s string

	for dl := declarationList; dl != nil; dl = dl.DeclarationList {
		s += dl.Declaration.String() + "\n"
	}

	return s
}
