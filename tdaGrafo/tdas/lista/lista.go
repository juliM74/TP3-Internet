package lista

type IteradorLista[T any] interface {
	// VerActual devuelve el valor del elemento actual del iterador.
	// Si el iterador ya termino (es decir, no hay un elemento actual), entra en panico con un mensaje "El iterador termino de iterar".
	VerActual() T

	// HaySiguiente indica si el iterador no ha alcanzado el final de la lista.
	HaySiguiente() bool

	// Siguiente avanza el iterador al próximo elemento de la lista.
	// Si el iterador ya termino, entra en panico con un mensaje "El iterador termino de iterar".
	Siguiente()

	// Insertar agrega un nuevo elemento en la posicion actual del iterador.
	// El nuevo elemento se inserta antes del elemento actual y el iterador pasa a apuntar a el.
	Insertar(T)

	// Borrar elimina el nodo actual de la lista y devuelve su valor.
	// Si el iterador ya termino, entra en panico con un mensaje "El iterador termino de iterar".
	Borrar() T
}

type Lista[T any] interface {
	// EstaVacia devuelve true si la lista no contiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero inserta un nuevo elemento al principio de la lista.
	InsertarPrimero(T)

	// InsertarUltimo inserta un nuevo elemento al final de la lista.
	InsertarUltimo(T)

	// Elimina el primer elemento de la lista y lo devuelve.
	// Si la lista está vacía, entra en pánico con el mensaje: "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero devuelve el primer elemento de la lista sin eliminarlo.a.
	// Si la lista está vacía, entra en pánico con el mensaje: "La lista esta vacia".
	VerPrimero() T

	// VerUltimo devuelve el último elemento de la lista sin eliminarlo.
	// Si la lista está vacía, entra en pánico con el mensaje: "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos en la lista.
	Largo() int

	// Iterar aplica la función visitar a cada elemento de la lista, desde el primero hasta el último.
	// La iteración se detiene si la función visitar devuelve false
	Iterar(visitar func(T) bool)

	// Iterador devuelve un nuevo iterador externo para recorrer o modificar la lista.
	Iterador() IteradorLista[T]
}
