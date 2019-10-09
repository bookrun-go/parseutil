package tree

type MultiLeafTree struct {
	Parent   *MultiLeafTree
	Children []*MultiLeafTree
	Data     interface{}
}

func NewMultiLeafTree(parent *MultiLeafTree, data interface{}) *MultiLeafTree {
	return &MultiLeafTree{
		Parent: parent,
		Data:   data,
	}
}

func (t *MultiLeafTree) AddLeaf(data interface{}) *MultiLeafTree {
	leafTree := NewMultiLeafTree(t, data)

	if t.Children == nil {
		t.Children = make([]*MultiLeafTree, 0)
	}
	t.Children = append(t.Children, leafTree)
	return leafTree
}
