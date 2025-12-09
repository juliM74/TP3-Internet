package aeropuertos

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"tp3/biblioteca"
	grafo "tp3/tdaGrafo"
	TDADiccionario "tp3/tdaGrafo/tdas/hash/diccionario"
	parser "tp3/tp3/aeropuertos/parser"
)

// sistema representa la implementacion del sistema de aeropuertos
type sistema struct {
	grafo           grafo.Grafo[string, any]
	ciudades        TDADiccionario.Diccionario[string, []string] // lista de aeropuertos (ciudad)
	ultimosCaminos  [][]string                                   // ultimo recorrido (para KML)
	ultimaAerolinea []biblioteca.Arista                          // ultima aerolinea generada (para KML)
	coordenadas     TDADiccionario.Diccionario[string, [2]float64]
}

func crearSistema(g grafo.Grafo[string, any], ciudades TDADiccionario.Diccionario[string, []string], coordenadas TDADiccionario.Diccionario[string, [2]float64]) SistemaDeAeropuertos {
	return &sistema{
		grafo:       g,
		ciudades:    ciudades,
		coordenadas: TDADiccionario.CrearHash[string, [2]float64](),
	}
}

// CargarDatos carga los aeropuertos y vuelos desde los archivos CSV y devuelve el sistema
func CargarDatos(aeropuertosCSV, vuelosCSV string) (SistemaDeAeropuertos, error) {
	grafoVuelos := grafo.CrearGrafo[string, any]()
	ciudades := TDADiccionario.CrearHash[string, []string]()
	coordenadas := TDADiccionario.CrearHash[string, [2]float64]()

	archivoAeropuertos, err := os.Open(aeropuertosCSV)
	if err != nil {
		return nil, fmt.Errorf("error al abrir archivo de aeropuertos: %w", err)
	}
	defer archivoAeropuertos.Close()

	if err := parser.CargarAeropuertos(archivoAeropuertos, grafoVuelos, ciudades, coordenadas); err != nil {
		return nil, fmt.Errorf("error al cargar aeropuertos: %w", err)
	}

	archivoVuelos, err := os.Open(vuelosCSV)
	if err != nil {
		return nil, fmt.Errorf("error al abrir archivo de vuelos: %w", err)
	}
	defer archivoVuelos.Close()

	if err := parser.CargarVuelos(archivoVuelos, grafoVuelos); err != nil {
		return nil, fmt.Errorf("error al cargar vuelos: %w", err)
	}

	return crearSistema(grafoVuelos, ciudades, coordenadas), nil
}

// implementacion de los metodos de la interfaz
func (s *sistema) Camino_mas(tipo, ciudadOrigen, ciudadDestino string) ([]string, int, error) {
	ciudadOrigen = strings.TrimSpace(ciudadOrigen)
	ciudadDestino = strings.TrimSpace(ciudadDestino)

	if !s.ciudades.Pertenece(ciudadOrigen) || !s.ciudades.Pertenece(ciudadDestino) {
		return nil, -1, fmt.Errorf("una o ambas ciudades no existen")
	}

	origenes := s.ciudades.Obtener(ciudadOrigen)
	destinos := s.ciudades.Obtener(ciudadDestino)

	if len(origenes) == 0 || len(destinos) == 0 {
		return nil, -1, fmt.Errorf("una o ambas ciudades no existen")
	}

	var obtenerPeso func(v, w string) int
	switch tipo {
	case "barato":
		obtenerPeso = func(v, w string) int {
			dato, existe := s.grafo.ObtenerDatos(v, w)
			if !existe {
				dato, existe = s.grafo.ObtenerDatos(w, v)
				if !existe {
					return math.MaxInt32
				}
			}
			return dato.(parser.DatosVuelo).Tiempo
		}

	case "rapido":
		obtenerPeso = func(v, w string) int {
			dato, existe := s.grafo.ObtenerDatos(v, w)
			if !existe {
				return math.MaxInt32
			}
			return dato.(parser.DatosVuelo).Costo
		}
	default:
		return nil, -1, fmt.Errorf("tipo de camino invalido")
	}

	camino := biblioteca.CaminoMinimo(s.grafo, origenes, destinos, obtenerPeso)

	if camino == nil {
		return nil, -1, fmt.Errorf("no hay camino entre las ciudades")
	}
	costo := biblioteca.CostoDeCamino(s.grafo, camino, obtenerPeso)
	s.ultimosCaminos = [][]string{camino}
	return camino, costo, nil
}

