package main

import (
	"fmt"
)

type BstNode struct {
	left  *BstNode
	value int
	right *BstNode
}

func (bstNode *BstNode) Add(value int){
	if bstNode == nil{
		return
	}else if value < bstNode.value{
		if(bstNode.left == nil){
			bstNode.left = &BstNode{value: value}
		}else{
			bstNode.left.Add(value)
		}
	}else if value > bstNode.value{
		if(bstNode.right == nil){
			bstNode.right = &BstNode{value: value}
		}else{
			bstNode.right.Add(value)
		}
	}
}

func (bstNode *BstNode) Search(value int) bool{
	if bstNode == nil{
		return false
	}
	if bstNode.value == value{
		return true
	}
	if value <= bstNode.value{
		return bstNode.left.Search(value)
	}else{
		return bstNode.right.Search(value)
	}
}

func (bstNode *BstNode) Min() int{
	if bstNode == nil{
		return -1
	}
	if bstNode.left != nil{
		return bstNode.left.Min()
	}

	return bstNode.value
}

func (bstNode *BstNode) Max() int{
	if bstNode == nil{
		return -1
	}
	if bstNode.right != nil{
		return bstNode.right.Max()
	}
	
	return bstNode.value
}

// raiz esquerda direita
func (bstNode *BstNode) PrintPre(){
	if bstNode == nil{
		return
	}
	fmt.Print(bstNode.value," ")
	bstNode.left.PrintPre()
	bstNode.right.PrintPre()
	
}
// esquerda raiz direita
func (bstNode *BstNode) PrintIn(){
	if bstNode == nil{
		return
	}
	bstNode.left.PrintIn()
	fmt.Print(bstNode.value," ")
	bstNode.right.PrintIn()
}
// esquerda direita raiz
func (bstNode *BstNode) PrintPos(){
	if bstNode == nil{
		return
	}
	bstNode.left.PrintPos()
	bstNode.right.PrintPos()	
	fmt.Print(bstNode.value," ")

}

func (bstNode *BstNode) PrintLevels(){
	if bstNode == nil{
		return
	}
	queue := []*BstNode{bstNode}
	for len(queue) > 0{
		current := queue[0]
		queue = queue[1:]
		fmt.Print(current.value," ")
		
		if current.left != nil{
			queue = append(queue, current.left)
			
		}
		if current.right != nil{
			queue = append(queue, current.right)
		}
	}
}

func (bstNode *BstNode) Height() int{
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

func (bstNode *BstNode) Remove(value int) *BstNode{
	if bstNode == nil {
		return nil
	}

	if value < bstNode.value {
			bstNode.left = bstNode.left.Remove(value)
	} else if value > bstNode.value {
			bstNode.right = bstNode.right.Remove(value)
	} else {
		if bstNode.left == nil && bstNode.right == nil {
			return nil
		}
		if bstNode.left == nil {
				return bstNode.right 
		}
		if bstNode.right == nil {
				return bstNode.left 
		}

		temp := bstNode.right
		for temp.left != nil {
				temp = temp.left
		}

		bstNode.value = temp.value
		bstNode.right = bstNode.right.Remove(temp.value)
		}

		return bstNode
}
func (bstNode *BstNode) isBst() bool{
	if bstNode == nil {
		return true
	}
	if bstNode.left != nil && bstNode.left.value > bstNode.value {
		return false
	}
	if bstNode.right != nil && bstNode.right.value < bstNode.value {
		return false
	}

	if !bstNode.left.isBst() || !bstNode.right.isBst() {
		return false
	}
	return true
}
func (bstNode *BstNode) Size() int{
	if bstNode == nil {
		return 0
	}
	return 1 + bstNode.left.Size() + bstNode.right.Size()
}
func convertToBalancedBst(v []int, ini int, fim int) *BstNode{
	if ini > fim {
		return nil
	}
	mid := (ini + fim) / 2
	node := &BstNode{value: v[mid]}
	node.left = convertToBalancedBst(v, ini, mid-1)
	node.right = convertToBalancedBst(v, mid+1, fim)
	return node
}


func main() {
	var root *BstNode
	root = &BstNode{value: 50}
	root.Add(30)
	root.Add(45)
	root.Add(32)
	root.Add(10)
	root.Add(90)
	root.Add(55)
	root.Add(49)
	root.Add(-5)
	root.Add(88)
	root.Add(110)
	root.Add(40)

	fmt.Println("Procurando valor que nao existe",root.Search(-1))
	fmt.Println("Procurando valor que nao existe",root.Search(49))
    fmt.Println("O valor minimo é: ",root.Min())
	fmt.Println("O valor maximo é: ",root.Max())
	
	fmt.Print("\nPre-ordem: ")
	root.PrintPre()
	fmt.Print("\nIn-ordem:  ")
	root.PrintIn()
	fmt.Print("\nPos-ordem: ")
	root.PrintPos()
	fmt.Print("\nEm nível:  ")
	root.PrintLevels()
	fmt.Println("\nA altura da árvore é: ",root.Height())
	fmt.Println("A BST tem tamanho(antes de remover): ",root.Size())
	root = root.Remove(45)
	fmt.Println("É uma BST: ",root.isBst())
	fmt.Println("A BST tem tamanho(pós remover um elemento): ",root.Size())

	fmt.Println("\n\nConvertendo vetor em BST balanceada")
	vetor := []int{1,2,3,4,5,6,7,8,9,10}
	arvore_balanciada := convertToBalancedBst(vetor,0,len(vetor)-1)
	fmt.Print("Navegaçao em nivel:  ")
	arvore_balanciada.PrintLevels()
}
