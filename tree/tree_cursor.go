package tree

import "errors"

type TreeCursor struct {
	CurrentTree *MultiLeafTree
	Head        *MultiLeafTree
}

func NewCursor(data interface{}) *TreeCursor {
	rootLeaf := NewMultiLeafTree(nil, data)
	return &TreeCursor{
		CurrentTree: rootLeaf,
		Head:        rootLeaf,
	}
}

func (tc *TreeCursor) AddToNewLeaf(data interface{}) {
	newLeaf := tc.CurrentTree.AddLeaf(data)
	tc.CurrentTree = newLeaf
}

func (tc *TreeCursor) ToParent() error {
	if tc.CurrentTree.Parent == nil {
		return errors.New("没有父节点")
	}
	tc.CurrentTree = tc.CurrentTree.Parent
	return nil
}
