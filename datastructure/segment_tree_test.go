package datastructure

import (
	"fmt"
	"testing"
)

func TestSegmentTree_Build(t *testing.T) {
	st := NewSegmentTree([]int{1, 2, 3, 4, 5})
	fmt.Println(st.Tree)
	treeExpected := []int{15, 6, 9, 3, 3, 4, 5, 1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if !equalSlice(treeExpected, st.Tree) {
		t.Errorf("expected %v but got %v", treeExpected, st.Tree)
	}
}

func TestSegmentTree_Update(t *testing.T) {
	st := NewSegmentTree([]int{1, 2, 3, 4, 5})
	tree := st.Tree
	st.Update(0, 0, 5, 2, 5)
	fmt.Println("tree updated by update()")
	fmt.Println(tree)
	updatedTree := []int{17, 8, 9, 3, 5, 4, 5, 1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if !equalSlice(updatedTree, tree) {
		t.Errorf("expected %v but got %v", updatedTree, tree)
	}
}

func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
