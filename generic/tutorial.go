package generic

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"
	"testing"
)

import (
	"golang.org/x/exp/constraints"
)

type Integer interface {
	int | int8 | int16 | int32 | int64
}

func min[T Integer](a, b T) T {
	if a < b {
		return a
	}

	return b
}

func min2[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}

	return b
}

func isSame[T comparable](a, b T) bool {
	if a == b {
		return a == b
	} else if a != b {
		return b != a
	}

	return false
}

type Stringer interface {
	Integer
	String() string
}

func Print[T Integer](a T) T {

	return a
}

func String[T Integer](a T) T {

	return a
}

func TestA(t *testing.T) {
	m1 := min(3, 5)

	fmt.Println(m1)

	m2 := min2(3, 5)

	fmt.Println(m2)

	m3 := isSame(3, 5)

	fmt.Println(m3)

	s1 := String(4)

	fmt.Println(s1)

	var ms1 MyString = "Hello"
	var ms2 MyString = "World"

	fmt.Println(Equal(ms1, ms2))

	mapTest()

}

type ComparableHasher interface {
	comparable // ==, !=
	Hash() uint32
}

type MyString string

func (s MyString) Hash() uint32 {
	h := fnv.New32()
	h.Write([]byte(s))
	return h.Sum32()
}

func Equal[T ComparableHasher](a, b T) bool {
	if a == b {
		return true
	}
	return a.Hash() == b.Hash()
}

func Map[F, T any](s []F, f func(F) T) []T {
	rst := make([]T, len(s))

	for i, v := range s {
		rst[i] = f(v)
	}

	return rst
}

func mapTest() {
	doubled := Map([]int{1, 2, 3}, func(v int) int {
		return v * 2
	})

	fmt.Println(doubled)

	uppered := Map([]string{"hello", "world", "abc"}, func(v string) string {
		return strings.ToUpper(v)
	})

	fmt.Println(uppered)

	toString := Map([]int{1, 2, 3}, func(v int) string {
		return "str" + strconv.Itoa(v)

	})

	fmt.Println(toString)
}

type Node[T any] struct {
	val  T
	next *Node[T]
}

func (n *Node[T]) Push(v T) *Node[T] {
	node := NewNode(v)
	n.next = node
	return node
}

func NewNode[T any](v T) *Node[T] {
	return &Node[T]{val: v}
}
