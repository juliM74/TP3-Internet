package grafo

import (
	TDADiccionario "tp3/tdaGrafo/tdas/hash/diccionario"
)

type grafoImp[K comparable, W any] struct {
	adyacentes TDADiccionario.Diccionario[K, TDADiccionario.Diccionario[K, W]]
}

// CrearGrafo crea un grafo no dirigido vacio.
func CrearGrafo[K comparable, W any]() Grafo[K, W] {
	return &grafoImp[K, W]{
		adyacentes: TDADiccionario.CrearHash[K, TDADiccionario.Diccionario[K, W]](),
	}
}

func (g *grafoImp[K, W]) AgregarVertice(v K) {
	if !g.adyacentes.Pertenece(v) {
		g.adyacentes.Guardar(v, TDADiccionario.CrearHash[K, W]())
	}
}

func (g *grafoImp[K, W]) AgregarArista(origen, destino K, datos W) {
	if !g.Pertenece(origen) || !g.Pertenece(destino) {
		panic("Alguno de los vertices no existe en el grafo")
	}
	g.adyacentes.Obtener(origen).Guardar(destino, datos)
	g.adyacentes.Obtener(destino).Guardar(origen, datos)
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

func (g *grafoImp[K, W]) BorrarArista(origen, destino K) {
	if g.adyacentes.Pertenece(origen) {
		vecinos := g.adyacentes.Obtener(origen)
		vecinos.Borrar(destino)
	}
	if g.adyacentes.Pertenece(destino) {
		vecinos := g.adyacentes.Obtener(destino)
		vecinos.Borrar(origen)
	}
}
