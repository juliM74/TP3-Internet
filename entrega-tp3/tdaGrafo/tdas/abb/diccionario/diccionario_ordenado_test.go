package diccionario_test

import (
	"strings"
	"testing"
	TDADiccionario "tp3/tdaGrafo/tdas/abb/diccionario"

	"github.com/stretchr/testify/require"
)

func TestDiccionarioVacio1(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("X"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("X") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("X") })
}

func TestClavePorDefecto(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNumerico := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	require.False(t, dicNumerico.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNumerico.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNumerico.Borrar(0) })
}

func TestDiccionarioUnElemento(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar("X", 42)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("X"))
	require.False(t, dic.Pertenece("Y"))
	require.EqualValues(t, 42, dic.Obtener("X"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("Y") })
}

func TestGuardarElementos(t *testing.T) {
	claves := []string{"Manzana", "Banana", "Cereza"}
	valores := []string{"Rojo", "Amarillo", "Rojo"}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	for i, clave := range claves {
		require.False(t, dic.Pertenece(clave))
		dic.Guardar(clave, valores[i])
		require.True(t, dic.Pertenece(clave))
		require.EqualValues(t, valores[i], dic.Obtener(clave))
		require.EqualValues(t, i+1, dic.Cantidad())
	}
}

func TestReemplazoDeValores(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar("Perro", "come")
	dic.Guardar("Gato", "Maulla")
	require.EqualValues(t, "come", dic.Obtener("Perro"))
	require.EqualValues(t, "Maulla", dic.Obtener("Gato"))

	dic.Guardar("Perro", "llora")
	dic.Guardar("Gato", "Ronronea")
	require.EqualValues(t, "llora", dic.Obtener("Perro"))
	require.EqualValues(t, "Ronronea", dic.Obtener("Gato"))
}

func TestBorrarElementos(t *testing.T) {
	claves := []string{"Uno", "Dos", "Tres"}
	valores := []string{"Primero", "Segundo", "Tercero"}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	for i, clave := range claves {
		dic.Guardar(clave, valores[i])
	}

	for i, clave := range claves {
		require.True(t, dic.Pertenece(clave))
		require.EqualValues(t, valores[i], dic.Borrar(clave))
		require.False(t, dic.Pertenece(clave))
	}
	require.EqualValues(t, 0, dic.Cantidad())
}

func TestIteradorInterno(t *testing.T) {
	claves := []string{"A", "B", "C"}
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	for i, clave := range claves {
		dic.Guardar(clave, i)
	}

	recorridas := make([]string, 0)
	dic.Iterar(func(clave string, valor int) bool {
		recorridas = append(recorridas, clave)
		return true
	})

	require.Equal(t, claves, recorridas)
}

func TestIteradorConCorte(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar("a", "uno")
	dic.Guardar("b", "dos")
	dic.Guardar("c", "tres")
	dic.Guardar("d", "cuatro")
	dic.Guardar("e", "cinco")

	visitadas := []string{}
	continuar := true
	dic.Iterar(func(clave string, _ string) bool {
		if clave == "c" {
			continuar = false
			return false
		}
		visitadas = append(visitadas, clave)
		return true
	})

	require.True(t, !continuar)
	require.Equal(t, []string{"a", "b"}, visitadas)
}

func TestIteradorTrasBorradoPrimerElemento(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar("a", "uno")
	dic.Guardar("b", "dos")
	dic.Guardar("c", "tres")

	dic.Borrar("a")

	visitados := []string{}
	dic.Iterar(func(clave string, _ string) bool {
		visitados = append(visitados, clave)
		return true
	})

	require.Equal(t, []string{"b", "c"}, visitados)
}

func TestIteradorTrasBorradoUltimoElemento(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar("a", "uno")
	dic.Guardar("b", "dos")
	dic.Guardar("c", "tres")

	dic.Borrar("c")

	visitados := []string{}
	dic.Iterar(func(clave string, _ string) bool {
		visitados = append(visitados, clave)
		return true
	})

	require.Equal(t, []string{"a", "b"}, visitados)
}

func TestVolumenGuardarYBorrar(t *testing.T) {
	const n = 10000
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	for i := 0; i < n; i++ {
		dic.Guardar(i, i)
	}
	require.Equal(t, n, dic.Cantidad())
	for i := 0; i < n; i++ {
		require.True(t, dic.Pertenece(i))
		require.Equal(t, i, dic.Borrar(i))
	}
	require.Equal(t, 0, dic.Cantidad())
}
func TestIteradorTrasBorradoTodos(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar("a", "uno")
	dic.Guardar("b", "dos")

	dic.Borrar("a")
	dic.Borrar("b")

	visitados := []string{}
	dic.Iterar(func(clave string, _ string) bool {
		visitados = append(visitados, clave)
		return true
	})

	require.Empty(t, visitados)
}

