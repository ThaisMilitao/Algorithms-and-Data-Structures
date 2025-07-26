package main

import (
	"fmt"
)

type BstNode struct {
	left  *BstNode
	value int
	right *BstNode

	height int
	bf     int
}

func (bstNode *BstNode) Search(value int) bool {
	if bstNode == nil {
		return false
	}
	if bstNode.value == value {
		return true
	}
	if value <= bstNode.value {
		return bstNode.left.Search(value)
	} else {
		return bstNode.right.Search(value)
	}
}

func (bstNode *BstNode) Min() int {
	if bstNode == nil {
		return -1
	}
	if bstNode.left != nil {
		return bstNode.left.Min()
	}

	return bstNode.value
}

func (bstNode *BstNode) Max() int {
	if bstNode == nil {
		return -1
	}
	if bstNode.right != nil {
		return bstNode.right.Max()
	}

	return bstNode.value
}

func (bstNode *BstNode) PrintLevels() {
	if bstNode == nil {
		return
	}
	queue := []*BstNode{bstNode}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Print(current.value, " ")

		if current.left != nil {
			queue = append(queue, current.left)

		}
		if current.right != nil {
			queue = append(queue, current.right)
		}
	}
}

func (bstNode *BstNode) Height() int {
	if bstNode == nil {
		return -1
	}

	leftHeight := bstNode.left.Height()
	rightHeight := bstNode.right.Height()

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func (bstNode *BstNode) Size() int {
	if bstNode == nil {
		return 0
	}
	return 1 + bstNode.left.Size() + bstNode.right.Size()
}

func (bstNode *BstNode) Add(value int) *BstNode {
	if bstNode == nil {
		bstNode = &BstNode{value: value}
		bstNode.updateProperties()
		return bstNode
	}

	if value <= bstNode.value {
		if bstNode.left == nil {
			bstNode.left = &BstNode{value: value}
		} else {
			bstNode.left = bstNode.left.Add(value)
		}
	} else {
		if bstNode.right == nil {
			bstNode.right = &BstNode{value: value}
		} else {
			bstNode.right = bstNode.right.Add(value)
		}
	}
	bstNode.updateProperties()
	return bstNode.Rebalance()
}

func (bstNode *BstNode) Remove(value int) *BstNode {
	if bstNode == nil {
		return nil
	}

	if bstNode.value == value {
		if bstNode.left == nil && bstNode.right == nil {
			return nil
		} else if bstNode.left == nil && bstNode.right != nil {
			return bstNode.right
		} else if bstNode.left != nil && bstNode.right == nil {
			return bstNode.left
		} else {
			maxLeft := bstNode.left.Max()
			bstNode.value = maxLeft
			bstNode.left.Remove(maxLeft)
		}
	} else if value < bstNode.value {
		if bstNode.left != nil {
			bstNode.left = bstNode.left.Remove(value)
		}
	} else {
		if bstNode.right != nil {
			bstNode.right = bstNode.right.Remove(value)
		}
	}

	bstNode.updateProperties()
	return bstNode.Rebalance()
}

func (bstNode *BstNode) updateProperties() {
	hl := 0
	hr := 0
	if bstNode.left != nil {
		hl = bstNode.left.height
	}
	if bstNode.right != nil {
		hr = bstNode.right.height
	}
	bstNode.bf = hr - hl
	if bstNode.left == nil && bstNode.right == nil {
		bstNode.height = 0
	} else if hl > hr {
		bstNode.height = hl + 1
	} else {
		bstNode.height = hr + 1
	}
}

func (bstNode *BstNode) rotateRight() *BstNode {
	newRoot := bstNode.left
	bstNode.left = newRoot.right
	newRoot.right = bstNode

	bstNode.updateProperties()
	newRoot.updateProperties()

	return newRoot
}

func (bstNode *BstNode) rotateLeft() *BstNode {
	newRoot := bstNode.right
	bstNode.right = newRoot.left
	newRoot.left = bstNode
	bstNode.updateProperties()
	newRoot.updateProperties()

	return newRoot
}

// Case 1: Rebalance Left Left
// O balanciamento consiste em rotacionar a raiz para a direita
func (bstNode *BstNode) RebalanceLeftLeft() *BstNode {
	return bstNode.rotateRight()
}

// Case 2: Rebalance Left Right
// Para balancear fazemos a rotação a esquerda da nova raiz e logo após a rotação a direita da antiga raiz.
func (bstNode *BstNode) RebalanceLeftRight() *BstNode {
	bstNode.left = bstNode.left.rotateLeft()
	return bstNode.rotateRight()
}

// Case 3: Rebalance Right Right
// o balanceamento consiste em rotacionar a raiz para a esquerda
func (bstNode *BstNode) RebalanceRightRight() *BstNode {
	return bstNode.rotateLeft()
}

// Case 4: Rebalance Right Left
// Para balancear fazemos a rotação a direita da nova raiz e logo após a rotação a esquerda da antiga raiz.
func (bstNode *BstNode) RebalanceRightLeft() *BstNode {
	bstNode.right = bstNode.right.rotateRight()
	return bstNode.rotateLeft()
}

func (bstNode *BstNode) Rebalance() *BstNode {
	if bstNode.bf == -2 { //desbalanceada à esquerda
		if bstNode.left.bf == -1 || bstNode.left.bf == 0 {
			// print(bstNode.RebalanceLeftLeft().value)
			return bstNode.RebalanceLeftLeft()
		} else if bstNode.left.bf == 1 { // esq-dir
			return bstNode.RebalanceLeftRight()
		}
	} else if bstNode.bf == 2 { //desbalanceada à direita
		if bstNode.right.bf == 1 || bstNode.right.bf == 0 {
			return bstNode.RebalanceRightRight()
		} else if bstNode.right.bf == -1 { //dir-esq
			return bstNode.RebalanceRightLeft()
		}
	}
	return bstNode
}

func (root *BstNode) IsAVL() bool {
	if root == nil {
		return true
	}
	if root.bf < -1 || root.bf > 1 {
		return false
	}
	return root.left.IsAVL() && root.right.IsAVL()
}

func main() {
	var root *BstNode
	root = root.Add(20)
	root = root.Add(8)
	root = root.Add(1)
	root = root.Add(9)
	root = root.Add(19)
	root = root.Add(11)

	// fmt.Print("\nEm nível:  ")
	root.PrintLevels()

	// fmt.Println("É uma AVL: ",root.IsAVL())

}
