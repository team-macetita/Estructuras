package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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

func (nodo *Stack) dot() string {
	dot := "digraph G {\nstack [shape=none, margin=0, label=<<TABLE BORDER=\"0\" CELLBORDER=\"1\" CELLSPACING=\"5\" CELLPADDING=\"20\">\n"

	actual := nodo.ultimo
	for actual != nil {
		dot += "<tr>\n"
		dot += "<td width=\"60\" height=\"60\"><font point-size=\"20\">" + strconv.Itoa(actual.valor) + "</font></td>\n"
		dot += "</tr>\n"
		actual = actual.anterior
	}
	dot += "</TABLE>>];\n"
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
	l1 := Stack{}
	l1.insertar(1)
	l1.insertar(2)
	l1.insertar(3)
	l1.insertar(4)
	generarGrafo(l1.dot())
	generarImg()
}