func TestIteradorTrasBorradoYReinsercion(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar("x", "uno")
	dic.Guardar("y", "dos")
	dic.Guardar("z", "tres")

	dic.Borrar("y")
	dic.Guardar("y", "nuevo")

	visitados := []string{}
	dic.Iterar(func(clave string, _ string) bool {
		visitados = append(visitados, clave)
		return true
	})

	require.Equal(t, []string{"x", "y", "z"}, visitados)
	require.Equal(t, "nuevo", dic.Obtener("y"))
}

func TestIteradorRangoParcial(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	dic.Guardar(5, "cinco")
	dic.Guardar(10, "diez")
	dic.Guardar(15, "quince")
	dic.Guardar(20, "veinte")

	desde, hasta := 10, 15
	iter := dic.IteradorRango(&desde, &hasta)

	visitados := []int{}
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		visitados = append(visitados, clave)
		iter.Siguiente()
	}

	require.Equal(t, []int{10, 15}, visitados)
}

func TestRecorridoInversoConRecursividad(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](func(k1, k2 int) int { return k1 - k2 })
	abb.Guardar(10, "diez")
	abb.Guardar(5, "cinco")
	abb.Guardar(15, "quince")
	abb.Guardar(3, "tres")
	abb.Guardar(7, "siete")
	abb.Guardar(12, "doce")
	abb.Guardar(18, "dieciocho")

	resultado := []int{}
	abb.Iterar(func(clave int, _ string) bool {

		resultado = append([]int{clave}, resultado...)
		return true
	})

	require.Equal(t, []int{18, 15, 12, 10, 7, 5, 3}, resultado)
}

func TestBorrarCasosBorde(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })

	dic.Guardar(2, "dos")
	dic.Guardar(1, "uno")
	dic.Guardar(3, "tres")
	require.Equal(t, "dos", dic.Borrar(2))
	require.False(t, dic.Pertenece(2))
	dic.Guardar(4, "cuatro")
	require.Equal(t, "tres", dic.Borrar(3))
	require.False(t, dic.Pertenece(3))

	require.Equal(t, "cuatro", dic.Borrar(4))
	require.False(t, dic.Pertenece(4))
}

func TestDeteccionMaximosYMinimos(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	abb.Guardar(10, "diez")
	abb.Guardar(5, "cinco")
	abb.Guardar(15, "quince")
	abb.Guardar(3, "tres")
	abb.Guardar(7, "siete")

	min := abb.IteradorRango(nil, nil)
	require.True(t, min.HaySiguiente())
	claveMin, _ := min.VerActual()
	require.Equal(t, 3, claveMin)

	max := abb.IteradorRango(nil, nil)
	var claveMax int
	for max.HaySiguiente() {
		claveMax, _ = max.VerActual()
		max.Siguiente()
	}
	require.Equal(t, 15, claveMax)
}

func TestRecorridos(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	abb.Guardar(10, "diez")
	abb.Guardar(5, "cinco")
	abb.Guardar(15, "quince")
	abb.Guardar(3, "tres")
	abb.Guardar(7, "siete")

	enOrden := []int{}
	abb.Iterar(func(clave int, _ string) bool {
		enOrden = append(enOrden, clave)
		return true
	})
	require.Equal(t, []int{3, 5, 7, 10, 15}, enOrden)
}

func TestIteradorTrasInsercionesYBorrados(t *testing.T) {
	t.Log("Valida que el iterador funcione correctamente tras múltiples inserciones y borrados")
	dic := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })

	dic.Guardar(10, "diez")
	dic.Guardar(5, "cinco")
	dic.Guardar(15, "quince")
	dic.Guardar(3, "tres")
	dic.Guardar(7, "siete")

	dic.Borrar(5)
	dic.Borrar(15)

	dic.Guardar(6, "seis")
	dic.Guardar(8, "ocho")

	esperado := []int{3, 6, 7, 8, 10}
	recorrido := []int{}
	dic.Iterar(func(clave int, _ string) bool {
		recorrido = append(recorrido, clave)
		return true
	})

	require.Equal(t, esperado, recorrido)
}

