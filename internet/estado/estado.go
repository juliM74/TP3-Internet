package estado

import (
	grafo "tp3/tdaGrafo"
)

// Interfaz del TDA Estado.
// Maneja todo el estado global del programa, cacheos y acceso al grafo.
type Estado interface {

	// Acceso al grafo principal
	Grafo() grafo.Grafo[string, int]

	// PageRank
	TienePagerank() bool
	GuardarPagerank(pagina string, valor float64)
	ObtenerPagerank(pagina string) (float64, bool)
	IterarPagerank(visitar func(string, float64) bool)
	MarcarPagerankCalculado()

	// Componentes Fuertemente Conectadas (CFC)
	TieneCFC(pagina string) bool
	GuardarCFC(pagina string, cfc []string)
	ObtenerCFC(pagina string) []string

	// Comunidades
	TieneComunidades() bool
	GuardarEtiqueta(pagina string, etiqueta int)
	ObtenerEtiqueta(pagina string) (int, bool)
	IterarEtiquetas(visitar func(string, int))
	MarcarComunidadesCalculadas()

	// Clustering local
	TieneClusteringLocal(pagina string) bool
	GuardarClusteringLocal(pagina string, valor float64)
	ObtenerClusteringLocal(pagina string) float64

	// Clustering global
	TieneClusteringPromedio() bool
	GuardarClusteringPromedio(valor float64)
	ObtenerClusteringPromedio() float64
}
