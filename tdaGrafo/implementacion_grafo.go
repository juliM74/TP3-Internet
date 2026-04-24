package grafo

import (
	TDADiccionario "tp3/tdaGrafo/tdas/hash/diccionario"
)

type grafoImp[K comparable, W any] struct {
	adyacentes TDADiccionario.Diccionario[K, TDADiccionario.Diccionario[K, W]]
	esDirigido bool
}

func SonIguales[K comparable](a, b K) bool {
	return a == b
}

func CrearGrafo[K comparable, W any](esDirigido bool) Grafo[K, W] {
	return &grafoImp[K, W]{
		adyacentes: TDADiccionario.CrearHash[K, TDADiccionario.Diccionario[K, W]](SonIguales),
		esDirigido: esDirigido,
	}
}

func (g *grafoImp[K, W]) AgregarVertice(v K) {
	if !g.adyacentes.Pertenece(v) {
		g.adyacentes.Guardar(v, TDADiccionario.CrearHash[K, W](SonIguales))
	}
}

func (g *grafoImp[K, W]) BorrarVertice(vertice K) {
	if !g.Pertenece(vertice) {
		panic("No existe tal vertice en el grafo")
	}

	// SI ES NO DIRIGIDO
	if !g.esDirigido {
		vecinos := g.ObtenerVecinos(vertice)
		for _, vecino := range vecinos {
			g.adyacentes.Obtener(vecino).Borrar(vertice)
		}
	} else {
		vertices := g.Vertices()
		for _, v := range vertices {
			if v != vertice {
				adyacentesDeV := g.adyacentes.Obtener(v)
				if adyacentesDeV.Pertenece(vertice) {
					adyacentesDeV.Borrar(vertice)
				}
			}
		}
	}
	g.adyacentes.Borrar(vertice)
}

func (g *grafoImp[K, W]) AgregarArista(origen, destino K, datos W) {
	if !g.Pertenece(origen) || !g.Pertenece(destino) {
		panic("Alguno de los vertices no existe en el grafo")
	}
	g.adyacentes.Obtener(origen).Guardar(destino, datos)
	if !g.esDirigido {
		g.adyacentes.Obtener(destino).Guardar(origen, datos)
	}
}

func (g *grafoImp[K, W]) BorrarArista(origen, destino K) {
	if !g.adyacentes.Pertenece(origen) || !g.adyacentes.Pertenece(destino) {
		panic("No existe tal origen o destino en el grafo")
	}
	g.adyacentes.Obtener(origen).Borrar(destino)
	if !g.esDirigido {
		g.adyacentes.Obtener(destino).Borrar(origen)
	}
}

func (g *grafoImp[K, W]) ObtenerDatos(origen, destino K) (W, bool) {
	if !g.Pertenece(origen) {
		var cero W
		return cero, false
	}
	vecinos := g.adyacentes.Obtener(origen)
	if !vecinos.Pertenece(destino) {
		var cero W
		return cero, false
	}
	return vecinos.Obtener(destino), true
}

func (g *grafoImp[K, W]) ObtenerVecinos(v K) []K {
	if !g.Pertenece(v) {
		panic("El vertice no pertenece al grafo")
	}
	vecinos := []K{}
	g.adyacentes.Obtener(v).Iterar(func(clave K, _ W) bool {
		vecinos = append(vecinos, clave)
		return true
	})
	return vecinos
}

func (g *grafoImp[K, W]) Pertenece(v K) bool {
	return g.adyacentes.Pertenece(v)
}

func (g *grafoImp[K, W]) Vertices() []K {
	vertices := []K{}
	g.adyacentes.Iterar(func(clave K, _ TDADiccionario.Diccionario[K, W]) bool {
		vertices = append(vertices, clave)
		return true
	})
	return vertices
}
