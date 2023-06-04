package main

import (
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	number int
	hight  int
	left   *Node
	right  *Node
}

type AVLTree struct {
	root *Node
}

func (this *AVLTree) insert(number int) {
	this.root = this.insert1(number, this.root)
}

func (this *AVLTree) insert1(number1 int, node *Node) *Node {
	if node == nil {
		return &Node{
			number: number1,
			left:   nil,
			right:  nil,
		}
	}
	if number1 < node.number {
		node.left = this.insert1(number1, node.left)
		if this.getHeight(node.left)-this.getHeight(node.right) == 2 {
			if number1 < node.left.number {
				node = this.rotateLeftChild(node)
			} else {
				node = this.doubleRotateLeftChild(node)
			}
		}
	} else if number1 > node.number {
		node.right = this.insert1(number1, node.right)
		if this.getHeight(node.right)-this.getHeight(node.left) == 2 {
			if number1 > node.right.number {
				node = this.rotateRightChild(node)
			} else {
				node = this.doubleRotateRightChild(node)
			}
		}
	}
	node.hight = this.getMaxHeight(this.getHeight(node.left), this.getHeight(node.right)) + 1
	return node
}

func (this *AVLTree) getMaxHeight(left, right int) int {
	if left > right {
		return left
	}
	return right
}

func (this *AVLTree) getHeight(node *Node) int {
	if node == nil {
		return -1
	}
	return node.hight
}

func (this *AVLTree) rotateLeftChild(node *Node) *Node {
	auxiliar := node.left
	node.left = auxiliar.right
	auxiliar.right = node
	node.hight = this.getMaxHeight(this.getHeight(node.left), this.getHeight(node.right)) + 1
	auxiliar.hight = this.getMaxHeight(this.getHeight(auxiliar.left), node.hight) + 1
	return auxiliar
}

func (this *AVLTree) rotateRightChild(node *Node) *Node {
	auxiliar := node.right
	node.right = auxiliar.left
	auxiliar.left = node
	node.hight = this.getMaxHeight(this.getHeight(node.left), this.getHeight(node.right)) + 1
	auxiliar.hight = this.getMaxHeight(this.getHeight(auxiliar.right), node.hight) + 1
	return auxiliar
}

func (this *AVLTree) doubleRotateLeftChild(node *Node) *Node {
	node.left = this.rotateRightChild(node.left)
	return this.rotateLeftChild(node)
}

func (this *AVLTree) doubleRotateRightChild(node *Node) *Node {
	node.right = this.rotateLeftChild(node.right)
	return this.rotateRightChild(node)
}

func (this *AVLTree) dot() string {
	dot := "digraph Tree{\n\tnode [shape = record];"
	dot += this.dot1(this.root)
	dot += "\n}"
	return dot
}

func (this *AVLTree) dot1(node *Node) string {
	dot := ""
	if node.left == nil && node.right == nil {
		dot = "\n\tnode_" + strconv.Itoa(node.number) + "[label=\"<C3>" + strconv.Itoa(node.number) + "\"];"
	} else {
		dot = "\n\tnode_" + strconv.Itoa(node.number) + "[label=\"<C0> | <C3>" + strconv.Itoa(node.number) + " | <C1>\"];"
	}
	if node.left != nil {
		dot += this.dot1(node.left)
		dot += "\n\tnode_" + strconv.Itoa(node.number) + ":C0 -> node_" + strconv.Itoa(node.left.number) + ":C3;"
	}
	if node.right != nil {
		dot += this.dot1(node.right)
		dot += "\n\tnode_" + strconv.Itoa(node.number) + ":C1 -> node_" + strconv.Itoa(node.right.number) + ":C3;"
	}
	return dot
}

func (this *AVLTree) preorder() {
	this.preorder1(this.root)
	fmt.Println()
}

func (this *AVLTree) preorder1(node *Node) {
	if node != nil {
		fmt.Print(strconv.Itoa(node.number) + " ")
		this.preorder1(node.left)
		this.preorder1(node.right)
	}
}

func (this *AVLTree) inorder() {
	this.inorder1(this.root)
	fmt.Println()
}

func (this *AVLTree) inorder1(node *Node) {
	if node != nil {
		this.inorder1(node.left)
		fmt.Print(strconv.Itoa(node.number) + " ")
		this.inorder1(node.right)
	}
}

func (this *AVLTree) postorder() {
	this.postorder1(this.root)
	fmt.Println()
}

func (this *AVLTree) postorder1(node *Node) {
	if node != nil {
		this.postorder1(node.left)
		this.postorder1(node.right)
		fmt.Print(strconv.Itoa(node.number) + " ")
	}
}

func generarGrafo(dot string) {
	nombreArchivo := "grafo.dot"
	nuevoContenido := dot

	// Eliminar el archivo existente
	err := os.Remove(nombreArchivo)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Error al eliminar el archivo:", err)
		return
	}

	// Crear un nuevo archivo
	file, err := os.Create(nombreArchivo)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()

	// Escribir el nuevo contenido en el archivo
	_, err = file.WriteString(nuevoContenido)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

	fmt.Println("Se ha escrito el nuevo contenido en el archivo.")
}

func main() {
	tree := AVLTree{}
	tree.insert(50)
	tree.insert(17)
	tree.insert(84)
	tree.insert(32)
	tree.insert(26)
	tree.insert(25)
	tree.insert(47)
	tree.insert(95)
	tree.insert(12)
	tree.insert(49)
	tree.insert(100)
	tree.insert(200)
	tree.insert(85)
	tree.insert(210)
	tree.insert(90)
	generarGrafo(tree.dot())
	fmt.Println("PREORDER")
	tree.preorder()
	fmt.Println("\nINORDER")
	tree.inorder()
	fmt.Println("\nPOSTORDER")
	tree.postorder()
}
