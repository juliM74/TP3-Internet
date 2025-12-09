package procesador

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tp3/tp3/aeropuertos"
)

func parseArgs(linea string) ([]string, error) {
	partes := strings.SplitN(linea, " ", 2)
	if len(partes) < 1 {
		return nil, fmt.Errorf("comando vacio")
	}
	args := []string{partes[0]}
	if len(partes) == 2 {
		// separa  argumentos por coma
		for _, arg := range strings.Split(partes[1], ",") {
			args = append(args, strings.TrimSpace(arg))
		}
	}
	return args, nil
}

// ProcesarEntrada procesa los comandos de entrada estandar.
func ProcesarEntrada(scanner *bufio.Scanner, sistema aeropuertos.SistemaDeAeropuertos) {
	for scanner.Scan() {
		linea := strings.TrimSpace(scanner.Text())
		if len(linea) == 0 {
			continue
		}

		partes, err := parseArgs(linea)
		if err != nil || len(partes) == 0 {
			fmt.Fprintln(os.Stderr, "Error en comando", linea)
			continue
		}
		comando := partes[0]
		args := partes[1:]

		switch comando {
		case "camino_mas":
			if len(args) != 3 {
				continue
			}

			camino, _, err := sistema.Camino_mas(args[0], args[1], args[2])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error en comando camino_mas")
				continue
			}

			for i, aeropuerto := range camino {
				if i > 0 {
					fmt.Print(" -> ")
				}
				fmt.Print(aeropuerto)
			}
			fmt.Println()

		case "camino_escalas":
			if len(args) != 2 {
				fmt.Fprintln(os.Stderr, "Error en comando camino_escalas")
				continue
			}

			camino, escalas, err := sistema.Camino_escalas(args[0], args[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error en comando camino_escalas")
				continue
			}

			for i, aeropuerto := range camino {
				if i > 0 {
					fmt.Print(" -> ")
				}
				fmt.Print(aeropuerto)
			}
			fmt.Println()
			fmt.Println("Cantidad de escalas:", escalas)

		case "centralidad":
			if len(args) != 1 {
				fmt.Fprintln(os.Stderr, "Error en comando centralidad")
				continue
			}
			n, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error en comando centralidad")
				continue
			}
			resultado, err := sistema.Centralidad(n)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error en comando centralidad")
			} else {
				for _, aeropuerto := range resultado {
					fmt.Println(aeropuerto)
				}
				fmt.Println("OK")
			}

		case "nueva_aerolinea":
			err := sistema.Nueva_aerolinea("mst.csv")

			if err != nil {
				fmt.Fprintln(os.Stderr, "Error en comando nueva_aerolinea")
			} else {
				fmt.Println("OK")
			}

		case "itinerario":
			if len(args) != 1 {
				fmt.Fprintln(os.Stderr, "Error en comando itinerario")
				continue
			}
			orden, caminos, err := sistema.Itinerario(args[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error en comando itinerario")
				return
			}
			fmt.Println(strings.Join(orden, ", "))
			for _, camino := range caminos {
				fmt.Println(strings.Join(camino, " -> "))
			}

		case "exportar_kml":
			if len(args) != 1 {
				fmt.Fprintln(os.Stderr, "Error en comando exportar_kml")
				continue
			}
			err := sistema.Exportar_kml(args[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error en comando exportar_kml")
			} else {
				fmt.Println("OK")
			}

		default:
			fmt.Fprintln(os.Stderr, "Error en comando", comando)
		}
	}
}
