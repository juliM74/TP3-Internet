package biblioteca

import (
	grafo "tp3/tdaGrafo"
	TDACola "tp3/tdaGrafo/tdas/cola"
	TDAHeap "tp3/tdaGrafo/tdas/cola_prioridad"
	TDADiccionario "tp3/tdaGrafo/tdas/hash/diccionario"
	TDAPila "tp3/tdaGrafo/tdas/pila"
)

// ========================= ESTRUCTURAS AUXILIARES =========================

type elementoDistancia[K comparable] struct {
	vertice   K
	distancia int
}

type aristaHeap[K comparable] struct {
	origen  K
	destino K
	peso    int
}

type Arista[K comparable] struct {
	Origen  K
	Destino K
	Peso    int
}

type itemRanking[K comparable, V comparable] struct {
	clave K
	valor V
}

// ========================= ALGORITMOS DE CAMINOS =========================

// CaminoMinimo aplica Dijkstra. Devuelve el camino y su costo total.
// Si no hay camino, devuelve nil y -1.
func CaminoMinimo[K comparable, W any](g grafo.Grafo[K, W], origen, destino K, obtenerPeso func(W) int, cmp func(K, K) bool) ([]K, int) {
	dist := TDADiccionario.CrearHash[K, int](cmp)
	padre := TDADiccionario.CrearHash[K, K](cmp)
	visitados := TDADiccionario.CrearHash[K, bool](cmp)

	// Heap de Minimos
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
	return nil, -1
}

// CaminoMinimoBFS encuentra el camino con menos aristas
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

// MSTPrim aplica el algoritmo de Prim para encontrar el Arbol de Tendido Minimo.
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

	// Carga inicial de aristas
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

		aristasMST = append(aristasMST, Arista[K]{
			Origen: arista.origen, Destino: arista.destino, Peso: arista.peso,
		})
		visitados.Guardar(arista.destino, true)

		for _, w := range g.ObtenerVecinos(arista.destino) {
			if !visitados.Pertenece(w) {
				dato, _ := g.ObtenerDatos(arista.destino, w)
				h.Encolar(aristaHeap[K]{
					origen: arista.destino, destino: w, peso: obtenerPeso(dato),
				})
			}
		}
	}
	return aristasMST
}

// ========================= ALGORITMOS DE ORDEN Y RANKING =========================

// OrdenTopologico devuelve un orden de ejecucion valido. Retorna false si hay ciclo.
func OrdenTopologico[K comparable, W any](g grafo.Grafo[K, W], cmp func(K, K) bool) ([]K, bool) {
	grados := TDADiccionario.CrearHash[K, int](cmp)
	cola := TDACola.CrearColaEnlazada[K]()
	orden := []K{}
	vertices := g.Vertices()

	for _, v := range vertices {
		grados.Guardar(v, 0)
	}
	for _, v := range vertices {
		for _, w := range g.ObtenerVecinos(v) {
			grados.Guardar(w, grados.Obtener(w)+1)
		}
	}
	for _, v := range vertices {
		if grados.Obtener(v) == 0 {
			cola.Encolar(v)
		}
	}

	for !cola.EstaVacia() {
		v := cola.Desencolar()
		orden = append(orden, v)
		for _, w := range g.ObtenerVecinos(v) {
			nuevo := grados.Obtener(w) - 1
			grados.Guardar(w, nuevo)
			if nuevo == 0 {
				cola.Encolar(w)
			}
		}
	}
	if len(orden) != len(vertices) {
		return nil, false
	}
	return orden, true
}

// PageRank calcula la importancia de cada pagina
func PageRank[K comparable, W any](g grafo.Grafo[K, W], d float64, iteraciones int, _ func(K, K) bool) TDADiccionario.Diccionario[K, float64] {
	vertices := g.Vertices()
	n := float64(len(vertices))
	pr := TDADiccionario.CrearHash[K, float64](sonIguales)
	base := (1.0 - d) / n

	for _, v := range vertices {
		pr.Guardar(v, 1.0/n)
	}

	for i := 0; i < iteraciones; i++ {
		prAux := TDADiccionario.CrearHash[K, float64](sonIguales)
		for _, v := range vertices {
			prAux.Guardar(v, base)
		}

		for _, v := range vertices {
			if !pr.Pertenece(v) {
				continue
			}
			vecinos := g.ObtenerVecinos(v)
			cant := float64(len(vecinos))

			if cant > 0 {
				aporte := (pr.Obtener(v) * d) / cant
				for _, w := range vecinos {
					if prAux.Pertenece(w) {
						prAux.Guardar(w, prAux.Obtener(w)+aporte)
					}
				}
			}
		}
		pr = prAux
	}
	return pr
}

