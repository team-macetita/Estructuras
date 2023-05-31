package main

import (
	"fmt"
	"strconv"
)

type Nodo struct {
	valor     int
	siguiente *Nodo
}

type ListaEnlazada struct {
	primero *Nodo
	ultimo  *Nodo
}

func (nodo *ListaEnlazada) insertar(valor int) {
	if nodo.primero != nil {
		nodo.ultimo.siguiente = &Nodo{valor: valor}
		nodo.ultimo = nodo.ultimo.siguiente
		return
	}
	nodo.primero = &Nodo{valor: valor}
	nodo.ultimo = nodo.primero
}

func (nodo *ListaEnlazada) recorrer() string {
	actual := nodo.primero
	lista := ""
	for actual != nil {
		lista += " -> " + strconv.Itoa(actual.valor)
		actual = actual.siguiente
	}
	return lista
}

func main() {
	fmt.Println("Lista Simple")
	l1 := ListaEnlazada{}
	l1.insertar(1)
	l1.insertar(2)
	l1.insertar(3)
	l1.insertar(4)
	fmt.Println(l1.recorrer())
}
