package pila

const CANTIDAD_INICIAL_ELEMENTOS = 5
const FACTOR_REDIMENSION = 2
const FACTOR_MINIMA_CARGA = 4

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{datos: make([]T, CANTIDAD_INICIAL_ELEMENTOS)}
}

func (p *pilaDinamica[T]) redimensionar(nuevaCapacidad int) { // redimensiona *2 porque entre mas aumentan los elementos mas debe crecer el slice
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, p.datos[:p.cantidad])
	p.datos = nuevosDatos
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(valor T) {
	// chequear que llego a la capacidad maxima
	if p.cantidad == cap(p.datos) {
		nuevaCapacidad := p.cantidad * FACTOR_REDIMENSION
		p.redimensionar(nuevaCapacidad)
	}

	p.datos[p.cantidad] = valor
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	valor := p.datos[p.cantidad-1]
	p.cantidad--
	if ((p.cantidad * FACTOR_MINIMA_CARGA) <= cap(p.datos)) && CANTIDAD_INICIAL_ELEMENTOS < p.cantidad { // luego de desapilar, chequeo la capacidad de la pila y redimensionar para abajo de ser necesaario
		p.redimensionar(cap(p.datos) / FACTOR_REDIMENSION)
	}
	return valor
}
