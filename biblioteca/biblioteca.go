package biblioteca

import (
	grafo "tp3/tdaGrafo"
	TDACola "tp3/tdaGrafo/tdas/cola"
	TDAHeap "tp3/tdaGrafo/tdas/cola_prioridad"
	TDADiccionario "tp3/tdaGrafo/tdas/hash/diccionario"
)

// ==================================== ESTRUCTURAS AUXILIARES ====================================

type elementoDistancia[K comparable] struct {
	vertice   K
	distancia int
}

// aristaHeap se usa en Prim para guardar la conexión completa en el heap.

type aristaHeap[K comparable] struct {
	origen  K
	destino K
	peso    int
}

// Arista es la estructura publica que devolvemos en el MST.
type Arista[K comparable] struct {
	Origen  K
	Destino K
	Peso    int
}

// ==================================== ALGORITMOS ====================================

// CaminoMinimo aplica el algoritmo de Dijkstra.
func CaminoMinimo[K comparable, W any](g grafo.Grafo[K, W], origen, destino K, obtenerPeso func(W) int, cmp func(K, K) bool) ([]K, int) {
	dist := TDADiccionario.CrearHash[K, int](cmp)
	padre := TDADiccionario.CrearHash[K, K](cmp)
	visitados := TDADiccionario.CrearHash[K, bool](cmp)

	h := TDAHeap.CrearHeap(func(a, b elementoDistancia[K]) int {
		return b.distancia - a.distancia
	})

	dist.Guardar(origen, 0)
	h.Encolar(elementoDistancia[K]{vertice: origen, distancia: 0})

	for !h.EstaVacia() {
		actual := h.Desencolar().vertice

		if actual == destino {
			return reconstruirCamino(padre, origen, destino), dist.Obtener(destino)
		}

		if visitados.Pertenece(actual) {
			continue
		}
		visitados.Guardar(actual, true)

		for _, vecino := range g.ObtenerVecinos(actual) {
			if visitados.Pertenece(vecino) {
				continue
			}

			datoArista, _ := g.ObtenerDatos(actual, vecino)
			peso := obtenerPeso(datoArista)
			nuevaDist := dist.Obtener(actual) + peso

			if !dist.Pertenece(vecino) || nuevaDist < dist.Obtener(vecino) {
				dist.Guardar(vecino, nuevaDist)
				padre.Guardar(vecino, actual)
				h.Encolar(elementoDistancia[K]{vertice: vecino, distancia: nuevaDist})
			}
		}
	}

	return nil, -1 // No hay camino
}

