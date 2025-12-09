package parser

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	grafo "tp3/tdaGrafo"
	TDADiccionario "tp3/tdaGrafo/tdas/hash/diccionario"
)

type DatosVuelo struct {
	Costo      int
	Tiempo     int
	Frecuencia int
}

func CargarAeropuertos(r io.Reader, g grafo.Grafo[string, any], ciudades TDADiccionario.Diccionario[string, []string], coordenadas TDADiccionario.Diccionario[string, [2]float64]) error {
	scanner := bufio.NewScanner(r)
	linea := 0

	for scanner.Scan() {
		linea++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		fields := strings.Split(line, ",")
		if len(fields) != 4 {
			return fmt.Errorf("linea %d mal formada en aeropuertos.csv: %q", linea, line)
		}

		ciudad := strings.TrimSpace(fields[0])
		aeropuerto := strings.TrimSpace(fields[1])
		latitud, errLat := strconv.ParseFloat(strings.TrimSpace(fields[2]), 64)
		longitud, errLon := strconv.ParseFloat(strings.TrimSpace(fields[3]), 64)
		if aeropuerto == "" || ciudad == "" || errLat != nil || errLon != nil {
			return fmt.Errorf("linea %d con datos invalidos: %q", linea, line)
		}

		if !g.Pertenece(aeropuerto) {
			g.AgregarVertice(aeropuerto)
		}

		coordenadas.Guardar(aeropuerto, [2]float64{latitud, longitud})

		actuales := []string{}
		if ciudades.Pertenece(ciudad) {
			actuales = ciudades.Obtener(ciudad)
		}
		ciudades.Guardar(ciudad, append(actuales, aeropuerto))
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error al leer aeropuertos.csv: %w", err)
	}
	return nil
}

func CargarVuelos(r io.Reader, g grafo.Grafo[string, any]) error {
	scanner := bufio.NewScanner(r)
	linea := 0

	for scanner.Scan() {
		linea++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		fields := strings.Split(line, ",")
		if len(fields) != 5 {
			return fmt.Errorf("linea %d mal formada en vuelos.csv: %q", linea, line)
		}

		origen := strings.TrimSpace(fields[0])
		destino := strings.TrimSpace(fields[1])
		costo, err1 := strconv.Atoi(fields[2])
		tiempo, err2 := strconv.Atoi(fields[3])
		frecuencia, err3 := strconv.Atoi(fields[4])
		if origen == "" || destino == "" || err1 != nil || err2 != nil || err3 != nil {
			return fmt.Errorf("linea %d con datos invalidos: %q", linea, line)
		}

		if !g.Pertenece(origen) || !g.Pertenece(destino) {
			return fmt.Errorf("linea %d con aeropuerto inexistente: %q", linea, line)
		}

		nuevo := DatosVuelo{Costo: costo, Tiempo: tiempo, Frecuencia: frecuencia}

		if viejo, ok := g.ObtenerDatos(origen, destino); ok {
			if viejo.(DatosVuelo).Costo <= costo {
				continue
			}
			g.BorrarArista(origen, destino)
			g.BorrarArista(destino, origen)
		}

		g.AgregarArista(origen, destino, nuevo)
		g.AgregarArista(destino, origen, nuevo)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error al leer vuelos.csv: %w", err)
	}
	return nil
}
