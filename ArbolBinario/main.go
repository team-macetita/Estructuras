package main

import (
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	number int
	left   *Node
	right  *Node
}

type Tree struct {
	root *Node
}

func (node *Tree) insert(number int) {
	node.root = node.insert1(number, node.root)
}

func (node *Tree) insert1(number1 int, nodeNew *Node) *Node {
	if nodeNew == nil {
		return &Node{
			number: number1,
			left:   nil,
			right:  nil,
		}
	}
	if number1 < nodeNew.number {
		nodeNew.left = node.insert1(number1, nodeNew.left)
	} else if number1 > nodeNew.number {
		nodeNew.right = node.insert1(number1, nodeNew.right)
	}
	return nodeNew
}

func (node *Tree) dot() string {
	dot := "digraph Tree{\n\tnode [shape = record];"
	dot += node.dot1(node.root)
	dot += "\n}"
	return dot
}

func (node *Tree) dot1(node1 *Node) string {
	dot := ""
	if node1.left == nil && node1.right == nil {
		dot = "\n\tnode_" + strconv.Itoa(node1.number) + "[label=\"" + strconv.Itoa(node1.number) + "\"];"
	} else {
		dot = "\n\tnode_" + strconv.Itoa(node1.number) + "[label=\"<C0> | " + strconv.Itoa(node1.number) + " | <C1>\"];"
	}
	if node1.left != nil {
		dot += node.dot1(node1.left)
		dot += "\n\tnode_" + strconv.Itoa(node1.number) + ":C0 -> node_" + strconv.Itoa(node1.left.number) + ";"
	}
	if node1.right != nil {
		dot += node.dot1(node1.right)
		dot += "\n\tnode_" + strconv.Itoa(node1.number) + ":C1 -> node_" + strconv.Itoa(node1.right.number) + ";"
	}
	return dot
}

func (node *Tree) preorder() {
	node.preorder1(node.root)
	fmt.Println()
}

func (node *Tree) preorder1(node1 *Node) {
	if node1 != nil {
		fmt.Print(strconv.Itoa(node1.number) + " ")
		node.preorder1(node1.left)
		node.preorder1(node1.right)
	}
}

func (node *Tree) inorder() {
	node.inorder1(node.root)
	fmt.Println()
}

func (node *Tree) inorder1(node1 *Node) {
	if node1 != nil {
		node.inorder1(node1.left)
		fmt.Print(strconv.Itoa(node1.number) + " ")
		node.inorder1(node1.right)
	}
}

func (node *Tree) postorder() {
	node.postorder1(node.root)
	fmt.Println()
}

func (node *Tree) postorder1(node1 *Node) {
	if node1 != nil {
		node.postorder1(node1.left)
		node.postorder1(node1.right)
		fmt.Print(strconv.Itoa(node1.number) + " ")
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
	tree := Tree{}
	tree.insert(50)
	tree.insert(17)
	tree.insert(84)
	tree.insert(32)
	tree.insert(26)
	tree.insert(25)
	tree.insert(47)
	tree.insert(95)
	tree.insert(12)
	generarGrafo(tree.dot())
	fmt.Println("PREORDER")
	tree.preorder()
	fmt.Println("\nINORDER")
	tree.inorder()
	fmt.Println("\nPOSTORDER")
	tree.postorder()
}
