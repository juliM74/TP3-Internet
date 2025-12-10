package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tp3/internet/utils"
	"tp3/internet/estado"
)

// ===============================
//           CONSTANTES
// ===============================

const (
	ARGS_ESPERADOS       = 2
	MENSAJE_USO          = "Uso: ./netstats <archivo_wiki.tsv>"
	SEPARADOR_COMANDO    = " "
	ERROR_LECTURA_STDIN  = "Error leyendo entrada estándar:"
)

func main() {

	// ==========================
	//   1. Validación argumentos
	// ==========================

	if len(os.Args) != ARGS_ESPERADOS {
		fmt.Println(MENSAJE_USO)
		return
	}

	rutaArchivo := os.Args[1]

	// ==========================
	//   2. Cargar el grafo
	// ==========================

	g := utils.CargarGrafo(rutaArchivo)

	// ==========================
	//   3. Crear Estado
	// ==========================

	est := estado.NuevoEstado(g)

	// ==========================
	//   4. Leer comandos STDIN
	// ==========================

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		linea := strings.TrimSpace(scanner.Text())
		if linea == "" {
			continue
		}

		// comando y parámetros
		partes := strings.SplitN(linea, SEPARADOR_COMANDO, 2)
		comando := partes[0]

		parametros := ""
		if len(partes) == 2 {
			parametros = partes[1]
		}

		// ==========================
		//   5. Ejecutar comando
		// ==========================

		utils.EjecutarLinea(est, comando, parametros)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(ERROR_LECTURA_STDIN, err)
	}
}