package main

import (
	"fmt"
	"strconv"
)

type Nodo struct {
	valor     int
	siguiente *Nodo
	anterior  *Nodo
}

type ListaDoble struct {
	primero *Nodo
	ultimo  *Nodo
}

func (nodo *ListaDoble) insertar(valor int) {
	if nodo.primero != nil {
		nodo.ultimo.siguiente = &Nodo{valor: valor}
		nodo.ultimo.siguiente.anterior = nodo.ultimo
		nodo.ultimo = nodo.ultimo.siguiente
		return
	}
	nodo.primero = &Nodo{valor: valor}
	nodo.ultimo = nodo.primero
}

func (nodo *ListaDoble) recorrer1() string {
	actual := nodo.primero
	lista := ""
	for actual != nil {
		lista += " -> " + strconv.Itoa(actual.valor)
		actual = actual.siguiente
	}
	return lista
}

func (nodo *ListaDoble) recorrer2() string {
	actual := nodo.ultimo
	lista := ""
	for actual != nil {
		lista += " -> " + strconv.Itoa(actual.valor)
		actual = actual.anterior
	}
	return lista
}

func main() {
	l1 := ListaDoble{}
	l1.insertar(1)
	l1.insertar(2)
	l1.insertar(3)
	l1.insertar(4)
	fmt.Println(l1.recorrer1())
	fmt.Println(l1.recorrer2())
}
