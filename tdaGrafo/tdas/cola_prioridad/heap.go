package cola_prioridad

const CAPACIDAD_INICIAL = 10
const REDIM_AGRANDAR = 2
const REDIMEN_REDUCIR = 4
const FACTOR_REDUCIR = 2

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func swap[T any](a, b *T) {
	*a, *b = *b, *a
}

func nuevoHeap[T any](capacidad int, cmp func(T, T) int) *colaConPrioridad[T] {
	return &colaConPrioridad[T]{
		datos: make([]T, capacidad),
		cant:  0,
		cmp:   cmp,
	}
}

func CrearHeap[T comparable](cmp func(T, T) int) ColaPrioridad[T] {
	return nuevoHeap(CAPACIDAD_INICIAL, cmp)
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	capacidad := len(arreglo)
	if capacidad < CAPACIDAD_INICIAL {
		capacidad = CAPACIDAD_INICIAL
	}
	heap := &colaConPrioridad[T]{
		datos: make([]T, capacidad),
		cant:  len(arreglo),
		cmp:   funcion_cmp,
	}
	copy(heap.datos, arreglo)
	heapify(heap.datos, heap.cant, heap.cmp)
	return heap
}

func (heap *colaConPrioridad[T]) EstaVacia() bool {
	return heap.cant == 0
}

func (heap *colaConPrioridad[T]) Encolar(elem T) {
	esNecesarioRedimensionar, capNueva := heap.hayQueAgrandar()
	if esNecesarioRedimensionar {
		heap.redimensionar(capNueva)
	}
	heap.datos[heap.cant] = elem
	upheap(heap.datos, heap.cant, heap.cmp)
	heap.cant++
}

func (heap *colaConPrioridad[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.datos[0]
}

func (heap *colaConPrioridad[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	datoDesencolado := heap.datos[0]
	heap.cant--
	swap(&heap.datos[0], &heap.datos[heap.cant])
	downheap(heap.datos, 0, heap.cant, heap.cmp)

	if esNecesarioAchicar, capNueva := heap.hayQueAchicar(); esNecesarioAchicar {
		heap.redimensionar(capNueva)
	}
	return datoDesencolado
}

func (heap *colaConPrioridad[T]) Cantidad() int {
	return heap.cant
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	n := len(elementos)
	heapify(elementos, n, funcion_cmp)
	for i := n - 1; i > 0; i-- {
		swap(&elementos[0], &elementos[i])
		downheap(elementos, 0, i, funcion_cmp)
	}
}

func heapify[T any](datos []T, cant int, cmp func(T, T) int) {
	for i := (cant / 2) - 1; i >= 0; i-- {
		downheap(datos, i, cant, cmp)
	}
}

func upheap[T any](datos []T, pos int, cmp func(T, T) int) {
	if pos == 0 {
		return
	}
	padre := (pos - 1) / 2
	if cmp(datos[pos], datos[padre]) > 0 {
		swap(&datos[pos], &datos[padre])
		upheap(datos, padre, cmp)
	}
}

func downheap[T any](datos []T, pos, cant int, cmp func(T, T) int) {
	hijoIzq := 2*pos + 1
	hijoDer := 2*pos + 2
	mayor := pos

	if hijoIzq < cant && cmp(datos[hijoIzq], datos[mayor]) > 0 {
		mayor = hijoIzq
	}
	if hijoDer < cant && cmp(datos[hijoDer], datos[mayor]) > 0 {
		mayor = hijoDer
	}
	if mayor != pos {
		swap(&datos[pos], &datos[mayor])
		downheap(datos, mayor, cant, cmp)
	}
}

func (heap *colaConPrioridad[T]) hayQueAgrandar() (bool, int) {
	capacidad := cap(heap.datos)
	if heap.cant == capacidad {
		return true, capacidad * REDIM_AGRANDAR
	}
	return false, 0
}

func (heap *colaConPrioridad[T]) hayQueAchicar() (bool, int) {
	capacidad := cap(heap.datos)
	if capacidad > CAPACIDAD_INICIAL && heap.cant*REDIMEN_REDUCIR <= capacidad {
		return true, capacidad / FACTOR_REDUCIR
	}
	return false, 0
}

func (heap *colaConPrioridad[T]) redimensionar(capNueva int) {
	array := make([]T, capNueva)
	copy(array, heap.datos[:heap.cant])
	heap.datos = array
}