func (s *sistema) Camino_escalas(ciudadOrigen, ciudadDestino string) ([]string, int, error) {
	ciudadOrigen = strings.TrimSpace(ciudadOrigen)
	ciudadDestino = strings.TrimSpace(ciudadDestino)

	if !s.ciudades.Pertenece(ciudadOrigen) || !s.ciudades.Pertenece(ciudadDestino) {
		return nil, -1, fmt.Errorf("una o ambas ciudades no existen")
	}

	origenes := s.ciudades.Obtener(ciudadOrigen)
	destinos := s.ciudades.Obtener(ciudadDestino)

	if len(origenes) <= 0 || len(destinos) <= 0 {
		return nil, -1, fmt.Errorf("una o ambas ciudades no existen")
	}

	camino, escalas := biblioteca.CaminoMinimoEscalas(s.grafo, origenes, destinos)
	if camino == nil {
		return nil, -1, fmt.Errorf("no hay camino entre las ciudades")
	}

	// Guardar el camino como ultimo recorrido (para KML)
	s.ultimosCaminos = [][]string{camino}
	return camino, escalas, nil
}

func (s *sistema) Centralidad(n int) ([]string, error) {
	if n <= 0 {
		return nil, fmt.Errorf("n debe ser mayor a 0")
	}

	obtenerPeso := func(v, w string) int {
		datos, ok := s.grafo.ObtenerDatos(v, w)
		if !ok {
			return math.MaxInt32 // Evitar caminos no validos
		}
		return datos.(parser.DatosVuelo).Frecuencia
	}

	resultado := biblioteca.AeropuertosCentrales(s.grafo, n, obtenerPeso)
	if len(resultado) == 0 {
		return nil, fmt.Errorf("no se pudo calcular la centralidad")
	}
	return resultado, nil
}

func (s *sistema) Nueva_aerolinea(rutaArchivo string) error {
	obtenerPeso := func(v, w string) int {
		datos, _ := s.grafo.ObtenerDatos(v, w)
		return datos.(parser.DatosVuelo).Costo
	}

	aristas := biblioteca.NuevaAerolinea(s.grafo, obtenerPeso)
	if len(aristas) == 0 {
		return fmt.Errorf("no se pudo generar la aerolínea")
	}

	// Guarda el MST en el sistema (para el comando exportar_kml)
	s.ultimaAerolinea = aristas

	// Escribe las rutas en el archivo CSV
	archivo, err := os.Create(rutaArchivo)
	if err != nil {
		return fmt.Errorf("error al crear el archivo: %w", err)
	}
	defer archivo.Close()

	for _, arista := range aristas {
		// tiempo y frecuencia
		datos, _ := s.grafo.ObtenerDatos(arista.Origen, arista.Destino)
		datoVuelo := datos.(parser.DatosVuelo)

		fmt.Fprintf(archivo, "%s,%s,%d,%d,%d\n",
			arista.Origen, arista.Destino, arista.Costo, datoVuelo.Tiempo, datoVuelo.Frecuencia)
	}

	return nil
}
func (s *sistema) Itinerario(ruta string) ([]string, [][]string, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, nil, fmt.Errorf("error al abrir archivo de ruta: %w", err)
	}
	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)

	if !scanner.Scan() {
		return nil, nil, fmt.Errorf("archivo vacío o mal formado")
	}

	linea := strings.TrimSpace(scanner.Text())
	if linea == "" {
		return nil, nil, fmt.Errorf("no hay ciudades en la primera línea")
	}

	// Extraer ciudades a visitar
	ciudades := []string{}
	for _, c := range strings.Split(linea, ",") {
		c = strings.TrimSpace(c)
		if c != "" {
			ciudades = append(ciudades, c)
		}
	}

	if len(ciudades) == 0 {
		return nil, nil, fmt.Errorf("no hay ciudades válidas en el itinerario")
	}

	// Crear grafo dirigido con las restricciones
	restricciones := grafo.CrearGrafo[string, any]()
	for _, ciudad := range ciudades {
		restricciones.AgregarVertice(ciudad)
	}

	for scanner.Scan() {
		linea := strings.TrimSpace(scanner.Text())
		if linea == "" {
			continue
		}
		partes := strings.Split(linea, ",")
		if len(partes) != 2 {
			return nil, nil, fmt.Errorf("restricción mal formada: %s", linea)
		}
		origen := strings.TrimSpace(partes[0])
		destino := strings.TrimSpace(partes[1])
		if !restricciones.Pertenece(origen) || !restricciones.Pertenece(destino) {
			return nil, nil, fmt.Errorf("ciudad de restricción no válida: %s", linea)
		}
		restricciones.AgregarArista(origen, destino, nil)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error al leer el archivo de restricciones: %w", err)
	}

	orden, ok := biblioteca.OrdenTopologico(restricciones)
	if !ok {
		return nil, nil, fmt.Errorf("hay un ciclo en las restricciones")
	}

	caminos := [][]string{}
	for i := 0; i < len(orden)-1; i++ {
		origenes := s.ciudades.Obtener(orden[i])
		destinos := s.ciudades.Obtener(orden[i+1])
		if len(origenes) == 0 || len(destinos) == 0 {
			return nil, nil, fmt.Errorf("una ciudad no tiene aeropuertos")
		}

		camino := biblioteca.CaminoMinimo(s.grafo, origenes, destinos, func(v, w string) int {
			dato, ok := s.grafo.ObtenerDatos(v, w)
			if !ok {
				return math.MaxInt32
			}
			return dato.(parser.DatosVuelo).Tiempo
		})

		if camino == nil {
			return nil, nil, fmt.Errorf("no hay camino entre %s y %s", orden[i], orden[i+1])
		}
		caminos = append(caminos, camino)
	}

	s.ultimosCaminos = caminos
	return orden, caminos, nil
}

