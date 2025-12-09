package aeropuertos

// SistemaDeAeropuertos representa el sistema que gestiona los aeropuertos y vuelos
// Permite consultar caminos, centralidad, MST y exportar KML
type SistemaDeAeropuertos interface {
	// Camino_mas calcula el camino mas rapido o mas barato entre dos ciudades
	// criterio puede ser "rapido" o "barato"
	// Devuelve la lista de aeropuertos en el camino, el costo total y un error en caso de fallo
	Camino_mas(criterio, origen, destino string) ([]string, int, error)

	// Camino_escalas calcula el camino con la menor cantidad de escalas entre dos ciudades
	// Devuelve la lista de aeropuertos en el camino, la cantidad de escalas y un error en caso de fallo
	Camino_escalas(origen, destino string) ([]string, int, error)

	// Centralidad devuelve los n aeropuertos mas importantes segun centralidad
	Centralidad(n int) ([]string, error)

	// Nueva_aerolinea genera un archivo con las rutas de la nueva aerolinea (MST)
	// Devuelve un error si no se pudo generar el archivo
	Nueva_aerolinea(rutaSalida string) error

	// Itinerario calcula el itinerario cultural segun el archivo de ruta
	// Devuelve el orden de las ciudades, los caminos entre ellas y un error si hubo problema
	Itinerario(ruta string) ([]string, [][]string, error)

	// Exportar_kml exporta un archivo KML con el ultimo recorrido realizado
	// Devuelve un error si no se pudo exportar
	Exportar_kml(rutaSalida string) error
}
