package main

import (
	"fmt"
	"strconv"
)

type Nodo struct {
	valor    int
	anterior *Nodo
}

type Stack struct {
	ultimo *Nodo
}

func (nodo *Stack) insertar(valor int) {
	if nodo.ultimo != nil {
		newNode := &Nodo{valor: valor}
		newNode.anterior = nodo.ultimo
		nodo.ultimo = newNode
		return
	}
	nodo.ultimo = &Nodo{valor: valor}
}

func (nodo *Stack) pop() interface{} {
	if nodo.ultimo != nil {
		temporal := nodo.ultimo.valor
		nodo.ultimo = nodo.ultimo.anterior
		return temporal
	}
	return nil
}

func (nodo *Stack) recorrer() string {
	actual := nodo.ultimo
	stack := ""
	for actual != nil {
		stack += " -> " + strconv.Itoa(actual.valor)
		actual = actual.anterior
	}
	return stack
}

func main() {
	l1 := Stack{}
	l1.insertar(1)
	l1.insertar(2)
	l1.insertar(3)
	l1.insertar(4)
	fmt.Println(l1.recorrer())

	l1.pop()
	l1.pop()
	fmt.Println(l1.recorrer())
}