// TopN obtiene los N elementos mayores usando un iterador y un comparador.
func TopN[K comparable, V comparable](iterador func(func(K, V) bool), n int, cmp func(V, V) int) []K {
	h := TDAHeap.CrearHeap(func(a, b itemRanking[K, V]) int {
		return cmp(b.valor, a.valor) // Min-Heap
	})

	iterador(func(clave K, valor V) bool {
		dato := itemRanking[K, V]{clave, valor}
		if h.Cantidad() < n {
			h.Encolar(dato)
		} else if cmp(valor, h.VerMax().valor) > 0 {
			h.Desencolar()
			h.Encolar(dato)
		}
		return true
	})

	resultado := make([]K, h.Cantidad())
	for i := h.Cantidad() - 1; i >= 0; i-- {
		resultado[i] = h.Desencolar().clave
	}
	return resultado
}

// ========================= CONECTIVIDAD Y COMUNIDADES =========================

// CFCSoloDe devuelve la componente fuertemente conexa de un vertice especifico.
func CFCSoloDe[K comparable, W any](g grafo.Grafo[K, W], origen K, cmp func(K, K) bool) []K {
	todas := ComponentesFuertementeConexas(g, cmp)
	for _, comp := range todas {
		for _, v := range comp {
			if cmp(v, origen) {
				return comp
			}
		}
	}
	return nil
}

// ComponentesFuertementeConexas calcula todas las componentes fuertemente conexas del grafo.
// Utiliza el algoritmo de Tarjan con comparacion estricta para evitar conflictos de Hash.
func ComponentesFuertementeConexas[K comparable, W any](g grafo.Grafo[K, W], _ func(K, K) bool) [][]K {
	visitados := TDADiccionario.CrearHash[K, int](sonIguales)
	masBajo := TDADiccionario.CrearHash[K, int](sonIguales)
	enPila := TDADiccionario.CrearHash[K, bool](sonIguales)

	pila := TDAPila.CrearPilaDinamica[K]()
	var componentes [][]K
	ordenGlobal := 0

	var dfs func(K)
	dfs = func(v K) {
		visitados.Guardar(v, ordenGlobal)
		masBajo.Guardar(v, ordenGlobal)
		enPila.Guardar(v, true)
		pila.Apilar(v)
		ordenGlobal++

		for _, w := range g.ObtenerVecinos(v) {
			if !visitados.Pertenece(w) {
				dfs(w)
				if masBajo.Obtener(w) < masBajo.Obtener(v) {
					masBajo.Guardar(v, masBajo.Obtener(w))
				}
			} else if enPila.Pertenece(w) {
				if visitados.Obtener(w) < masBajo.Obtener(v) {
					masBajo.Guardar(v, visitados.Obtener(w))
				}
			}
		}

		if masBajo.Obtener(v) == visitados.Obtener(v) {
			nuevaCFC := []K{}
			for {
				tope := pila.Desapilar()
				enPila.Borrar(tope)
				nuevaCFC = append(nuevaCFC, tope)
				if sonIguales(tope, v) {
					break
				}
			}
			componentes = append(componentes, nuevaCFC)
		}
	}

	for _, v := range g.Vertices() {
		if !visitados.Pertenece(v) {
			dfs(v)
		}
	}
	return componentes
}

