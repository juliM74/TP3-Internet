package grafo

// Grafo representa un grafo no dirigido generico con vertices de tipo K y datos en las aristas de tipo W.
type Grafo[K comparable, W any] interface {

	// AgregarVertice agrega un nuevo vertice al grafo.
	// Si el vertice ya existe, no hace nada.
	AgregarVertice(v K)

	// AgregarArista agrega una arista no dirigida para los vertices 'origen' y 'destino',
	// almacenando los datos especificados en la arista.
	// Si alguno de los vertices no pertenece al grafo, entra en panico.
	AgregarArista(origen, destino K, datos W)

	// ObtenerDatos devuelve los datos asociados a la arista entre 'origen' y 'destino'.
	// Si no existe la arista, el segundo valor devuelto sera false.
	ObtenerDatos(origen, destino K) (W, bool)

	// ObtenerVecinos devuelve un slice con los vertices adyacentes al vertice 'v'.
	// Si el vertice no pertenece al grafo, entra en panico.
	ObtenerVecinos(v K) []K

	// Pertenece indica si un vertice pertenece al grafo.
	Pertenece(v K) bool

	// Vertices devuelve un slice con todos los vertices del grafo.
	Vertices() []K

	// BorrarArista borra una arista del grafo.
	// Si la arista no existe, entra en panico.
	BorrarArista(origen, destino K)

	//BorrarVertice elimina el vertice del grafo.
	// Si el vertice no existe, entra en panico.
	BorrarVertice(vertice K)
}
