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