// LabelPropagation detecta comunidades propagando etiquetas mayoritarias.
func LabelPropagation[K comparable, W any](g grafo.Grafo[K, W], cmp func(K, K) bool) TDADiccionario.Diccionario[K, int] {
	vertices := g.Vertices()
	etiquetas := TDADiccionario.CrearHash[K, int](cmp)

	for i, v := range vertices {
		etiquetas.Guardar(v, i)
	}

	for i := 0; i < 10; i++ { // Max 10 iteraciones
		cambios := false
		for _, v := range vertices {
			vecinos := g.ObtenerVecinos(v)
			if len(vecinos) == 0 {
				continue
			}

			frecuencias := make(map[int]int)
			for _, w := range vecinos {
				if etiquetas.Pertenece(w) {
					frecuencias[etiquetas.Obtener(w)]++
				}
			}

			maxFrec := -1
			mejorEtiqueta := -1
			for etiq, frec := range frecuencias {
				if frec > maxFrec {
					maxFrec = frec
					mejorEtiqueta = etiq
				}
			}

			if mejorEtiqueta != -1 {
				actual := etiquetas.Obtener(v)
				if mejorEtiqueta != actual {
					etiquetas.Guardar(v, mejorEtiqueta)
					cambios = true
				}
			}
		}
		if !cambios {
			break
		}
	}
	return etiquetas
}

// ========================= METRICAS =========================

// ClusteringVertice calcula coeficiente local.
func ClusteringVertice[K comparable, W any](g grafo.Grafo[K, W], v K, cmp func(K, K) bool) float64 {
	vecinos := g.ObtenerVecinos(v)
	k := len(vecinos)
	if k < 2 {
		return 0.0
	}

	vecinosSet := TDADiccionario.CrearHash[K, bool](cmp)
	for _, vec := range vecinos {
		vecinosSet.Guardar(vec, true)
	}

	aristas := 0
	for _, w := range vecinos {
		for _, x := range g.ObtenerVecinos(w) {
			if cmp(w, x) {
				continue
			}
			if vecinosSet.Pertenece(x) {
				aristas++
			}
		}
	}
	return float64(aristas) / float64(k*(k-1))
}

// ClusteringPromedio calcula promedio global de la red.
func ClusteringPromedio[K comparable, W any](g grafo.Grafo[K, W], cmp func(K, K) bool) float64 {
	vertices := g.Vertices()
	n := len(vertices)
	if n == 0 {
		return 0.0
	}

	suma := 0.0
	for _, v := range vertices {
		suma += ClusteringVertice(g, v, cmp)
	}
	return suma / float64(n)
}

// Diametro calcula el camino minimo mas largo. O(V*(V+E)).
func Diametro[K comparable, W any](g grafo.Grafo[K, W], cmp func(K, K) bool) []K {
	var maxCamino []K
	maxDistGlobal := -1

	for _, origen := range g.Vertices() {
		_, padres, ultimo, distMaxLocal := bfsCompleto(g, origen, cmp)

		if distMaxLocal > maxDistGlobal {
			maxDistGlobal = distMaxLocal
			maxCamino = reconstruirCamino(padres, origen, ultimo)
		}
	}
	return maxCamino
}

// CantidadEnRango cuenta nodos a distancia exacta n (BFS por niveles).
func CantidadEnRango[K comparable, W any](g grafo.Grafo[K, W], origen K, n int, cmp func(K, K) bool) int {
	if n < 0 {
		return 0
	}
	visitados := TDADiccionario.CrearHash[K, bool](cmp)
	cola := TDACola.CrearColaEnlazada[elementoDistancia[K]]()

	visitados.Guardar(origen, true)
	cola.Encolar(elementoDistancia[K]{vertice: origen, distancia: 0})
	cantidad := 0

	for !cola.EstaVacia() {
		actual := cola.Desencolar()
		if actual.distancia == n {
			cantidad++
			continue
		}
		for _, w := range g.ObtenerVecinos(actual.vertice) {
			if !visitados.Pertenece(w) {
				visitados.Guardar(w, true)
				cola.Encolar(elementoDistancia[K]{vertice: w, distancia: actual.distancia + 1})
			}
		}
	}
	return cantidad
}

// ========================= NAVEGACION Y LECTURA =========================

// CicloLargoN busca ciclo de largo exacto N (Backtracking).
func CicloLargoN[K comparable, W any](g grafo.Grafo[K, W], origen K, n int, cmp func(K, K) bool) []K {
	visitados := TDADiccionario.CrearHash[K, bool](cmp)
	camino := []K{origen}
	visitados.Guardar(origen, true)
	return backtrackingCiclo(g, origen, origen, n, camino, visitados, cmp)
}

