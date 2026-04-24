package diccionario

import (
	TDAPila "tp3/tdaGrafo/tdas/pila"
)

type nodoAbb[K comparable, V any] struct {
	izq   *nodoAbb[K, V]
	der   *nodoAbb[K, V]
	clave K
	dato  V
}

type abb[K comparable, V any] struct {
	raiz *nodoAbb[K, V]
	cant int
	cmp  func(K, K) int
}

type iterAbb[K comparable, V any] struct {
	abb   *abb[K, V]
	nodo  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
}

func crearNodoABB[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{
		clave: clave,
		dato:  dato,
	}
}

func (a *abb[K, V]) recorrer(clave K, accion func(**nodoAbb[K, V])) {
	nodo := &a.raiz
	for *nodo != nil {
		cmp := a.cmp(clave, (*nodo).clave)
		if cmp == 0 {
			break
		}
		if cmp < 0 {
			nodo = &(*nodo).izq
		} else {
			nodo = &(*nodo).der
		}
	}
	accion(nodo)
}

func (a *abb[K, V]) buscarNodo(clave K) *nodoAbb[K, V] {
	var res *nodoAbb[K, V]
	a.recorrer(clave, func(nodo **nodoAbb[K, V]) {
		res = *nodo
	})
	return res
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{
		raiz: nil,
		cant: 0,
		cmp:  funcion_cmp,
	}
}

func (a *abb[K, V]) Guardar(clave K, dato V) {
	a.recorrer(clave, func(nodo **nodoAbb[K, V]) {
		if *nodo == nil {
			a.cant++
			*nodo = crearNodoABB(clave, dato)
		} else {
			(*nodo).dato = dato
		}
	})
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	return a.buscarNodo(clave) != nil
}

func (a *abb[K, V]) Obtener(clave K) V {
	nodo := a.buscarNodo(clave)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return nodo.dato
}

func quitarMayor[K comparable, V any](nodo **nodoAbb[K, V]) *nodoAbb[K, V] {
	for (*nodo).der != nil {
		nodo = &(*nodo).der
	}
	mayor := *nodo
	*nodo = mayor.izq
	return mayor
}

func (abb *abb[K, V]) Borrar(clave K) V {
	var valorBorrado V
	var encontrado bool
	abb.recorrer(clave, func(ptrNodo **nodoAbb[K, V]) {
		if *ptrNodo == nil {
			return
		}
		encontrado = true
		valorBorrado = (*ptrNodo).dato
		switch {
		case (*ptrNodo).izq == nil && (*ptrNodo).der == nil:
			*ptrNodo = nil
		case (*ptrNodo).izq == nil:
			*ptrNodo = (*ptrNodo).der
		case (*ptrNodo).der == nil:
			*ptrNodo = (*ptrNodo).izq
		default:
			mayor := quitarMayor(&(*ptrNodo).izq)
			(*ptrNodo).clave = mayor.clave
			(*ptrNodo).dato = mayor.dato
		}
	})
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}
	abb.cant--
	return valorBorrado
}

func (a *abb[K, V]) Cantidad() int {
	return a.cant
}

func (a *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	a.IterarRango(nil, nil, visitar)
}
func (a *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	a._iterarRango(a.raiz, desde, hasta, visitar)
}

func (a *abb[K, V]) _iterarRango(nodo *nodoAbb[K, V], desde *K, hasta *K, visitar func(clave K, dato V) bool) bool {
	if nodo == nil {
		return true
	}
	if desde != nil && a.cmp(nodo.clave, *desde) < 0 {
		return a._iterarRango(nodo.der, desde, hasta, visitar)
	}

	if hasta != nil && a.cmp(nodo.clave, *hasta) > 0 {
		return a._iterarRango(nodo.izq, desde, hasta, visitar)
	}

	if !a._iterarRango(nodo.izq, desde, hasta, visitar) {
		return false
	}
	if !visitar(nodo.clave, nodo.dato) {
		return false
	}
	return a._iterarRango(nodo.der, desde, hasta, visitar)
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (nodo *nodoAbb[K, V]) apilarHijosIzquierdos(iter *iterAbb[K, V]) {
	for nodo != nil {
		if iter.hasta != nil && iter.abb.cmp(nodo.clave, *iter.hasta) > 0 {
			nodo = nodo.izq
			continue
		}
		if iter.desde != nil && iter.abb.cmp(nodo.clave, *iter.desde) < 0 {
			nodo = nodo.der
			continue
		}
		iter.nodo.Apilar(nodo)
		nodo = nodo.izq
	}
}

func (iter *iterAbb[K, V]) HaySiguiente() bool {
	return !iter.nodo.EstaVacia()
}

func (iter *iterAbb[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodoActual := iter.nodo.VerTope()
	return nodoActual.clave, nodoActual.dato
}

func (iter *iterAbb[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	actual := iter.nodo.Desapilar()
	if actual.der != nil {
		actual.der.apilarHijosIzquierdos(iter)
	}
}

func (abb *abb[K, V]) crearIterador() *iterAbb[K, V] {
	iter := &iterAbb[K, V]{
		abb:  abb,
		nodo: TDAPila.CrearPilaDinamica[*nodoAbb[K, V]](),
	}
	return iter
}

func (a *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := a.crearIterador()
	iter.desde = desde
	iter.hasta = hasta
	a.raiz.apilarHijosIzquierdos(iter)
	return iter
}