func (s *sistema) Exportar_kml(rutaArchivo string) error {
	if len(s.ultimosCaminos) == 0 || len(s.ultimosCaminos[0]) < 2 {
		return fmt.Errorf("no hay camino disponible para exportar")
	}

	camino := s.ultimosCaminos[0]

	archivo, err := os.Create(rutaArchivo)
	if err != nil {
		return fmt.Errorf("error al crear archivo: %w", err)
	}
	defer archivo.Close()

	fmt.Fprintln(archivo, `<?xml version="1.0" encoding="UTF-8"?>`)
	fmt.Fprintln(archivo, `<kml xmlns="http://earth.google.com/kml/2.1">`)
	fmt.Fprintln(archivo, `<Document>`)

	// Marcadores para cada aeropuerto
	for _, codigo := range camino {
		latLon := s.coordenadas.Obtener(codigo)
		lat := latLon[0]
		lon := latLon[1]
		fmt.Fprintln(archivo, `  <Placemark>`)
		fmt.Fprintf(archivo, `    <name>%s</name>`+"\n", codigo)
		fmt.Fprintln(archivo, `    <Point>`)
		fmt.Fprintf(archivo, `      <coordinates>%f,%f</coordinates>`+"\n", lon, lat)
		fmt.Fprintln(archivo, `    </Point>`)
		fmt.Fprintln(archivo, `  </Placemark>`)
	}

	fmt.Fprintln(archivo, "")

	// Conexiones entre aeropuertos
	for i := 0; i < len(camino)-1; i++ {
		origen := camino[i]
		destino := camino[i+1]
		latLonOrig := s.coordenadas.Obtener(origen)
		latLonDest := s.coordenadas.Obtener(destino)

		fmt.Fprintln(archivo, `  <Placemark>`)
		fmt.Fprintln(archivo, `    <LineString>`)
		fmt.Fprintf(archivo, `      <coordinates>%f,%f %f,%f</coordinates>`+"\n",
			latLonOrig[1], latLonOrig[0], latLonDest[1], latLonDest[0])
		fmt.Fprintln(archivo, `    </LineString>`)
		fmt.Fprintln(archivo, `  </Placemark>`)
	}

	fmt.Fprintln(archivo, `</Document>`)
	fmt.Fprintln(archivo, `</kml>`)

	return nil
}
