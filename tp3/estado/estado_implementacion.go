package estado

import (
	"strings"
	grafo "tp3/tdaGrafo"
	TDADiccionario "tp3/tdaGrafo/tdas/hash/diccionario"
)

// =============================
//   IMPLEMENTACIÓN DEL TDA ESTADO
// =============================

type estadoConcreto struct {
	grafo grafo.Grafo[string, int]

	// PageRank
	pagerank          TDADiccionario.Diccionario[string, float64]
	pagerankCalculado bool

	// CFC cache
	cfcCache TDADiccionario.Diccionario[string, []string]

	// Comunidades
	etiquetas             TDADiccionario.Diccionario[string, int]
	comunidadesCalculadas bool

	// Clustering
	clusteringLocal             TDADiccionario.Diccionario[string, float64]
	clusteringPromedio          float64
	clusteringPromedioCalculado bool
}

// Constructor
func NuevoEstado(g grafo.Grafo[string, int]) Estado {
	return &estadoConcreto{
		grafo:           g,
		pagerank:        TDADiccionario.CrearHash[string, float64](strings.EqualFold),
		cfcCache:        TDADiccionario.CrearHash[string, []string](strings.EqualFold),
		etiquetas:       TDADiccionario.CrearHash[string, int](strings.EqualFold),
		clusteringLocal: TDADiccionario.CrearHash[string, float64](strings.EqualFold),
	}
}

// =============================
//   MÉTODOS DE LA INTERFAZ
// =============================

// -------- GRAFO --------
func (estado *estadoConcreto) Grafo() grafo.Grafo[string, int] {
	return estado.grafo
}

// -------- PAGE RANK --------
func (estado *estadoConcreto) TienePagerank() bool {
	return estado.pagerankCalculado
}

func (estado *estadoConcreto) GuardarPagerank(pagina string, valor float64) {
	estado.pagerank.Guardar(pagina, valor)
}

func (estado *estadoConcreto) ObtenerPagerank(pagina string) (float64, bool) {
	if !estado.pagerank.Pertenece(pagina) {
		return 0, false
	}
	return estado.pagerank.Obtener(pagina), true
}

func (estado *estadoConcreto) IterarPagerank(visitar func(string, float64)) {
	estado.pagerank.Iterar(func(k string, v float64) bool {
		visitar(k, v)
		return true
	})
}

func (estado *estadoConcreto) MarcarPagerankCalculado() {
	estado.pagerankCalculado = true
}

// -------- CFC CACHE --------
func (estado *estadoConcreto) TieneCFC(pagina string) bool {
	return estado.cfcCache.Pertenece(pagina)
}

func (estado *estadoConcreto) GuardarCFC(pagina string, cfc []string) {
	estado.cfcCache.Guardar(pagina, cfc)
}

func (estado *estadoConcreto) ObtenerCFC(pagina string) []string {
	if !estado.cfcCache.Pertenece(pagina) {
		return nil
	}
	return estado.cfcCache.Obtener(pagina)
}

// -------- COMUNIDADES --------
func (estado *estadoConcreto) TieneComunidades() bool {
	return estado.comunidadesCalculadas
}

func (estado *estadoConcreto) GuardarEtiqueta(pagina string, etiqueta int) {
	estado.etiquetas.Guardar(pagina, etiqueta)
}

func (estado *estadoConcreto) ObtenerEtiqueta(pagina string) (int, bool) {
	if !estado.etiquetas.Pertenece(pagina) {
		return 0, false
	}
	return estado.etiquetas.Obtener(pagina), true
}

func (estado *estadoConcreto) IterarEtiquetas(visitar func(string, int)) {
	estado.etiquetas.Iterar(func(k string, v int) bool {
		visitar(k, v)
		return true
	})
}

func (estado *estadoConcreto) MarcarComunidadesCalculadas() {
	estado.comunidadesCalculadas = true
}

// -------- CLUSTERING LOCAL --------
func (estado *estadoConcreto) TieneClusteringLocal(pagina string) bool {
	return estado.clusteringLocal.Pertenece(pagina)
}

func (estado *estadoConcreto) GuardarClusteringLocal(pagina string, valor float64) {
	estado.clusteringLocal.Guardar(pagina, valor)
}

func (estado *estadoConcreto) ObtenerClusteringLocal(pagina string) float64 {
	return estado.clusteringLocal.Obtener(pagina)
}

// -------- CLUSTERING PROMEDIO --------
func (estado *estadoConcreto) TieneClusteringPromedio() bool {
	return estado.clusteringPromedioCalculado
}

func (estado *estadoConcreto) GuardarClusteringPromedio(valor float64) {
	estado.clusteringPromedio = valor
	estado.clusteringPromedioCalculado = true
}

func (estado *estadoConcreto) ObtenerClusteringPromedio() float64 {
	return estado.clusteringPromedio
}
