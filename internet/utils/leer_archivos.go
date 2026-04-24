package utils

import (
	"bufio"
	"os"
	"strings"
	grafo "tp3/tdaGrafo"
)

const capacidadMax = 4 * 1024 * 1024

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
		if !g.Pertenece(origen) {
			g.AgregarVertice(origen)
		}

		for i := 1; i < len(campos); i++ {
			destino := campos[i]
			if destino == "" {
				continue
			}

			if !g.Pertenece(destino) {
				g.AgregarVertice(destino)
			}
			g.AgregarArista(origen, destino, 1)
		}
	}

	if err := scanner.Err(); err != nil {
		panic("Error leyendo el archivo: " + err.Error())
	}

	return g
}
