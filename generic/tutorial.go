package generic

import (
	"fmt"
	"golang.org/x/exp/maps"
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

type Addable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | complex64 | complex128 | string
}

func add[T Addable](a, b T) T {
	return a + b
}

//fmt.Println(add(1, 2))
//fmt.Println(add("foo", "bar"))

func MapExample() {
	m := map[int]string{
		123: "foo",
		456: "bar",
	}
	fmt.Println(maps.Keys(m))   // [123 456]
	fmt.Println(maps.Values(m)) // [foo bar]
	m2 := maps.Clone(m)
	fmt.Println(maps.Equal(m2, m)) // true
	m3 := map[int]string{
		789: "baz",
	}
	maps.Copy(m3, m)
	fmt.Println(maps.Equal(m3, m)) // false
}

func some[T any](a []T, f func(T) bool) bool {
	for _, e := range a {
		if f(e) {
			return true
		}
	}

	return false
}

func SomeExample() {
	someEven := some([]int{2, 3, 4, 5, 6}, func(v int) bool {
		return v%2 == 0
	})
	if someEven {
		fmt.Println("Some odd")
	}

	someEven = some([]int{1, 3, 5, 7}, func(v int) bool {
		return v%2 == 0
	})
	if !someEven {
		fmt.Println("Only odd")
	}
}
