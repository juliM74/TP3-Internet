package main

import (
	"bufio"
	"fmt"
	"os"
	"tp3/tp3/aeropuertos"
	procesador "tp3/tp3/aeropuertos/procesador"
)

func main() {
	var aeropuertosCSV, vuelosCSV string

	if len(os.Args) == 3 {
		aeropuertosCSV = os.Args[1]
		vuelosCSV = os.Args[2]
	} else {
		aeropuertosCSV = "aeropuertos.csv"
		vuelosCSV = "vuelos.csv"
	}

	sistema, err := aeropuertos.CargarDatos(aeropuertosCSV, vuelosCSV)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error cargando datos:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	procesador.ProcesarEntrada(scanner, sistema)
}
