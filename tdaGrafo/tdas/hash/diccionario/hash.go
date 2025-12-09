package diccionario

import (
	"fmt"
	"hash/fnv"
	TDAlista "tp3/tdaGrafo/tdas/lista"
)

const (
	CAPACIDAD_INICIAL      = 53
	FACTOR_CARGA_AUMENTO   = 2.0
	FACTOR_CARGA_REDUCCION = 0.5
	FACTOR_REDIMENSION     = 2
)

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	tabla     []TDAlista.Lista[parClaveValor[K, V]]
	cantidad  int // cantidad real de elementos guardados
	capacidad int // cantidad de posiciones en la tabla
	cmp       func(K, K) bool
}

type iteradorHash[K comparable, V any] struct {
	iterLista   TDAlista.IteradorLista[parClaveValor[K, V]]
	diccionario *hashAbierto[K, V]
	posDic      int
}

func crearTabla[K comparable, V any](capacidad int) []TDAlista.Lista[parClaveValor[K, V]] {
	tabla := make([]TDAlista.Lista[parClaveValor[K, V]], capacidad)
	for i := range tabla {
		tabla[i] = TDAlista.CrearListaEnlazada[parClaveValor[K, V]]()
	}
	return tabla
}

func CrearHash[K comparable, V any](cmp func(K, K) bool) Diccionario[K, V] {
	diccionario := new(hashAbierto[K, V])
	diccionario.capacidad = CAPACIDAD_INICIAL
	diccionario.tabla = crearTabla[K, V](diccionario.capacidad)
	diccionario.cantidad = 0
	diccionario.cmp = cmp
	return diccionario
}

func (d *hashAbierto[K, V]) redimensionar(nuevaCapacidad int) {
	nuevaTabla := crearTabla[K, V](nuevaCapacidad)

	for _, lista := range d.tabla {
		iter := lista.Iterador()
		for iter.HaySiguiente() {
			campo := iter.VerActual()
			indice := funcionHashing(campo.clave, nuevaCapacidad)
			nuevaTabla[indice].InsertarUltimo(campo)
			iter.Siguiente()
		}
	}

	d.tabla = nuevaTabla
	d.capacidad = nuevaCapacidad
}

func (d *hashAbierto[K, V]) buscarNodoClave(clave K) (TDAlista.IteradorLista[parClaveValor[K, V]], bool) {
	indice := funcionHashing(clave, d.capacidad)
	lista := d.tabla[indice]
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		campo := iter.VerActual()
		if d.cmp(campo.clave, clave) {
			return iter, true
		}
		iter.Siguiente()
	}
	return iter, false
}

func (d *hashAbierto[K, V]) Guardar(clave K, dato V) {
	factorCarga := float64(d.cantidad+1) / float64(d.capacidad)
	if factorCarga > FACTOR_CARGA_AUMENTO {
		d.redimensionar(d.capacidad * FACTOR_REDIMENSION)
	}

	iter, ok := d.buscarNodoClave(clave)
	indice := funcionHashing(clave, d.capacidad)

	if ok {
		campo := iter.VerActual()
		campo.dato = dato
		iter.Borrar()
		iter.Insertar(campo)
		return
	}

	d.tabla[indice].InsertarUltimo(parClaveValor[K, V]{clave, dato})
	d.cantidad++
}

func (d *hashAbierto[K, V]) Pertenece(clave K) bool {
	_, ok := d.buscarNodoClave(clave)
	return ok
}

func (d *hashAbierto[K, V]) Obtener(clave K) V {
	iter, ok := d.buscarNodoClave(clave)
	if !ok {
		panic("La clave no pertenece al diccionario")
	}
	campo := iter.VerActual()
	return campo.dato
}

func (d *hashAbierto[K, V]) Borrar(clave K) V {
	indice := funcionHashing(clave, d.capacidad)
	lista := d.tabla[indice]
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		campo := iter.VerActual()
		if d.cmp(campo.clave, clave) {
			iter.Borrar()
			d.cantidad--
			factorCarga := float64(d.cantidad) / float64(d.capacidad)
			if factorCarga < FACTOR_CARGA_REDUCCION && d.capacidad > CAPACIDAD_INICIAL {
				d.redimensionar(d.capacidad / FACTOR_REDIMENSION)
			}
			return campo.dato
		}
		iter.Siguiente()
	}
	panic("La clave no pertenece al diccionario")
}

func (d *hashAbierto[K, V]) Cantidad() int {
	return d.cantidad
}

func (d *hashAbierto[K, V]) Iterar(visitar func(K, V) bool) {
	for _, lista := range d.tabla {
		iter := lista.Iterador()
		for iter.HaySiguiente() {
			campo := iter.VerActual()
			if !visitar(campo.clave, campo.dato) {
				return
			}
			iter.Siguiente()
		}
	}
}

func (d *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	it := &iteradorHash[K, V]{
		diccionario: d,
		posDic:      0,
	}
	it.avanzarHastaElemento()
	return it
}

func (it *iteradorHash[K, V]) avanzarHastaElemento() {
	for it.posDic < len(it.diccionario.tabla) {
		it.iterLista = it.diccionario.tabla[it.posDic].Iterador()
		if it.iterLista.HaySiguiente() {
			return
		}
		it.posDic++
	}
	it.iterLista = nil
}

func (it *iteradorHash[K, V]) HaySiguiente() bool {
	return it.iterLista != nil && it.iterLista.HaySiguiente()
}

func (it *iteradorHash[K, V]) VerActual() (K, V) {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	campo := it.iterLista.VerActual()
	return campo.clave, campo.dato
}

func (it *iteradorHash[K, V]) Siguiente() {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	it.iterLista.Siguiente()
	if !it.iterLista.HaySiguiente() {
		it.posDic++
		it.avanzarHastaElemento()
	}
}

func funcionHashing[K comparable](clave K, tam int) int {
	h := fnv.New64a()
	h.Write([]byte(fmt.Sprintf("%v", clave)))
	return int(h.Sum64() % uint64(tam))
}
