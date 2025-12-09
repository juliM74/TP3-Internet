package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func crearNodoLista[T any](nuevoDato T, apuntado *nodoLista[T]) *nodoLista[T] {
	return &nodoLista[T]{dato: nuevoDato, siguiente: apuntado}
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.primero == nil && l.ultimo == nil
}

func (l *listaEnlazada[T]) InsertarPrimero(dato T) {
	NuevoNodo := &nodoLista[T]{dato: dato, siguiente: l.primero}
	if l.EstaVacia() {
		l.ultimo = NuevoNodo
	}
	l.primero = NuevoNodo
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(dato T) {
	if l.EstaVacia() {
		l.InsertarPrimero(dato)
		return
	}
	NuevoNodo := &nodoLista[T]{dato: dato, siguiente: nil}
	l.ultimo.siguiente = NuevoNodo
	l.ultimo = NuevoNodo
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
	} else {
		return l.primero.dato
	}
}

func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	} else {
		return l.ultimo.dato
	}
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
