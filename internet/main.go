package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tp3/internet/estado"
	"tp3/internet/utils"
)

const (
	ARGS_ESPERADOS      = 2
	MENSAJE_USO         = "Uso: ./netstats <archivo_wiki.tsv>"
	SEPARADOR_COMANDO   = " "
	ERROR_LECTURA_STDIN = "Error leyendo entrada estándar:"
)

func main() {
	if len(os.Args) != ARGS_ESPERADOS {
		fmt.Println(MENSAJE_USO)
		return
	}

	rutaArchivo := os.Args[1]
	g := utils.CargarGrafo(rutaArchivo)

	est := estado.NuevoEstado(g)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		linea := strings.TrimSpace(scanner.Text())
		if len(linea) == 0 || strings.HasPrefix(linea, "#") {
			continue
		}
		if linea == "" {
			continue
		}
		partes := strings.SplitN(linea, SEPARADOR_COMANDO, 2)
		comando := partes[0]

		parametros := ""
		if len(partes) == 2 {
			parametros = partes[1]
		}

		utils.EjecutarLinea(est, comando, parametros)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(ERROR_LECTURA_STDIN, err)
	}
}
