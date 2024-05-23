package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
	tree "github.com/untref-ayp2/data-structures/tree"
)

/* Construct the following tree
          15
        /    \
       /      \
      10       20
     / \      /  \
    /   \    /    \
   8    12  16    25
*/

func TestSecLargeElem(t *testing.T) {
	bst := tree.NewBinarySearchTree[int]()
	bst.Insert(15)
	bst.Insert(10)
	bst.Insert(20)
	bst.Insert(8)
	bst.Insert(12)
	bst.Insert(16)
	bst.Insert(25)

	p, _ := SecondLargestElement(bst)
	assert.Equal(t, 20, p)
}

func TestNilSecLargeElem(t *testing.T) {
	bst := tree.NewBinarySearchTree[int]()

	_, err := SecondLargestElement(bst)
	assert.EqualError(t, err, "No hay valores")
}
