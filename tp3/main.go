package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1. Verificación de argumentos
	// El programa espera: ./program file.tsv
	argumentos := os.Args
	if len(argumentos) != 2 {
		fmt.Println("Uso: ./netstats <archivo_grafo>")
		return
	}

	rutaArchivo := argumentos[1]

	// 2. Carga del Grafo (Usamos la función de lectura_archivos.go)
	// Nota: Como lectura_archivos.go es package main, no hace falta importarlo,
	// Go lo ve automáticamente al compilar el paquete.
	fmt.Println("Cargando grafo...")
	g := CargarGrafo(rutaArchivo)
	fmt.Printf("Grafo cargado con %d vértices y una cantidad inmensa de aristas :)\n", len(g.Vertices()))

	// 3. Procesamiento de Comandos (Loop Principal)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		linea := input.Text()

		// Limpiamos espacios extra al inicio/final
		linea = strings.TrimSpace(linea)
		if len(linea) == 0 {
			continue
		}

		// Separamos el comando de los parámetros.
		// Ejemplo: "camino Argentina,Francia" -> ["camino", "Argentina,Francia"]
		// Usamos SplitN con n=2 para que solo corte en el primer espacio.
		partes := strings.SplitN(linea, " ", 2)
		comando := partes[0]
		parametros := ""
		if len(partes) > 1 {
			parametros = partes[1]
		}

		// 4. Dispatcher (Selector de comandos)
		// Las funciones a las que llamamos aquí (Ej: EjecutarCamino)
		// las implementaremos en el archivo 'comandos.go'.
		switch comando {
		case "listar_operaciones":
			EjecutarListarOperaciones()
		case "camino":
			EjecutarCamino(g, parametros)
		case "mas_importantes":
			EjecutarMasImportantes(g, parametros)
		case "conectados":
			EjecutarConectados(g, parametros)
		case "ciclo":
			EjecutarCiclo(g, parametros)
		case "lectura":
			EjecutarLectura(g, parametros)
		case "diametro":
			EjecutarDiametro(g)
		case "rango":
			EjecutarRango(g, parametros)
		case "comunidad":
			EjecutarComunidad(g, parametros)
		case "navegacion":
			EjecutarNavegacion(g, parametros)
		case "clustering":
			EjecutarClustering(g, parametros)
		default:
			fmt.Printf("Comando desconocido: %s\n", comando)
		}
	}
}
