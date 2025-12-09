package main

import (
	"bufio"
	"os"
	"strings"
	grafo "tp3/tdaGrafo"
)

const capacidadMax = 1024 * 1024 * 4

func CargarGrafo(rutaArchivo string) grafo.Grafo[string, int] {
	// 1. Abrir el archivo
	archivo, err := os.Open(rutaArchivo)
	if err != nil {
		panic("Error al abrir el archivo: " + err.Error())
	}
	defer archivo.Close()

	// 2. Crear el grafo DIRIGIDO
	g := grafo.CrearGrafo[string, int](true)
	scanner := bufio.NewScanner(archivo)

	buf := make([]byte, capacidadMax)
	scanner.Buffer(buf, capacidadMax)

	for scanner.Scan() {
		linea := scanner.Text()
		campos := strings.Split(linea, "\t")

		if len(campos) == 0 {
			continue
		}
		origen := campos[0]
		g.AgregarVertice(origen)

		for i := 1; i < len(campos); i++ {
			destino := campos[i]

			if destino == "" {
				continue
			}

			g.AgregarVertice(destino)
			g.AgregarArista(origen, destino, 1)
		}
	}

	if err := scanner.Err(); err != nil {
		panic("Error leyendo el archivo: " + err.Error())
	}

	return g
}
