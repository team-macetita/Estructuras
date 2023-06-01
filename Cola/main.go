package main

import (
	"fmt"
	"strconv"
)

type Nodo struct {
	valor     int
	siguiente *Nodo
}

type Queue struct {
	primero *Nodo
	ultimo  *Nodo
}

func (nodo *Queue) insertar(valor int) {
	if nodo.primero != nil {
		nodo.ultimo.siguiente = &Nodo{valor: valor}
		nodo.ultimo = nodo.ultimo.siguiente
		return
	}
	nodo.primero = &Nodo{valor: valor}
	nodo.ultimo = nodo.primero
}

func (nodo *Queue) pop() interface{} {
	if nodo.primero != nil {
		temporal := nodo.primero.valor
		nodo.primero = nodo.primero.siguiente
		return temporal
	}
	return nil
}

func (nodo *Queue) recorrer() string {
	actual := nodo.primero
	queue := ""
	for actual != nil {
		queue += " -> " + strconv.Itoa(actual.valor)
		actual = actual.siguiente
	}
	return queue
}

func main() {
	q1 := Queue{}
	q1.insertar(1)
	q1.insertar(2)
	q1.insertar(3)
	q1.insertar(4)
	fmt.Println(q1.recorrer())

	q1.pop()
	q1.pop()
	fmt.Println(q1.recorrer())
}
