package datastructure

type FenwickTree struct {
	n    int
	tree []int
}

func NewFenwickTree(n int) *FenwickTree {
	return &FenwickTree{
		n:    n,
		tree: make([]int, n+1),
	}
}

func (ft *FenwickTree) Update(i int, val int) {
	i += 1
	for i <= ft.n {
		ft.tree[i] += val
		i += i & (-i)
	}
}

func (ft *FenwickTree) Query(i int) int {
	i += 1
	res := 0
	for i > 0 {
		res += ft.tree[i]
		i -= i & (-i)
	}
	return res
}