// CaminoMinimoBFS encuentra el camino con menos aristas (sin pesos).
func CaminoMinimoBFS[K comparable, W any](g grafo.Grafo[K, W], origen, destino K, cmp func(K, K) bool) []K {
	visitados := TDADiccionario.CrearHash[K, bool](cmp)
	padre := TDADiccionario.CrearHash[K, K](cmp)
	cola := TDACola.CrearColaEnlazada[K]()

	cola.Encolar(origen)
	visitados.Guardar(origen, true)

	encontrado := false
	for !cola.EstaVacia() {
		v := cola.Desencolar()
		if v == destino {
			encontrado = true
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

	if !encontrado {
		return nil
	}
	return reconstruirCamino(padre, origen, destino)
}

// MSTPrim aplica el algoritmo de Prim para encontrar el arbol de tendido minimo.
// Devuelve una lista de Aristas que forman el MST.
func MSTPrim[K comparable, W any](g grafo.Grafo[K, W], obtenerPeso func(W) int, cmp func(K, K) bool) []Arista[K] {
	visitados := TDADiccionario.CrearHash[K, bool](cmp)
	aristasMST := []Arista[K]{}

	h := TDAHeap.CrearHeap(func(a, b aristaHeap[K]) int {
		return b.peso - a.peso
	})

	vertices := g.Vertices()
	if len(vertices) == 0 {
		return aristasMST
	}

	inicio := vertices[0]
	visitados.Guardar(inicio, true)

	for _, vecino := range g.ObtenerVecinos(inicio) {
		dato, _ := g.ObtenerDatos(inicio, vecino)
		h.Encolar(aristaHeap[K]{
			origen:  inicio,
			destino: vecino,
			peso:    obtenerPeso(dato),
		})
	}

	for !h.EstaVacia() && visitados.Cantidad() < len(vertices) {
		arista := h.Desencolar()

		if visitados.Pertenece(arista.destino) {
			continue
		}

		// Agregamos al MST
		aristasMST = append(aristasMST, Arista[K]{
			Origen:  arista.origen,
			Destino: arista.destino,
			Peso:    arista.peso,
		})

		visitados.Guardar(arista.destino, true)

		// Agregamos las nuevas aristas disponibles desde el nodo recien visitado
		nuevoOrigen := arista.destino
		for _, w := range g.ObtenerVecinos(nuevoOrigen) {
			if !visitados.Pertenece(w) {
				dato, _ := g.ObtenerDatos(nuevoOrigen, w)
				h.Encolar(aristaHeap[K]{
					origen:  nuevoOrigen,
					destino: w,
					peso:    obtenerPeso(dato),
				})
			}
		}
	}

	return aristasMST
}

// Centralidad calcula los n vertices mas importantes usando Centralidad de Intermediacion.
// costo: O(V * E log V).
func Centralidad[K comparable, W any](g grafo.Grafo[K, W], n int, obtenerPeso func(W) int, cmp func(K, K) bool) []K {
	contador := TDADiccionario.CrearHash[K, int](cmp)
	vertices := g.Vertices()

	// Inicializamos contadores
	for _, v := range vertices {
		contador.Guardar(v, 0)
	}

	// Corremos Dijkstra desde cada vertice
	for _, origen := range vertices {
		// Dijkstra interno modificado para centralidad
		dist := TDADiccionario.CrearHash[K, int](cmp)
		padre := TDADiccionario.CrearHash[K, K](cmp)
		h := TDAHeap.CrearHeap(func(a, b elementoDistancia[K]) int {
			return b.distancia - a.distancia
		})

		dist.Guardar(origen, 0)
		h.Encolar(elementoDistancia[K]{vertice: origen, distancia: 0})

		visitados := TDADiccionario.CrearHash[K, bool](cmp)

		for !h.EstaVacia() {
			v := h.Desencolar().vertice
			if visitados.Pertenece(v) {
				continue
			}
			visitados.Guardar(v, true)

			for _, w := range g.ObtenerVecinos(v) {
				dato, _ := g.ObtenerDatos(v, w)
				nuevaDist := dist.Obtener(v) + obtenerPeso(dato)

				if !dist.Pertenece(w) || nuevaDist < dist.Obtener(w) {
					dist.Guardar(w, nuevaDist)
					padre.Guardar(w, v)
					h.Encolar(elementoDistancia[K]{vertice: w, distancia: nuevaDist})
				}
			}
		}

		// Reconstruimos caminos hacia todos los demas para sumar al contador
		for _, destino := range vertices {
			if destino == origen || !dist.Pertenece(destino) {
				continue
			}

			actual := padre.Obtener(destino)
			for actual != origen {
				contador.Guardar(actual, contador.Obtener(actual)+1)
				if !padre.Pertenece(actual) {
					break
				}
				actual = padre.Obtener(actual)
			}
		}
	}

	// Usamos un Heap para obtener los Top N
	type nodoRanking struct {
		vertice K
		valor   int
	}
	heapRanking := TDAHeap.CrearHeap(func(a, b nodoRanking) int {
		return a.valor - b.valor // Heap de Maximos
	})

	contador.Iterar(func(v K, cant int) bool {
		heapRanking.Encolar(nodoRanking{v, cant})
		return true
	})

	resultado := []K{}
	for i := 0; i < n && !heapRanking.EstaVacia(); i++ {
		resultado = append(resultado, heapRanking.Desencolar().vertice)
	}
	return resultado
}

// OrdenTopologico devuelve el orden de ejecucion para grafos dirigidos.
// Retorna false si detecta un ciclo.
func OrdenTopologico[K comparable, W any](g grafo.Grafo[K, W], cmp func(K, K) bool) ([]K, bool) {
	grados := TDADiccionario.CrearHash[K, int](cmp)
	cola := TDACola.CrearColaEnlazada[K]()
	orden := []K{}

	vertices := g.Vertices()
	for _, v := range vertices {
		grados.Guardar(v, 0)
	}

	// Calcular grados de entrada
	for _, v := range vertices {
		for _, w := range g.ObtenerVecinos(v) {
			grados.Guardar(w, grados.Obtener(w)+1)
		}
	}

	// Encolar grado 0
	grados.Iterar(func(v K, grado int) bool {
		if grado == 0 {
			cola.Encolar(v)
		}
		return true
	})

	for !cola.EstaVacia() {
		v := cola.Desencolar()
		orden = append(orden, v)

		for _, w := range g.ObtenerVecinos(v) {
			nuevoGrado := grados.Obtener(w) - 1
			grados.Guardar(w, nuevoGrado)
			if nuevoGrado == 0 {
				cola.Encolar(w)
			}
		}
	}

	if len(orden) != len(vertices) {
		return nil, false // Hay ciclo
	}
	return orden, true
}

func reconstruirCamino[K comparable](padres TDADiccionario.Diccionario[K, K], origen, destino K) []K {
	camino := []K{}
	actual := destino

	for actual != origen {
		camino = append([]K{actual}, camino...)
		actual = padres.Obtener(actual)
	}
	camino = append([]K{origen}, camino...)
	return camino
}
