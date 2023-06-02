package main

import (
	"fmt"
	"strconv"
)

type NodoCabeza struct {
	indice    int
	anterior  *NodoCabeza
	siguiente *NodoCabeza
	acceso    *NodoInterno
	ultimo    *NodoInterno
}

type NodoInterno struct {
	numero    int
	fila      int
	columna   int
	derecha   *NodoInterno
	izquierda *NodoInterno
	arriba    *NodoInterno
	abajo     *NodoInterno
}

type ListaCabeza struct {
	primero *NodoCabeza
	ultimo  *NodoCabeza
}

func (nodo *ListaCabeza) insertar(indice int) {
	if nodo.primero != nil {
		if indice < nodo.primero.indice {
			nodo.primero.anterior = &NodoCabeza{indice: indice}
			nodo.primero.anterior.siguiente = nodo.primero
			nodo.primero = nodo.primero.anterior
		} else if indice > nodo.ultimo.indice {
			nodo.ultimo.siguiente = &NodoCabeza{indice: indice}
			nodo.ultimo.siguiente.anterior = nodo.ultimo
			nodo.ultimo = nodo.ultimo.siguiente
		} else {
			actual := nodo.primero
			for actual.siguiente != nil {
				if indice > actual.indice && indice < actual.siguiente.indice {
					tmp := &NodoCabeza{indice: indice}
					tmp.anterior = actual
					tmp.siguiente = actual.siguiente

					actual.siguiente.anterior = tmp
					actual.siguiente = tmp
					return
				}
				actual = actual.siguiente
			}
		}
		return
	}
	nodo.primero = &NodoCabeza{indice: indice}
	nodo.ultimo = nodo.primero
}

func (nodo *ListaCabeza) existeNodo(indice int) bool {
	if nodo.primero != nil {
		actual := nodo.primero
		for actual != nil {
			if actual.indice == indice {
				return true
			}
			actual = actual.siguiente
		}
	}
	return false
}

type MatrizDispersa struct {
	accesoF *ListaCabeza
	accesoC *ListaCabeza
}

func (nodo *MatrizDispersa) insertar(fila, columna, numero int) {
	if !nodo.accesoF.existeNodo(fila) {
		nodo.accesoF.insertar(fila)
	}
	if !nodo.accesoC.existeNodo(columna) {
		nodo.accesoC.insertar(columna)
	}
	nodoI := &NodoInterno{fila: fila, columna: columna, numero: numero}
	nodo.agregarFila(fila, nodoI)
	nodo.agregarColumna(columna, nodoI)
}

func (nodo *MatrizDispersa) agregarFila(fila int, nodoI *NodoInterno) {
	actualF := nodo.accesoF.primero
	for actualF != nil {
		if actualF.indice == fila {
			if actualF.acceso != nil {
				if nodoI.columna < actualF.acceso.columna {
					actualF.acceso.izquierda = nodoI
					actualF.acceso.izquierda.derecha = actualF.acceso
					actualF.acceso = actualF.acceso.izquierda
				} else if nodoI.columna > actualF.ultimo.columna {
					actualF.ultimo.derecha = nodoI
					actualF.ultimo.derecha.izquierda = actualF.ultimo
					actualF.ultimo = actualF.ultimo.derecha
				} else {
					actualC := actualF.acceso
					for actualC.derecha != nil {
						if nodoI.columna > actualC.columna && nodoI.columna < actualC.derecha.columna {
							nodoI.izquierda = actualC
							nodoI.derecha = actualC.derecha

							actualC.derecha.izquierda = nodoI
							actualC.derecha = nodoI
							return
						}
						actualC = actualC.derecha
					}
				}
				return
			}
			actualF.acceso = nodoI
			actualF.ultimo = actualF.acceso
			return
		}
		actualF = actualF.siguiente
	}
}

func (nodo *MatrizDispersa) agregarColumna(columna int, nodoI *NodoInterno) {
	actualC := nodo.accesoC.primero
	for actualC != nil {
		if actualC.indice == columna {
			if actualC.acceso != nil {
				if nodoI.fila < actualC.acceso.fila {
					actualC.acceso.arriba = nodoI
					actualC.acceso.arriba.abajo = actualC.acceso
					actualC.acceso = actualC.acceso.arriba
				} else if nodoI.fila > actualC.ultimo.fila {
					actualC.ultimo.abajo = nodoI
					actualC.ultimo.abajo.arriba = actualC.ultimo
					actualC.ultimo = actualC.ultimo.abajo
				} else {
					actualF := actualC.acceso
					for actualF.abajo != nil {
						if nodoI.fila > actualF.fila && nodoI.fila < actualF.abajo.fila {
							nodoI.arriba = actualF
							nodoI.abajo = actualF.abajo

							actualF.abajo.arriba = nodoI
							actualF.abajo = nodoI
							return
						}
						actualF = actualF.abajo
					}
				}
				return
			}
			actualC.acceso = nodoI
			actualC.ultimo = actualC.acceso
			return
		}
		actualC = actualC.siguiente
	}
}

func (nodo *MatrizDispersa) imprimir() {
	actualF := nodo.accesoF.primero
	for actualF != nil {
		actualC := actualF.acceso
		m := ""
		for actualC != nil {
			m += strconv.Itoa(actualC.numero)
			actualC = actualC.derecha
		}
		fmt.Println(m)
		actualF = actualF.siguiente
	}
}

func main() {
	matriz := MatrizDispersa{
		accesoF: &ListaCabeza{},
		accesoC: &ListaCabeza{},
	}

	matriz.insertar(2, 3, 5)
	matriz.insertar(0, 1, 5)
	matriz.insertar(1, 2, 5)
	matriz.insertar(1, 1, 5)

	matriz.insertar(0, 0, 1)
	matriz.insertar(0, 2, 1)
	matriz.insertar(0, 3, 1)

	matriz.insertar(1, 0, 1)
	matriz.insertar(1, 3, 1)

	matriz.insertar(2, 0, 1)
	matriz.insertar(2, 1, 1)
	matriz.insertar(2, 2, 1)

	matriz.imprimir()
}
