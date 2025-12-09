package lista

// Estructura de nodo (privada)
type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

// Estructura de la lista enlazada
type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

// Estructura de iterador
type iterListaEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

// Metodos de listaEnlazada

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func crearNodoLista[T any](nuevoDato T, apuntado *nodoLista[T]) *nodoLista[T] {
	return &nodoLista[T]{dato: nuevoDato, siguiente: apuntado}
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.primero == nil && l.ultimo == nil && l.largo == 0
}

func (l *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevoNodo := crearNodoLista(dato, l.primero)
	if l.EstaVacia() {
		l.ultimo = nuevoNodo
	}
	l.primero = nuevoNodo
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevoNodo := crearNodoLista(dato, nil)
	if l.EstaVacia() {
		l.primero = nuevoNodo
	} else {
		l.ultimo.siguiente = nuevoNodo
	}
	l.ultimo = nuevoNodo
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	dato := l.primero.dato
	l.primero = l.primero.siguiente
	if l.primero == nil {
		l.ultimo = nil
	}
	l.largo--
	return dato
}

func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.primero.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.ultimo.dato
}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	puntero := l.primero
	for puntero != nil {
		if !visitar(puntero.dato) {
			return
		}
		puntero = puntero.siguiente
	}
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{
		actual:   l.primero,
		anterior: nil,
		lista:    l,
	}
}

// Metodos del iterador externo

func (i *iterListaEnlazada[T]) VerActual() T {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	return i.actual.dato
}

func (i *iterListaEnlazada[T]) HaySiguiente() bool {
	return i.actual != nil
}

func (i *iterListaEnlazada[T]) Siguiente() {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	i.anterior = i.actual
	i.actual = i.actual.siguiente
}

func (i *iterListaEnlazada[T]) Insertar(nuevoElemento T) {
	nuevoNodo := crearNodoLista(nuevoElemento, i.actual)

	if i.anterior == nil {
		i.lista.primero = nuevoNodo
	} else {
		i.anterior.siguiente = nuevoNodo
	}
	if i.actual == nil {
		i.lista.ultimo = nuevoNodo
	}

	i.actual = nuevoNodo
	i.lista.largo++
}

func (i *iterListaEnlazada[T]) Borrar() T {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	borrado := i.actual.dato

	if i.anterior == nil {
		i.lista.primero = i.actual.siguiente
	} else {
		i.anterior.siguiente = i.actual.siguiente
	}
	if i.actual.siguiente == nil {
		i.lista.ultimo = i.anterior
	}

	i.actual = i.actual.siguiente
	i.lista.largo--

	return borrado
}
