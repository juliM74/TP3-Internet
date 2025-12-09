package cola_test

import (
	"testing"
	TDACola "tp3/tdaGrafo/tdas/cola"

	"github.com/stretchr/testify/require"
)

const ITERACIONES_VOLUMEN = 10000

type parNumeros struct {
	numero1 int
	numero2 int
}

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "no se puede ver el primer elemento de una cola vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "no se puede desencolar una cola vacia")
}

func TestPruebaDeVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < ITERACIONES_VOLUMEN; i++ {
		cola.Encolar(i)
	}
	// Desencolar y verificar orden
	for i := 0; i < ITERACIONES_VOLUMEN; i++ {
		require.Equal(t, i, cola.Desencolar(), "El elemento desencolado deberia ser el primer agregado FIFO")
	}
	require.True(t, cola.EstaVacia(), "La cola deberia estar vacia luego de desencolar por completo")
	// Luego de desencolar todo, verifico que se comporta como cola vacia
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "no se puede ver el primer elemento de una cola vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "no se puede desencolar una cola vacia")
}

func TestEncolarYDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	elementos := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(elementos); i++ {
		elemento := elementos[i]
		cola.Encolar(elemento)
	}
	// desencolar hasta quedar vacia
	for !cola.EstaVacia() {
		primeroActual := cola.VerPrimero()
		desencolado := cola.Desencolar()
		require.Equal(t, desencolado, primeroActual, "el primero deberia ser el primer desencolado FIFO")
	}
	require.True(t, cola.EstaVacia(), "La cola deberia estar vacia despues de desencolar todos los elementos")
}

func TestColaVaciaDespuesDeDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	elementos := []int{10, 20, 30}
	for _, elem := range elementos {
		cola.Encolar(elem)
	}
	for !cola.EstaVacia() {
		cola.Desencolar()
	}
	require.True(t, cola.EstaVacia(), "La cola debería estar vacía después de desencolar todos los elementos")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "El primero de una cola vacía no se puede ver")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "No se puede desencolar una cola vacía")
}

// encolar y desencolar con distintos tipos
// --------------------------------- (ejemplo con strings) ---------------------------------
func TestColaConStrings(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("Hola")
	cola.Encolar("Mundo")
	require.Equal(t, "Hola", cola.Desencolar(), "Debe desencolarse 'Hola' primero")
	require.Equal(t, "Mundo", cola.Desencolar(), "Debe desencolarse 'Mundo' después")
	require.True(t, cola.EstaVacia(), "La cola debería estar vacía después de desencolar todo")
}

// --------------------------------- (ejemplo con floats) ---------------------------------
func TestColaConFloats(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[float64]()
	cola.Encolar(3.14)
	cola.Encolar(2.71)
	require.Equal(t, 3.14, cola.Desencolar(), "Debe desencolarse 3.14 primero")
	require.Equal(t, 2.71, cola.Desencolar(), "Debe desencolarse 2.71 después")
	require.True(t, cola.EstaVacia(), "La cola debería estar vacía después de desencolar todo")
}

// --------------------------------- (ejemplo con structs) ---------------------------------
func TestColaDeStructs(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[parNumeros]()
	elementos := []parNumeros{
		{numero1: 1, numero2: 2},
		{numero1: 5, numero2: 80},
		{numero1: 7, numero2: 10},
	}
	for _, elem := range elementos {
		cola.Encolar(elem)
	}
	require.False(t, cola.EstaVacia(), "La cola no debería estar vacía")

	// Prueba de ver primero
	primero := cola.VerPrimero()
	require.Equal(t, 1, primero.numero1)
	require.Equal(t, 2, primero.numero2)

	// Desencolar y verificar FIFO
	for _, esperado := range elementos {
		desencolado := cola.Desencolar()
		require.Equal(t, esperado, desencolado)
	}

	require.True(t, cola.EstaVacia(), "La cola debería estar vacía después de desencolar todo")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "No se puede ver el primero de una cola vacía")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "No se puede desencolar una cola vacía")
}
