package pila_test

import (
	"testing"
	TDAPila "tp3/tdaGrafo/tdas/pila"

	"github.com/stretchr/testify/require"
)

const ITERACIONES_VOLUMEN = 10000

type parNumeros struct {
	numero1 int
	numero2 int
}
type structPrueba []parNumeros

// pila recién creada está vacía
func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "no se puede ver tope de pila vacia")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "no se puede desapilar cuando la pila esta vacia")
}

func TestPruebaDeVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	for i := 0; i < ITERACIONES_VOLUMEN; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope(), "El tope debería ser el último elemento apilado")
	}
	// Desapilar y verificar orden
	for i := ITERACIONES_VOLUMEN - 1; i >= 0; i-- {
		require.Equal(t, i, pila.Desapilar(), "El elemento desapilado no coincide con el esperado")
	}
	require.True(t, pila.EstaVacia(), "La pila debería estar vacía después de desapilar todo")
	// Luego de vaciar la pila, verifico que esta se comporte como tal
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "no se puede ver tope de pila vacia")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "no se puede desapilar cuando la pila esta vacia")

	nuevoElem := 9000000
	pila.Apilar(nuevoElem)
	require.Equal(t, nuevoElem, pila.VerTope())
	require.False(t, pila.EstaVacia(), "La pila ya no deberia estar vacia")

}

// apilar y desapilar varios elementos
func TestApilarYDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()
	elementos := []float64{1.2, 2.33, 3.784, 4.33, 5.123456789, 9.9999, 6.666}

	for i := 0; i < len(elementos); i++ {
		elemento := elementos[i]
		pila.Apilar(elemento)
		require.Equal(t, elemento, pila.VerTope(), "El tope debería ser el último elemento apilado")
	}
	// desapilar en orden inverso
	for i := len(elementos) - 1; i >= 0; i-- {
		require.Equal(t, elementos[i], pila.Desapilar(), "El elemento desapilado no coincide con el esperado")
	}
	require.True(t, pila.EstaVacia(), "La pila debería estar vacía después de desapilar todo")

	nuevoElem := 7.32
	pila.Apilar(nuevoElem)
	require.Equal(t, nuevoElem, pila.VerTope())
	require.False(t, pila.EstaVacia(), "La pila ya no deberia estar vacia")
}

// Probar que una pila con elementos, luego de ser desapilada funciona como una pila recien creada vacia
func TestPilaVaciaDespuesDeDesapilar(t *testing.T) {
	miNombre := "Julian"
	pila := TDAPila.CrearPilaDinamica[string]()
	elementos := []string{"Hola", "mi", "nombre", "es", miNombre}
	for _, elem := range elementos {
		pila.Apilar(elem)
	}
	require.Equal(t, miNombre, pila.VerTope())
	for !pila.EstaVacia() {
		pila.Desapilar()
	}
	require.True(t, pila.EstaVacia(), "La pila debería estar vacía después de desapilar todos los elementos")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "El tope de una pila vacía no se puede ver")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "No se puede desapilar una pila vacía")

	comoEsMiNombre := "Mi nombre sigue siendo Julian"
	pila.Apilar(comoEsMiNombre)
	require.Equal(t, comoEsMiNombre, pila.VerTope())
	require.False(t, pila.EstaVacia(), "La pila ya no deberia estar vacia")
}

// apilar y desapilar con distintos tipos

// --------------------------------- (ejemplo con structs) ---------------------------------
func TestPilaDeStructs(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[parNumeros]()
	elementos := structPrueba{
		{numero1: 1, numero2: 2},
		{numero1: 5, numero2: 80},
		{numero1: 7, numero2: 10},
	}
	// Cargo elementos a pila de structs
	for _, elem := range elementos {
		pila.Apilar(elem)
	}
	require.False(t, pila.EstaVacia(), "la pila no deberia estar vacia")
	// Prueba de ver tope
	tope := pila.VerTope()
	require.Equal(t, 7, tope.numero1)
	require.Equal(t, 10, tope.numero2)
	// Pruebo desapilar LIFO
	desapilado := pila.Desapilar()
	require.Equal(t, 7, desapilado.numero1)
	require.Equal(t, 10, desapilado.numero2)
	// Vacio la pila y pruebo si se comporta como una pila vacia
	for !pila.EstaVacia() {
		pila.Desapilar()
	}
	require.True(t, pila.EstaVacia(), "la pila deberia estar vacia despues de desapilar todo")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "no se puede ver tope de pila vacia")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "no se puede desapilar cuando la pila esta vacia")
}

// --------------------------------- (ejemplo con pilas) ---------------------------------
func TestPilaDePilas(t *testing.T) {
	pilaExterna := TDAPila.CrearPilaDinamica[TDAPila.Pila[int]]()

	pilaInterna1 := TDAPila.CrearPilaDinamica[int]()
	pilaInterna1.Apilar(1)
	pilaInterna1.Apilar(2)

	pilaInterna2 := TDAPila.CrearPilaDinamica[int]()
	pilaInterna2.Apilar(3)
	pilaInterna2.Apilar(4)

	pilaExterna.Apilar(pilaInterna1)
	pilaExterna.Apilar(pilaInterna2)

	require.Equal(t, 4, pilaExterna.VerTope().VerTope(), "El tope de la pila externa debería contener el tope 4")
	require.Equal(t, 4, pilaExterna.Desapilar().Desapilar(), "Desapilando la pila interna con tope 4")
	require.Equal(t, 2, pilaExterna.VerTope().VerTope(), "El nuevo tope debería ser 2")
}
