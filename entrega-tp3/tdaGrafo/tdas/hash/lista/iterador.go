package lista

type iterListaEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

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
	nuevoNodo := crearNodoLista(nuevoElemento, nil)

	if i.lista.EstaVacia() {
		i.lista.primero = nuevoNodo
		i.lista.ultimo = nuevoNodo
		i.actual = nuevoNodo
	} else if i.anterior == nil {
		nuevoNodo.siguiente = i.lista.primero
		i.lista.primero = nuevoNodo
		i.actual = nuevoNodo
	} else if i.actual == nil {
		i.lista.ultimo.siguiente = nuevoNodo
		i.lista.ultimo = nuevoNodo
		i.actual = nuevoNodo
	} else {
		nuevoNodo.siguiente = i.actual
		i.anterior.siguiente = nuevoNodo
		i.actual = nuevoNodo
	}

	i.lista.largo++
}

func (i *iterListaEnlazada[T]) Borrar() T {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	borrado := i.actual.dato

	if i.anterior == nil {
		i.lista.primero = i.actual.siguiente
		if i.lista.primero == nil {
			i.lista.ultimo = nil
		}
	} else {
		i.anterior.siguiente = i.actual.siguiente
		if i.actual.siguiente == nil {
			i.lista.ultimo = i.anterior
		}
	}

	i.actual = i.actual.siguiente
	i.lista.largo--

	return borrado
}
