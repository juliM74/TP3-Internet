package lista

type IteradorLista[T any] interface {
	// VerActual devuelve el valor del nodo actual del iterador.
	// Pre: HaySiguiente() es true.
	VerActual() T

	// HaySiguiente indica si el iterador no ha alcanzado el final de la lista.
	HaySiguiente() bool

	// Siguiente avanza el iterador al próximo nodo de la lista.
	// Pre: HaySiguiente() es true.
	Siguiente()

	// Insertar agrega un nuevo elemento en la posición actual del iterador.
	// Pre: el iterador fue creado.
	// Post: el nuevo elemento está en la posición actual y el iterador apunta a él.
	Insertar(T)

	// Borrar elimina el nodo actual de la lista y devuelve su valor.
	// Pre: HaySiguiente() es true.
	// Post: el iterador apunta al siguiente nodo.
	Borrar() T
}

type Lista[T any] interface {
	// EstaVacia devuelve true si la lista no contiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero inserta un nuevo elemento al principio de la lista.
	// Pre: la lista fue creada.
	// Post: el elemento fue insertado al inicio de la lista.
	InsertarPrimero(T)

	// InsertarUltimo inserta un nuevo elemento al final de la lista.
	// Pre: la lista fue creada.
	// Post: el elemento fue insertado al final de la lista.
	InsertarUltimo(T)

	// elimina el primer elemento de la lista y lo devuelve.
	// Pre: la lista no debe estar vacía.
	// Post: la cantidad de elementos disminuyó en 1. El primer elemento de la lista es ahora el segundo anterior.
	// Si la lista está vacía, entra en pánico con el mensaje: "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero devuelve el primer elemento de la lista sin eliminarlo.
	// Pre: la lista no debe estar vacía.
	// Post: no modifica la lista.
	// Si la lista está vacía, entra en pánico con el mensaje: "La lista esta vacia".
	VerPrimero() T

	// VerUltimo devuelve el último elemento de la lista sin eliminarlo.
	// Pre: la lista no debe estar vacía.
	// Post: no modifica la lista.
	// Si la lista está vacía, entra en pánico con el mensaje: "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos en la lista.
	Largo() int

	// Iterar aplica la función visitar a cada elemento de la lista, desde el primero hasta el último.
	// La iteración puede cortarse si visitar devuelve false.
	Iterar(visitar func(T) bool)

	// Iterador devuelve un nuevo iterador externo para recorrer o modificar la lista.
	Iterador() IteradorLista[T]
}
