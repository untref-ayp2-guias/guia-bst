package ejercicios

import (
	tree "github.com/untref-ayp2/data-structures/tree"
	"github.com/untref-ayp2/data-structures/types"
)

type TreeSet[T types.Ordered] struct {
	set *tree.BinarySearchTree[T]
}

func NewTreeSet[T types.Ordered](elements ...T) *TreeSet[T] {
	return nil
}

func (ts *TreeSet[T]) Add(elements ...T) {
}

func (ts *TreeSet[T]) Size() int {
	return 0
}

func (ts *TreeSet[T]) Contains(element T) bool {
	return false
}

func (ts *TreeSet[T]) Remove(element T) {
}

func (ts *TreeSet[T]) Values() []T {
	return nil
}

func (ts *TreeSet[T]) String() string {
	return ""
}
