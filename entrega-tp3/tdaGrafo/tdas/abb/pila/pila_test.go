package pila_test

import (
	"testing"
	TDAPila "tp3/tdaGrafo/tdas/pila"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "ver tope en la pila vacia debe lanzar panic")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "desapilar en pila vacia lanza panic")
}

func TestPilaApilarDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	pila.Apilar(5)
	require.False(t, pila.EstaVacia(), "la pila no debe estar vacía después de apilar")
	require.Equal(t, 5, pila.VerTope(), "el tope debe ser el último elemento apilado")

	pila.Apilar(10)
	require.Equal(t, 10, pila.VerTope(), "el tope debe ser el ultimo elemetno apilado")

	require.Equal(t, 10, pila.Desapilar(), "desapilar debe devolver el ultimo elemento apilado")
	require.Equal(t, 5, pila.VerTope(), "El nuevo tope debe ser el elemento anterior")
	require.Equal(t, 5, pila.Desapilar(), "Desapilar debe devolver el último elemento apilado")

	require.True(t, pila.EstaVacia(), "la pila debe quedar vacia despues de desapilar todos los elementos")

	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "ver tope en la pila vacía debe tirar panic")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "desapilar en pila vaciada debe lanzar panic")

}

func TestVolumenPila(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	volumen := 500000
	for i := 0; i < volumen; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope(), "el tope debe ser el ultimo elemento apilado")
	}

	for i := volumen - 1; i >= 0; i-- {
		require.Equal(t, i, pila.VerTope(), "el tope debe mantener el orden lifo")
		require.Equal(t, i, pila.Desapilar(), " desapilar debe seguir la regla LIFO, último en entrar, primero en salir")

	}

	require.True(t, pila.EstaVacia(), "la pila debe estar vacia despues de desapilar todos sus elementos")

	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "ver tope en pila vacia después debe lanzar panic")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "desapilar en pila vacia debe lanzar panic")

}

func TestTiposDatos(t *testing.T) {
	pilaString := TDAPila.CrearPilaDinamica[string]()
	pilaString.Apilar("mundo")
	pilaString.Apilar("comidaa")
	require.Equal(t, "comidaa", pilaString.VerTope(), "el tope de la pila de str debe ser correcto")
	require.Equal(t, "comidaa", pilaString.Desapilar(), "Desapilar de la pila de str debe ser correcto")
	require.Equal(t, "mundo", pilaString.Desapilar(), "Desapilar de la pila de str debe ser correcto")

	require.True(t, pilaString.EstaVacia(), "la pila debe quedar vacia tras desapilar todos los elementos ")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaString.VerTope() }, "Ver tope en pila de strings vacía debe lanzar pánico")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaString.Desapilar() }, "Desapilar en pila de strings vacía debe lanzar pánico")

	pilaFloat := TDAPila.CrearPilaDinamica[float64]()
	pilaFloat.Apilar(3.14)
	pilaFloat.Apilar(2.71)
	require.Equal(t, 2.71, pilaFloat.VerTope(), "Se espera que el tope de la pila sea el último elemento apilado")
	require.Equal(t, 2.71, pilaFloat.Desapilar(), "Se espera que desapilar devuelva elementos en ordeno")
	require.Equal(t, 3.14, pilaFloat.Desapilar(), "Desapilar de la pila de float64 debe ser correcto")

	pilaBool := TDAPila.CrearPilaDinamica[bool]()
	pilaBool.Apilar(true)
	pilaBool.Apilar(false)
	require.Equal(t, false, pilaBool.VerTope(), "El último valor apilado (false) debe estar en el tope")
	require.Equal(t, false, pilaBool.Desapilar(), "Al desapilar debe salir primero el valor false (último ingresado)")
	require.Equal(t, true, pilaBool.Desapilar(), "Al desapilar nuevamente debe obtenerse true (primer valor)")

}

func TestPilaVaciadaComoNueva(t *testing.T) {
	pilaNueva := TDAPila.CrearPilaDinamica[int]()

	pilaUsada := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 5; i++ {
		pilaUsada.Apilar(i)
	}

	for i := 0; i < 5; i++ {
		pilaUsada.Desapilar()
	}

	require.True(t, pilaNueva.EstaVacia(), "La pila nueva debe estar vacía")
	require.True(t, pilaUsada.EstaVacia(), "La pila vaciada debe estar vacía")

	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaNueva.VerTope() }, "Ver tope en pila nueva debe lanzar pánico")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaUsada.VerTope() }, "Ver tope en pila vaciada debe lanzar pánico")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaNueva.Desapilar() }, "Desapilar en pila nueva debe lanzar pánico")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaUsada.Desapilar() }, "Desapilar en pila vaciada debe lanzar pánico")

	pilaNueva.Apilar(10)
	pilaUsada.Apilar(10)
	require.Equal(t, pilaNueva.VerTope(), pilaUsada.VerTope(), "Ambas pilas deben tener el mismo tope")

}
