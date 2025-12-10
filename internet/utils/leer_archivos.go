package utils

import (
	"bufio"
	"os"
	"strings"
	grafo "tp3/tdaGrafo"
)

const capacidadMax = 4 * 1024 * 1024 // 4MB

// CargarGrafo lee un archivo TSV donde:
// columna 0 = artículo
// columnas 1..n = links salientes
// y construye un grafo dirigido con peso 1 por arista.
func CargarGrafo(rutaArchivo string) grafo.Grafo[string, int] {

	archivo, err := os.Open(rutaArchivo)
	if err != nil {
		panic("Error al abrir el archivo: " + err.Error())
	}
	defer archivo.Close()

	g := grafo.CrearGrafo[string, int](true)

	scanner := bufio.NewScanner(archivo)
	buf := make([]byte, capacidadMax)
	scanner.Buffer(buf, capacidadMax)

	for scanner.Scan() {
		linea := scanner.Text()
		if len(linea) == 0 {
			continue
		}

		campos := strings.Split(linea, "\t")
		origen := campos[0]

		// Asegurar que exista el vértice origen
		if !g.ExisteVertice(origen) {
			g.AgregarVertice(origen)
		}

		// Procesar los destinos
		for i := 1; i < len(campos); i++ {
			destino := campos[i]
			if destino == "" {
				continue
			}

			if !g.ExisteVertice(destino) {
				g.AgregarVertice(destino)
			}

			// Peso 1 por defecto (grafo no pesado según TP3)
			g.AgregarArista(origen, destino, 1)
		}
	}

	if err := scanner.Err(); err != nil {
		panic("Error leyendo el archivo: " + err.Error())
	}

	return g
}