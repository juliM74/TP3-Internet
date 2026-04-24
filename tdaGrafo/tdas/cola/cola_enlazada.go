package cola

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

type colaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{}
}

func crearNodo[T any](nuevoElemento T, apuntado *nodo[T]) *nodo[T] {
	return &nodo[T]{dato: nuevoElemento, siguiente: apuntado}
}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.primero.dato
}

func (c *colaEnlazada[T]) Encolar(nuevoElemento T) {
	nuevoNodo := crearNodo(nuevoElemento, nil)
	if c.EstaVacia() { // cuando la cola esta vacia, se agrega el primer nodo (ambos punteros de la cola apuntan al primer nodo)
		c.primero = nuevoNodo
	} else { // como la cola tiene elementos, apuntamos el ultimo nodo al nuevo, y el puntero c.ultimo ahora apunta al nuevo nodo
		c.ultimo.siguiente = nuevoNodo
	}
	c.ultimo = nuevoNodo
}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	desencolado := c.primero.dato
	c.primero = c.primero.siguiente // el primer puntero de cola apunta al siguiente nodo del nodo desencolado
	if c.EstaVacia() {
		c.ultimo = nil
	}
	return desencolado
}