func backtrackingCiclo[K comparable, W any](g grafo.Grafo[K, W], inicio, actual K, n int, camino []K, visitados TDADiccionario.Diccionario[K, bool], cmp func(K, K) bool) []K {
	if len(camino) == n+1 {
		return nil
	}

	for _, w := range g.ObtenerVecinos(actual) {
		if len(camino) == n && cmp(w, inicio) {
			return append(camino, w)
		}
		if !visitados.Pertenece(w) {
			visitados.Guardar(w, true)
			res := backtrackingCiclo(g, inicio, w, n, append(camino, w), visitados, cmp)
			if res != nil {
				return res
			}
			visitados.Borrar(w)
		}
	}
	return nil
}

// Lectura2am: Orden Topologico sobre un subconjunto de vertices.
func Lectura2am[K comparable, W any](g grafo.Grafo[K, W], paginas []K, cmp func(K, K) bool) []K {
	grados := TDADiccionario.CrearHash[K, int](cmp)
	enSubconjunto := TDADiccionario.CrearHash[K, bool](cmp)
	cola := TDACola.CrearColaEnlazada[K]()

	for _, p := range paginas {
		grados.Guardar(p, 0)
		enSubconjunto.Guardar(p, true)
	}
	// Solo aristas internas
	for _, v := range paginas {
		for _, w := range g.ObtenerVecinos(v) {
			if enSubconjunto.Pertenece(w) {
				grados.Guardar(w, grados.Obtener(w)+1)
			}
		}
	}
	for _, p := range paginas {
		if grados.Obtener(p) == 0 {
			cola.Encolar(p)
		}
	}

	orden := []K{}
	for !cola.EstaVacia() {
		v := cola.Desencolar()
		orden = append(orden, v)
		for _, w := range g.ObtenerVecinos(v) {
			if enSubconjunto.Pertenece(w) {
				nuevo := grados.Obtener(w) - 1
				grados.Guardar(w, nuevo)
				if nuevo == 0 {
					cola.Encolar(w)
				}
			}
		}
	}
	if len(orden) != len(paginas) {
		return nil
	}
	invertirLista(orden)
	return orden
}

// PrimerLink navega siempre al primer vecino (max 20 saltos).
func PrimerLink[K comparable, W any](g grafo.Grafo[K, W], origen K, cmp func(K, K) bool) []K {
	camino := []K{origen}
	actual := origen
	visitados := TDADiccionario.CrearHash[K, bool](cmp)

	for i := 0; i < 20; i++ {
		visitados.Guardar(actual, true)
		vecinos := g.ObtenerVecinos(actual)
		if len(vecinos) == 0 {
			break
		}

		siguiente := vecinos[0]
		if visitados.Pertenece(siguiente) {
			break
		}

		camino = append(camino, siguiente)
		actual = siguiente
	}
	return camino
}

// ========================= AUXILIARES PRIVADOS =========================

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

// bfsCompleto: Auxiliar para Diametro. No corta al destino, explora todo.
func bfsCompleto[K comparable, W any](g grafo.Grafo[K, W], origen K, cmp func(K, K) bool) (TDADiccionario.Diccionario[K, int], TDADiccionario.Diccionario[K, K], K, int) {
	dist := TDADiccionario.CrearHash[K, int](cmp)
	padre := TDADiccionario.CrearHash[K, K](cmp)
	cola := TDACola.CrearColaEnlazada[K]()

	dist.Guardar(origen, 0)
	cola.Encolar(origen)

	var ultimo K = origen
	maxDist := 0

	for !cola.EstaVacia() {
		v := cola.Desencolar()
		dActual := dist.Obtener(v)
		if dActual > maxDist {
			maxDist = dActual
			ultimo = v
		}

		for _, w := range g.ObtenerVecinos(v) {
			if !dist.Pertenece(w) {
				dist.Guardar(w, dActual+1)
				padre.Guardar(w, v)
				cola.Encolar(w)
			}
		}
	}
	return dist, padre, ultimo, maxDist
}

// Funcion auxiliar para uso interno de los algoritmos.
// Garantiza consistencia con el Grafo.
func sonIguales[K comparable](a, b K) bool {
	return a == b
}

func invertirLista[K any](lista []K) {
	for i, j := 0, len(lista)-1; i < j; i, j = i+1, j-1 {
		lista[i], lista[j] = lista[j], lista[i]
	}
}
