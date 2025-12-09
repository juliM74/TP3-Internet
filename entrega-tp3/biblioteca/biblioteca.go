package biblioteca

import (
	"math"
	"strings"
	grafo "tp3/tdaGrafo"
	TDACola "tp3/tdaGrafo/tdas/cola"
	TDAHeap "tp3/tdaGrafo/tdas/cola_prioridad"
	TDADiccionario "tp3/tdaGrafo/tdas/hash/diccionario"
)

// nodoDistancia representa un nodo en el heap con su distancia acumulada
type nodoDistancia struct {
	nodo      string
	distancia int
}

// Arista representa una conexion entre dos aeropuertos con un costo
type Arista struct {
	Origen  string
	Destino string
	Costo   int
}

// compararPorDistancia compara dos nodos por distancia para el heap
func compararPorDistancia(a, b nodoDistancia) int {
	return a.distancia - b.distancia // heap minimo
}

// CaminoMinimo aplica Dijkstra en el grafo desde multiples aeropuertos origen
// hasta multiples aeropuertos destino, minimizando segun la funcion obtenerPeso.
// Devuelve el camino como slice de strings y su costo total.
// Si no hay camino, devuelve nil y -1.
func CaminoMinimo(g grafo.Grafo[string, any], origenes, destinos []string, obtenerPeso func(v, w string) int) []string {
	dist := TDADiccionario.CrearHash[string, int]()
	padre := TDADiccionario.CrearHash[string, string]()
	visitados := TDADiccionario.CrearHash[string, bool]()
	h := TDAHeap.CrearHeap(func(a, b nodoDistancia) int {
		return b.distancia - a.distancia
	})

	for _, origen := range origenes {
		dist.Guardar(origen, 0)
		padre.Guardar(origen, "")
		h.Encolar(nodoDistancia{nodo: origen, distancia: 0})
	}

	for !h.EstaVacia() {
		actual := h.Desencolar().nodo
		if visitados.Pertenece(actual) {
			continue
		}
		visitados.Guardar(actual, true)

		for _, vecino := range g.ObtenerVecinos(actual) {
			if visitados.Pertenece(vecino) {
				continue
			}
			peso := obtenerPeso(actual, vecino)
			nuevaDist := dist.Obtener(actual) + peso

			if !dist.Pertenece(vecino) || nuevaDist < dist.Obtener(vecino) {
				dist.Guardar(vecino, nuevaDist)
				padre.Guardar(vecino, actual)
				h.Encolar(nodoDistancia{nodo: vecino, distancia: nuevaDist})
			}
		}
	}

	// Buscar el destino con menor distancia
	mejorDestino := ""
	mejorDist := math.MaxInt32
	for _, destino := range destinos {
		if dist.Pertenece(destino) && dist.Obtener(destino) < mejorDist {
			mejorDist = dist.Obtener(destino)
			mejorDestino = destino
		}
	}

	if mejorDestino == "" {
		return nil
	}

	return reconstruirCamino(padre, mejorDestino)
}

func CostoDeCamino(g grafo.Grafo[string, any], camino []string, obtenerPeso func(v, w string) int) int {
	total := 0
	for i := 0; i < len(camino)-1; i++ {
		v, w := camino[i], camino[i+1]
		peso := obtenerPeso(v, w)
		total += peso
	}
	return total
}

// reconstruirCamino reconstruye el camino desde cualquier origen hasta el destino,
// utilizando un mapa de padres. El camino se devuelve como slice en orden correcto.
func reconstruirCamino(padres TDADiccionario.Diccionario[string, string], destino string) []string {
	camino := []string{}
	actual := destino
	for actual != "" {
		camino = append([]string{actual}, camino...)
		if !padres.Pertenece(actual) {
			break
		}
		actual = padres.Obtener(actual)
	}
	return camino
}

// CaminoMinimoEscalas aplica BFS para encontrar el camino con menor cantidad de escalas.
// Devuelve el camino y la cantidad de escalas. Si no hay camino, devuelve nil y -1.
func CaminoMinimoEscalas(g grafo.Grafo[string, any], origenes, destinos []string) ([]string, int) {
	padre := TDADiccionario.CrearHash[string, string]()
	visitados := TDADiccionario.CrearHash[string, bool]()
	cola := TDACola.CrearColaEnlazada[string]()

	// carga todos los origenes
	for _, origen := range origenes {
		cola.Encolar(origen)
		visitados.Guardar(origen, true)
		padre.Guardar(origen, "")
	}

	var destinoAlcanzado string

	for !cola.EstaVacia() {
		v := cola.Desencolar()

		// es destino?
		for _, destino := range destinos {
			if v == destino {
				destinoAlcanzado = v
				break
			}
		}
		if destinoAlcanzado != "" {
			break
		}

		for _, w := range g.ObtenerVecinos(v) {
			if !visitados.Pertenece(w) {
				visitados.Guardar(w, true)
				padre.Guardar(w, v)
				cola.Encolar(w)
			}
		}
	}

	if destinoAlcanzado == "" {
		return nil, -1 // No hay camino posible
	}

	camino := reconstruirCamino(padre, destinoAlcanzado)
	return camino, len(camino) - 1 // cantidad de escalas
}

