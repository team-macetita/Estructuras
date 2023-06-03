package main

import (
	"fmt"
	"os"
	"strconv"
)

type Nodo struct {
	valor     int
	siguiente *Nodo
}

type Queue struct {
	primero  *Nodo
	ultimo   *Nodo
	longitud int
}

func (nodo *Queue) insertar(valor int) {
	if nodo.primero != nil {
		nodo.ultimo.siguiente = &Nodo{valor: valor}
		nodo.ultimo = nodo.ultimo.siguiente
		nodo.longitud++
		return
	}
	nodo.primero = &Nodo{valor: valor}
	nodo.ultimo = nodo.primero
	nodo.longitud++
}

func (nodo *Queue) pop() interface{} {
	if nodo.primero != nil {
		temporal := nodo.primero.valor
		nodo.primero = nodo.primero.siguiente
		nodo.longitud--
		return temporal
	}
	return nil
}

func (nodo *Queue) dot() string {
	dot := "digraph G {\n"
	dot += "rankdir=LR;\n"
	dot += "node[shape=\"box\"];\n"

	actual := nodo.primero
	long := 1
	for actual != nil {
		if long < nodo.longitud {
			dot += strconv.Itoa(actual.valor) + " -> "
			long++
		} else {
			dot += strconv.Itoa(actual.valor)
		}
		if actual.siguiente == nil {
			dot += ";\n"
		}
		actual = actual.siguiente
	}
	dot += "}\n"

	return dot
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
	q1 := Queue{}
	q1.insertar(1)
	q1.insertar(2)
	q1.insertar(3)
	q1.insertar(4)
	generarGrafo(q1.dot())
}