func TestIteradorRangoTrasModificaciones(t *testing.T) {
	t.Log("Valida que el iterador de rango funcione correctamente tras inserciones y borrados")
	dic := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })

	dic.Guardar(10, "diez")
	dic.Guardar(5, "cinco")
	dic.Guardar(15, "quince")
	dic.Guardar(3, "tres")
	dic.Guardar(7, "siete")

	dic.Borrar(5)
	dic.Borrar(15)

	dic.Guardar(6, "seis")
	dic.Guardar(8, "ocho")

	desde, hasta := 6, 10
	iter := dic.IteradorRango(&desde, &hasta)

	esperadoClaves := []int{6, 7, 8, 10}
	esperadoDatos := []string{"seis", "siete", "ocho", "diez"}
	recorridoClaves := []int{}
	recorridoDatos := []string{}
	for iter.HaySiguiente() {
		clave, dato := iter.VerActual()
		recorridoClaves = append(recorridoClaves, clave)
		recorridoDatos = append(recorridoDatos, dato)
		iter.Siguiente()
	}

	require.Equal(t, esperadoClaves, recorridoClaves)
	require.Equal(t, esperadoDatos, recorridoDatos)
}
func TestIteradorEnArbolVacioTrasOperaciones(t *testing.T) {
	t.Log("Valida que el iterador funcione correctamente en un árbol vacío tras operaciones de inserción y borrado")
	dic := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })

	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })

	dic.Guardar(10, "diez")
	dic.Borrar(10)

	iter = dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorRangoSoloDesde(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	for _, k := range []int{2, 4, 6, 8, 10} {
		dic.Guardar(k, "")
	}
	desde := 6
	iter := dic.IteradorRango(&desde, nil)
	visitados := []int{}
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		visitados = append(visitados, clave)
		iter.Siguiente()
	}
	require.Equal(t, []int{6, 8, 10}, visitados)
}

func TestIteradorRangoSoloHasta(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	for _, k := range []int{2, 4, 6, 8, 10} {
		dic.Guardar(k, "")
	}
	hasta := 6
	iter := dic.IteradorRango(nil, &hasta)
	visitados := []int{}
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		visitados = append(visitados, clave)
		iter.Siguiente()
	}
	require.Equal(t, []int{2, 4, 6}, visitados)
}

func TestIteradorRangoAmbosNil(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	for _, k := range []int{2, 4, 6} {
		dic.Guardar(k, "")
	}
	iter := dic.IteradorRango(nil, nil)
	visitados := []int{}
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		visitados = append(visitados, clave)
		iter.Siguiente()
	}
	require.Equal(t, []int{2, 4, 6}, visitados)
}

func TestIteradorRangoCortePorFalse(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	for _, k := range []int{1, 2, 3, 4, 5} {
		dic.Guardar(k, "")
	}
	desde, hasta := 2, 5
	iter := dic.IteradorRango(&desde, &hasta)
	visitados := []int{}
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		visitados = append(visitados, clave)
		if clave == 3 {
			break
		}
		iter.Siguiente()
	}
	require.Equal(t, []int{2, 3}, visitados)
}

func TestIteradorInternoOrdenCompleto(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	claves := []int{5, 2, 8, 1, 3}
	for _, k := range claves {
		dic.Guardar(k, "")
	}
	recorrido := []int{}
	dic.Iterar(func(clave int, _ string) bool {
		recorrido = append(recorrido, clave)
		return true
	})
	require.Equal(t, []int{1, 2, 3, 5, 8}, recorrido)
}

func TestVolumenIteradorInterno(t *testing.T) {
	const n = 10000
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	for i := 0; i < n; i++ {
		dic.Guardar(i, i)
	}
	suma := 0
	dic.Iterar(func(clave int, valor int) bool {
		suma += valor
		return true
	})
	require.Equal(t, (n-1)*n/2, suma)
}

func TestIteradorInternoConCorte(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	for _, k := range []int{1, 2, 3, 4, 5} {
		dic.Guardar(k, "")
	}
	recorrido := []int{}
	dic.Iterar(func(clave int, _ string) bool {
		recorrido = append(recorrido, clave)
		return clave < 3
	})
	require.Equal(t, []int{1, 2, 3}, recorrido)
}

func TestIteradorExternoCompleto(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	claves := []int{4, 2, 6, 1, 3, 5, 7}
	for _, k := range claves {
		dic.Guardar(k, "")
	}
	iter := dic.Iterador()
	recorrido := []int{}
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		recorrido = append(recorrido, clave)
		iter.Siguiente()
	}
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, recorrido)
}
func TestVolumenIteradorExterno(t *testing.T) {
	const n = 10000
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	for i := 0; i < n; i++ {
		dic.Guardar(i, i)
	}
	iter := dic.Iterador()
	suma := 0
	for iter.HaySiguiente() {
		_, valor := iter.VerActual()
		suma += valor
		iter.Siguiente()
	}
	require.Equal(t, (n-1)*n/2, suma)
}
func TestIteradorExternoConRango(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	for _, k := range []int{10, 20, 30, 40, 50} {
		dic.Guardar(k, "")
	}
	desde, hasta := 20, 40
	iter := dic.IteradorRango(&desde, &hasta)
	recorrido := []int{}
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		recorrido = append(recorrido, clave)
		iter.Siguiente()
	}
	require.Equal(t, []int{20, 30, 40}, recorrido)
}

func TestIteradorExternoRangoVacio(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	for _, k := range []int{1, 2, 3} {
		dic.Guardar(k, "")
	}
	desde, hasta := 4, 5
	iter := dic.IteradorRango(&desde, &hasta)
	require.False(t, iter.HaySiguiente())
}
