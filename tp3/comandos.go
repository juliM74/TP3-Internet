package main

import (
	"fmt"
	grafo "tp3/tdaGrafo"
)

func EjecutarListarOperaciones() {
	// Este ya te lo dejo listo porque es trivial (O(1))
	fmt.Println("camino")
	fmt.Println("mas_importantes")
	fmt.Println("conectados")
	fmt.Println("ciclo")
	fmt.Println("lectura")
	fmt.Println("diametro")
	fmt.Println("rango")
	fmt.Println("comunidad")
	fmt.Println("navegacion")
	fmt.Println("clustering")
}

func EjecutarCamino(g grafo.Grafo[string, int], parametros string) {
	fmt.Println("Implementar: camino")
}

func EjecutarMasImportantes(g grafo.Grafo[string, int], parametros string) {
	fmt.Println("Implementar: mas_importantes")
}

func EjecutarConectados(g grafo.Grafo[string, int], parametros string) {
	fmt.Println("Implementar: conectados")
}

func EjecutarCiclo(g grafo.Grafo[string, int], parametros string) {
	fmt.Println("Implementar: ciclo")
}

func EjecutarLectura(g grafo.Grafo[string, int], parametros string) {
	fmt.Println("Implementar: lectura")
}

func EjecutarDiametro(g grafo.Grafo[string, int]) {
	fmt.Println("Implementar: diametro")
}

func EjecutarRango(g grafo.Grafo[string, int], parametros string) {
	fmt.Println("Implementar: rango")
}

func EjecutarComunidad(g grafo.Grafo[string, int], parametros string) {
	fmt.Println("Implementar: comunidad")
}

func EjecutarNavegacion(g grafo.Grafo[string, int], parametros string) {
	fmt.Println("Implementar: navegacion")
}

func EjecutarClustering(g grafo.Grafo[string, int], parametros string) {
	fmt.Println("Implementar: clustering")
}