// AeropuertosCentrales calcula los n aeropuertos mas centrales en el grafo.
// Usa Dijkstra desde cada nodo para contar cuantas veces aparece cada nodo en caminos minimos.
// Devuelve los n aeropuertos con mayor centralidad.
func AeropuertosCentrales(g grafo.Grafo[string, any], n int, obtenerPeso func(v, w string) int) []string {
	centralidad := TDADiccionario.CrearHash[string, int]()

	vertices := g.Vertices()
	for _, origen := range vertices {
		dist := TDADiccionario.CrearHash[string, int]()
		padre := TDADiccionario.CrearHash[string, string]()
		h := TDAHeap.CrearHeap(compararPorDistancia)

		dist.Guardar(origen, 0)
		h.Encolar(nodoDistancia{nodo: origen, distancia: 0})
		visitados := TDADiccionario.CrearHash[string, bool]()

		for !h.EstaVacia() {
			v := h.Desencolar().nodo
			if visitados.Pertenece(v) {
				continue
			}
			visitados.Guardar(v, true)

			for _, w := range g.ObtenerVecinos(v) {
				nuevaDist := dist.Obtener(v) + obtenerPeso(v, w)
				if !dist.Pertenece(w) || nuevaDist < dist.Obtener(w) {
					dist.Guardar(w, nuevaDist)
					padre.Guardar(w, v)
					h.Encolar(nodoDistancia{nodo: w, distancia: nuevaDist})
				}
			}
		}

		// Suma centralidad para cada nodo intermedio en caminos minimos
		for _, destino := range vertices {
			if destino == origen || !dist.Pertenece(destino) {
				continue
			}
			camino := reconstruirCamino(padre, destino)
			for i := 1; i < len(camino)-1; i++ { // ignora origen y destino
				aeropuerto := camino[i]
				contador := 0
				if centralidad.Pertenece(aeropuerto) {
					contador = centralidad.Obtener(aeropuerto)
				}
				centralidad.Guardar(aeropuerto, contador+1)
			}
		}
	}

	// Obtener los n aeropuertos con mayor centralidad
	todos := []nodoDistancia{}
	centralidad.Iterar(func(aeropuerto string, cantidad int) bool {
		todos = append(todos, nodoDistancia{nodo: aeropuerto, distancia: cantidad})
		return true
	})

	h := TDAHeap.CrearHeap(func(a, b nodoDistancia) int {
		return a.distancia - b.distancia // heap maximo
	})
	for _, nodo := range todos {
		h.Encolar(nodo)
	}

	resultado := []string{}
	for i := 0; i < n && !h.EstaVacia(); i++ {
		resultado = append(resultado, h.Desencolar().nodo)
	}

	return resultado
}

// NuevaAerolinea aplica el algoritmo de Prim para generar una nueva aerolinea (MST).
// Devuelve la lista de aristas que forman el MST.
func NuevaAerolinea(g grafo.Grafo[string, any], obtenerPeso func(v, w string) int) []Arista {
	visitados := TDADiccionario.CrearHash[string, bool]()
	h := TDAHeap.CrearHeap(compararPorDistancia)
	aristas := []Arista{}

	vertices := g.Vertices()
	if len(vertices) == 0 {
		return aristas
	}

	inicio := vertices[0]
	visitados.Guardar(inicio, true)
	for _, w := range g.ObtenerVecinos(inicio) {
		h.Encolar(nodoDistancia{
			nodo:      inicio + "->" + w,
			distancia: obtenerPeso(inicio, w),
		})
	}

	for !h.EstaVacia() && visitados.Cantidad() < len(vertices) {
		arista := h.Desencolar()
		// Separar origen y destino (solución robusta con Split)
		partes := strings.Split(arista.nodo, "->")
		if len(partes) != 2 {
			panic("Formato de arista inválido")
		}
		origen, destino := partes[0], partes[1]

		if visitados.Pertenece(destino) {
			continue
		}

		visitados.Guardar(destino, true)
		aristas = append(aristas, Arista{
			Origen:  origen,
			Destino: destino,
			Costo:   arista.distancia,
		})

		for _, w := range g.ObtenerVecinos(destino) {
			if !visitados.Pertenece(w) {
				h.Encolar(nodoDistancia{
					nodo:      destino + "->" + w,
					distancia: obtenerPeso(destino, w),
				})
			}
		}
	}

	return aristas
}

// OrdenTopologico calcula un orden topologico valido del grafo dirigido
// Devuelve el orden y un bool indicando si fue exitoso (false si hay ciclos)
func OrdenTopologico(g grafo.Grafo[string, any]) ([]string, bool) {
	grado := TDADiccionario.CrearHash[string, int]()
	orden := []string{}
	cola := TDACola.CrearColaEnlazada[string]()

	// Inicializa grados de entrada
	for _, v := range g.Vertices() {
		grado.Guardar(v, 0)
	}
	for _, v := range g.Vertices() {
		for _, w := range g.ObtenerVecinos(v) {
			grado.Guardar(w, grado.Obtener(w)+1)
		}
	}

	// Carga vertices con grado 0
	for _, v := range g.Vertices() {
		if grado.Obtener(v) == 0 {
			cola.Encolar(v)
		}
	}

	// Procesa el grafo
	for !cola.EstaVacia() {
		v := cola.Desencolar()
		orden = append(orden, v)
		for _, w := range g.ObtenerVecinos(v) {
			grado.Guardar(w, grado.Obtener(w)-1)
			if grado.Obtener(w) == 0 {
				cola.Encolar(w)
			}
		}
	}

	if len(orden) != len(g.Vertices()) {
		return nil, false // Hay ciclo
	}
	return orden, true
}
