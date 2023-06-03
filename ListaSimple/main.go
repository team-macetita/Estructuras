package main

import (
	"fmt"
	"os"
	"paquetes/estudiante"
)

type Nodo struct {
	valor     *estudiante.Estudiante
	siguiente *Nodo
}

type ListaEnlazada struct {
	primero  *Nodo
	ultimo   *Nodo
	longitud int
}

func (nodo *ListaEnlazada) insertar(valor *estudiante.Estudiante) {
	if nodo.primero != nil {
		nodo.ultimo.siguiente = &Nodo{valor: valor}
		nodo.ultimo = nodo.ultimo.siguiente
		nodo.longitud += 1
		return
	}
	nodo.primero = &Nodo{valor: valor}
	nodo.ultimo = nodo.primero
	nodo.longitud += 1
}

func (nodo *ListaEnlazada) dot() string {
	dot := "digraph G {\n"
	dot += "rankdir=LR;\n"
	dot += "node[shape=\"box\"];\n"

	actual := nodo.primero
	for actual != nil {
		dot += actual.valor.Nombre + " -> "
		if actual.siguiente == nil {
			dot += "null;\n"
		}
		actual = actual.siguiente
	}
	dot += "}"

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
	fmt.Println("Lista Simple")
	l1 := ListaEnlazada{}
	l1.insertar(&estudiante.Estudiante{Nombre: "Brandon", Apellido: "Tejaxun"})
	l1.insertar(&estudiante.Estudiante{Nombre: "Andy", Apellido: "Pichiya"})
	l1.insertar(&estudiante.Estudiante{Nombre: "Jefferson", Apellido: "TP"})
	generarGrafo(l1.dot())
}
