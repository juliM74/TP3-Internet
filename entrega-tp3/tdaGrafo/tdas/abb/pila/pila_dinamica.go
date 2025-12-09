package pila

const (
	CAPACIDAD_INICIAL = 5
	EXPANSION         = 2
	ESCALA_REDUCCION  = 4
	OFFSET            = 1
	FACTOR_REDUCCION  = 2
)

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{datos: make([]T, CAPACIDAD_INICIAL), cantidad: 0}
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

func (p *pilaDinamica[T]) redimensionar(NuevaCapacidad int) {
	nuevosDatos := make([]T, NuevaCapacidad)
	copy(nuevosDatos, p.datos[:p.cantidad])
	p.datos = nuevosDatos
}

func (p *pilaDinamica[T]) Apilar(valor T) {
	if len(p.datos) == p.cantidad {
		capacidad := p.cantidad*EXPANSION + OFFSET
		p.redimensionar(capacidad)
	}

	p.datos[p.cantidad] = valor
	p.cantidad++

}

func (p *pilaDinamica[T]) debeReducir() bool {
	return len(p.datos) > 0 && p.cantidad <= len(p.datos)/ESCALA_REDUCCION
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}

	p.cantidad--
	valor := p.datos[p.cantidad]

	if p.debeReducir() {
		capacidad := len(p.datos) / FACTOR_REDUCCION
		p.redimensionar(capacidad)
	}

	return valor
}
