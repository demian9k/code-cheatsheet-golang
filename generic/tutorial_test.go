package generic

import (
	"fmt"
	"golang.org/x/exp/slices"
	"testing"
)

func TestNode(t *testing.T) {

	node1 := NewNode(1)
	node1.Push(2).Push(3).Push(4)

	for node1 != nil {
		fmt.Print(node1.val, " - ")
		node1 = node1.next
	}
	fmt.Println()

	fmt.Println()

	node2 := NewNode("Hi")
	node2.Push("hello").Push("How are you")

	for node2 != nil {
		fmt.Print(node2.val, " - ")
		node2 = node2.next
	}
}

func TestSlice1(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}

	contains := slices.Contains(a, 5)
	fmt.Printf("slices.Cotains => %t", contains)

}

func Test_BinarySearch(t *testing.T) {

	//constraints.Ordered
	var ordered = []int{1, 2, 3, 4, 5, 6, 7}

	//인덱스 번호가 반환된다
	idx, _ := slices.BinarySearch(ordered, 6)

	fmt.Println(idx) //5

	var orderedStrings = []string{"a", "b", "c", "d", "e", "f", "g"}

	idx2, _ := slices.BinarySearch(orderedStrings, "c")

	fmt.Println(idx2) //2
}
func Test_BinarySearchFunc(t *testing.T) {

}

// Unused Capacity from slice
func Test_Clip(t *testing.T) {

	src := []string{"a", "b", "c", "d", "e", "f", "g"}

	var strs = make([]string, 10)

	for i, s := range src {
		strs[i] = s
	}

	fmt.Printf("strs => %s,  capacity => %d  length => %d \n", strs, cap(strs), len(strs))

	strs = strs[0:3]

	cliped := slices.Clip(strs)

	fmt.Printf("cliped => %s,  capacity => %d length => %d \n", cliped, cap(cliped), len(cliped))

}
func Test_Clone(t *testing.T) {

}
func Test_Compact(t *testing.T) {
	list := []string{"a", "b", "a", "b", "c", "d"}

	intList := []int{1, 2, 3, 4, 5, 1, 4, 7}

	slices.Compact(list)
	fmt.Println(list)

	slices.Compact(intList)

	fmt.Printf("intList ==> %v \n", intList)

	intList2 := []int{1, 1, 2, 2, 3, 3, 4, 7}

	slices.Compact(intList2)
	fmt.Printf("intList2 ==> %v \n", intList2)

	intList3 := []int{1, 1, 1, 3, 3, 3, 4, 4, 8}

	slices.Compact(intList3)
	fmt.Printf("intList3 ==> %v \n", intList3)
}

func Test_CompactFunc(t *testing.T) {}
func Test_Compare(t *testing.T)     {}
func Test_CompareFunc(t *testing.T) {}
func Test_Contains(t *testing.T)    {}
func Test_Delete(t *testing.T)      {}
func Test_Equal(t *testing.T)       {}
func Test_EqualFunc(t *testing.T)   {}
func Test_Grow(t *testing.T)        {}
func Test_Index(t *testing.T) {

	list := []string{"a", "b", "c", "d", "e", "f", "g"}

	//list 에서 "g" 의 인덱스를 구한다.
	idx := slices.Index(list, "g")

	fmt.Println(idx) // 6
}
func Test_IndexFunc(t *testing.T)      {}
func Test_Insert(t *testing.T)         {}
func Test_IsSorted(t *testing.T)       {}
func Test_IsSortedFunc(t *testing.T)   {}
func Test_Sort(t *testing.T)           {}
func Test_SortFunc(t *testing.T)       {}
func Test_SortStableFunc(t *testing.T) {}
