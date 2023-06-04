package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

type Nodo struct {
	valor     int
	siguiente *Nodo
	anterior  *Nodo
}

type ListaDoble struct {
	primero  *Nodo
	ultimo   *Nodo
	longitud int
}

func (nodo *ListaDoble) insertar(valor int) {
	if nodo.primero != nil {
		nodo.ultimo.siguiente = &Nodo{valor: valor}
		nodo.ultimo.siguiente.anterior = nodo.ultimo
		nodo.ultimo = nodo.ultimo.siguiente
		nodo.longitud += 1
		return
	}
	nodo.primero = &Nodo{valor: valor}
	nodo.ultimo = nodo.primero
	nodo.longitud += 1
}

func (nodo *ListaDoble) dot() string {
	dot := "digraph G {\n"
	dot += "rankdir=LR;\n"
	dot += "node[shape=\"box\"];\n"

	dot += "nodo1 [label=\"null\"]\n"
	dot += "nodo2 [label=\"null\"]\n"

	actual := nodo.primero
	contador := 0
	for contador < 2 && actual != nil {
		dot += strconv.Itoa(actual.valor) + " -> "
		if actual.siguiente == nil {
			dot += "nodo1;\n"
		}
		actual = actual.siguiente
	}

	actual = nodo.primero
	dot += "nodo2"
	for actual != nil {
		dot += " -> " + strconv.Itoa(actual.valor)
		if actual.siguiente == nil {
			dot += "[dir=\"back\"] ;\n"
		}
		actual = actual.siguiente
	}
	dot += "}"

	return dot
}

// func (nodo *ListaDoble) recorrer2() string {
// 	actual := nodo.ultimo
// 	lista := ""
// 	for actual != nil {
// 		lista += " -> " + strconv.Itoa(actual.valor)
// 		actual = actual.anterior
// 	}
// 	return lista
// }

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

func generarImg() {
	// Ruta del archivo .dot de entrada
	inputFile := "grafo.dot"

	// Ruta del archivo de imagen de salida
	outputFile := "grafo.png"

	// Comando para ejecutar Graphviz
	cmd := exec.Command("dot", "-Tpng", "-o", outputFile, inputFile)

	// Ejecutar el comando
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Archivo de imagen generado:", outputFile)
}

func main() {
	l1 := ListaDoble{}
	l1.insertar(1)
	l1.insertar(2)
	l1.insertar(3)
	l1.insertar(4)
	generarGrafo(l1.dot())
	generarImg()
}
