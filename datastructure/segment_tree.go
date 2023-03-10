package datastructure

type SegmentTree struct {
	Tree []int
}

func NewSegmentTree(arr []int) *SegmentTree {
	st := &SegmentTree{
		Tree: make([]int, 4*len(arr)),
	}
	st.build(arr, 0, 0, len(arr)-1)
	return st
}

func (st *SegmentTree) build(arr []int, treeIdx, leftIdx, rightIdx int) {
	if leftIdx == rightIdx {
		st.Tree[treeIdx] = arr[leftIdx]
	} else {
		midIdx := (leftIdx + rightIdx) / 2
		st.build(arr, 2*treeIdx+1, leftIdx, midIdx)
		st.build(arr, 2*treeIdx+2, midIdx+1, rightIdx)
		st.Tree[treeIdx] = st.Tree[2*treeIdx+1] + st.Tree[2*treeIdx+2]
	}
}

func (st *SegmentTree) Query(treeIdx, leftIdx, rightIdx, qLeftIdx, qRightIdx int) int {
	if qLeftIdx <= leftIdx && rightIdx <= qRightIdx {
		return st.Tree[treeIdx]
	} else if rightIdx < qLeftIdx || qRightIdx < leftIdx {
		return 0
	} else {
		midIdx := (leftIdx + rightIdx) / 2
		leftSum := st.Query(2*treeIdx+1, leftIdx, midIdx, qLeftIdx, qRightIdx)
		rightSum := st.Query(2*treeIdx+2, midIdx+1, rightIdx, qLeftIdx, qRightIdx)
		return leftSum + rightSum
	}
}

func (st *SegmentTree) Update(treeIdx, leftIdx, rightIdx, arrIdx, newVal int) {
	if leftIdx == rightIdx {
		st.Tree[treeIdx] = newVal
	} else {
		midIdx := (leftIdx + rightIdx) / 2
		if arrIdx <= midIdx {
			st.Update(2*treeIdx+1, leftIdx, midIdx, arrIdx, newVal)
		} else {
			st.Update(2*treeIdx+2, midIdx+1, rightIdx, arrIdx, newVal)
		}
		st.Tree[treeIdx] = st.Tree[2*treeIdx+1] + st.Tree[2*treeIdx+2]
	}
}
