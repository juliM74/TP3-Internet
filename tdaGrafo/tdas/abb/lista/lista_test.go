package lista_test

import (
	"testing"
	TDAlista "tp3/tdaGrafo/tdas/lista"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	t.Log("Test de lista vacia ")
	lista := TDAlista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, 0, lista.Largo())
}

func TestListaInsertarSinIter(t *testing.T) {
	t.Log("Test de lista al insertar ")
	lista := TDAlista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 1, lista.VerUltimo())
	require.EqualValues(t, 1, lista.Largo())

	lista.InsertarPrimero(2)
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 1, lista.VerUltimo())
	require.EqualValues(t, 2, lista.Largo())

	lista.InsertarPrimero(0)
	lista.InsertarUltimo(3)
	require.EqualValues(t, 0, lista.VerPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())
	require.EqualValues(t, 4, lista.Largo())

	require.EqualValues(t, 0, lista.BorrarPrimero())
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 3, lista.Largo())

	// Borrar los elementos restantes
	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.EqualValues(t, 1, lista.BorrarPrimero())
	require.EqualValues(t, 3, lista.BorrarPrimero())

	// Ahora sí debería estar vacía
	require.True(t, lista.EstaVacia(), "La lista debe quedar vacía después de borrar todos los elementos")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestIteradorBorradoPrincipio(t *testing.T) {
	t.Log("Prueba de borrado al principio con iterador")
	lista := TDAlista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	iter := lista.Iterador()
	require.EqualValues(t, 1, iter.Borrar())

	// Verifica que el primero sea el nuevo correcto
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 2, lista.Largo())

	// Verifica avance del iterador y elementos restantes
	require.True(t, iter.HaySiguiente())
	require.EqualValues(t, 2, iter.VerActual())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.EqualValues(t, 3, iter.VerActual())

	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorBorradoFinal(t *testing.T) {
	t.Log("Prueba de borrado al final con iterador")
	lista := TDAlista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	iter := lista.Iterador()
	iter.Siguiente()
	iter.Siguiente() // 3

	require.EqualValues(t, 3, iter.VerActual())
	require.EqualValues(t, 3, iter.Borrar())

	require.EqualValues(t, 2, lista.VerUltimo())
	require.EqualValues(t, 2, lista.Largo())
	require.False(t, iter.HaySiguiente()) // Chequea si no hay mas elementos

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })

	// Verificamos que siguen los otros elementos
	iter2 := lista.Iterador()
	require.True(t, iter2.HaySiguiente())
	require.EqualValues(t, 1, iter2.VerActual())
	iter2.Siguiente()
	require.EqualValues(t, 2, iter2.VerActual())
	require.True(t, iter2.HaySiguiente())
	iter2.Siguiente()
	require.False(t, iter2.HaySiguiente())
}

func TestIteradorBorradoMedio(t *testing.T) {
	t.Log("Prueba de borrado en medio con iterador")
	lista := TDAlista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	iter.Siguiente()
	require.EqualValues(t, 2, iter.Borrar())

	iter = lista.Iterador()
	require.EqualValues(t, 1, iter.VerActual())
	iter.Siguiente()
	require.EqualValues(t, 3, iter.VerActual())
	require.EqualValues(t, 2, lista.Largo())

	require.True(t, iter.HaySiguiente())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())

}

func TestIteradorInterno(t *testing.T) {
	t.Log("Prueba de iterador interno")
	lista := TDAlista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	suma := 0
	lista.Iterar(func(elem int) bool {
		suma += elem
		return true
	})

	require.EqualValues(t, 6, suma)
}

func TestIteradorInternoConCorte(t *testing.T) {
	t.Log("Prueba de iterador interno con corte")
	lista := TDAlista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)

	suma := 0
	lista.Iterar(func(elem int) bool {
		suma += elem
		return elem < 3
	})
	require.EqualValues(t, 6, suma)
}

func TestIteradorInternoSinCorteNiElementos(t *testing.T) {
	t.Log("Prueba de iterador interno sin corte y con lista vacia")
	lista := TDAlista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)

	// Caso sin corte
	sumaTotal := 0
	lista.Iterar(func(elem int) bool {
		sumaTotal += elem
		return true
	})
	require.EqualValues(t, 10, sumaTotal)

	// Caso sin elementos
	listaVacia := TDAlista.CrearListaEnlazada[int]()
	listaVacia.Iterar(func(elem int) bool {
		t.Errorf("No se esperaba iteración")
		return true
	})
}

func TestIteradorInternoConCorteVolumen(t *testing.T) {
	t.Log("Iterador interno con corte en lista de gran volumen")
	lista := TDAlista.CrearListaEnlazada[int]()
	for i := 0; i < 10000; i++ {
		lista.InsertarUltimo(i)
	}

	suma := 0
	lista.Iterar(func(elem int) bool {
		suma += elem
		return elem < 4999
	})

	require.Greater(t, suma, 0)
	require.Less(t, suma, 4999*5000/2+100)
}

func TestIteradorInternoSinElementosSinCorte(t *testing.T) {
	t.Log("Iterador interno sin elementos sin corte")
	lista := TDAlista.CrearListaEnlazada[int]()
	llamado := false

	lista.Iterar(func(elem int) bool {
		llamado = true
		return true
	})

	require.False(t, llamado, "No debería llamarse la función sobre una lista vacía")
}

func TestIteradorInternoSinElementosConCorte(t *testing.T) {
	t.Log("Iterador interno sin elementos con corte")
	lista := TDAlista.CrearListaEnlazada[int]()
	llamado := false

	lista.Iterar(func(elem int) bool {
		llamado = true
		return false
	})

	require.False(t, llamado, "No debería llamarse la función sobre una lista vacía, aunque haya corte")
}

func TestDiferentesTipos(t *testing.T) {
	t.Log("Prueba con diferentes tipos de datos")

	listaStr := TDAlista.CrearListaEnlazada[string]()
	listaStr.InsertarUltimo("comidaaa")
	listaStr.InsertarUltimo("boca")
	require.EqualValues(t, "comidaaa", listaStr.VerPrimero())
	require.EqualValues(t, "boca", listaStr.VerUltimo())

	listaFloat := TDAlista.CrearListaEnlazada[float64]()
	listaFloat.InsertarUltimo(3.14)
	listaFloat.InsertarUltimo(2.78)
	require.EqualValues(t, 3.14, listaFloat.VerPrimero())
	require.EqualValues(t, 2.78, listaFloat.VerUltimo())
}
func TestIteradorInsercionPrin(t *testing.T) {
	t.Log("Prueba: insertar un elemento en la posición en la que se crea el iterador")
	lista := TDAlista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	iter.Insertar(1)

	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 3, lista.Largo())

	iter = lista.Iterador()
	require.EqualValues(t, 1, iter.VerActual())
	iter.Siguiente()
	require.EqualValues(t, 2, iter.VerActual())
	iter.Siguiente()
	require.EqualValues(t, 3, iter.VerActual())
	iter.Siguiente()

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
}

func TestIteradorInsercionFinal(t *testing.T) {
	t.Log("Prueba: insertar un elemento cuando el iterador está al final")
	lista := TDAlista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)

	iter := lista.Iterador()
	iter.Siguiente()
	iter.Siguiente()

	iter.Insertar(3)

	require.EqualValues(t, 3, lista.VerUltimo())
	require.EqualValues(t, 3, lista.Largo())

	iter = lista.Iterador()
	require.EqualValues(t, 1, iter.VerActual())
	iter.Siguiente()
	require.EqualValues(t, 2, iter.VerActual())
	iter.Siguiente()
	require.EqualValues(t, 3, iter.VerActual())

	iter.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
}

func TestIteradorInsercionMedio(t *testing.T) {
	t.Log("Prueba: insertar un elemento en el medio")
	lista := TDAlista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	iter.Siguiente()

	iter.Insertar(2)
	require.EqualValues(t, 3, lista.Largo())

	iter = lista.Iterador()
	require.EqualValues(t, 1, iter.VerActual())
	iter.Siguiente()
	require.EqualValues(t, 2, iter.VerActual())
	iter.Siguiente()
	require.EqualValues(t, 3, iter.VerActual())

	listaVacia := TDAlista.CrearListaEnlazada[int]()
	iterVacio := listaVacia.Iterador()

	require.False(t, iterVacio.HaySiguiente())

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterVacio.VerActual() })
}

func TestIteradorHaySiguiente(t *testing.T) {
	t.Log("Prueba: comportamiento del método HaySiguiente")
	lista := TDAlista.CrearListaEnlazada[int]()

	iter := lista.Iterador()
	require.False(t, iter.HaySiguiente(), "En lista vacía no debe haber siguiente")

	lista.InsertarUltimo(1)
	iter = lista.Iterador()
	require.True(t, iter.HaySiguiente(), "En lista con un elemento debe haber siguiente")
	iter.Siguiente()
	require.False(t, iter.HaySiguiente(), "Al llegar al final no debe haber siguiente")

	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iter = lista.Iterador()
	require.True(t, iter.HaySiguiente())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())

	iter = lista.Iterador()
	iter.Borrar()
	require.True(t, iter.HaySiguiente(), "Después de borrar debe seguir habiendo elementos")
	iter.Siguiente()
	iter.Borrar()
	require.False(t, iter.HaySiguiente(), "Después de borrar el último no debe haber siguiente")
}

func TestIteradorBorraListaConUnicoElemento(t *testing.T) {
	t.Log("Prueba de borrar único elemento con iterador")
	lista := TDAlista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(123)
	iter := lista.Iterador()
	require.Equal(t, 123, iter.Borrar())

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())

	// Verificar que no se puede hacer nada despues de que la lista este vacia
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })

	// Agregar mas elementos y realizar nuevas verificaciones
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	iter = lista.Iterador()

	// Iterar y borrar de nuevo
	require.Equal(t, 1, iter.Borrar())
	require.Equal(t, 2, iter.Borrar())

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
}
func TestCantidadOperacionesIterador_InsertarPrimero(t *testing.T) {
	t.Log("Prueba de insertar primero")
	lista := TDAlista.CrearListaEnlazada[int]()
	maxIter := 100000

	for i := 1; i <= maxIter; i++ {
		lista.InsertarPrimero(i)
	}
	require.EqualValues(t, maxIter, lista.Largo())

	contador := maxIter
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		require.EqualValues(t, contador, iter.VerActual())
		contador--
	}
	require.EqualValues(t, 0, contador)
}

func TestCantidadOperacionesIterador_InsertarPrimeroConBorrarPrimero(t *testing.T) {
	t.Log("Prueba de insertar primero con borrar primero")
	lista := TDAlista.CrearListaEnlazada[int]()
	maxIter := 100000

	for i := 1; i <= maxIter; i++ {
		lista.InsertarPrimero(i)
	}

	for i := maxIter; i >= 1; i-- {
		require.EqualValues(t, i, lista.BorrarPrimero())
	}

	require.True(t, lista.EstaVacia())
}

func TestCantidadOperacionesIterador_InsertarUltimoConBorrarPrimero(t *testing.T) {
	t.Log("Prueba de insertar último con borrar primero")
	lista := TDAlista.CrearListaEnlazada[int]()
	maxIter := 100000

	for i := 1; i <= maxIter; i++ {
		lista.InsertarUltimo(i)
	}

	for i := 1; i <= maxIter; i++ {
		require.EqualValues(t, i, lista.BorrarPrimero())
	}

	require.True(t, lista.EstaVacia())
}

func TestCantidadOperacionesIterador_IteradorExternoConInsertarYBorrarPrimero(t *testing.T) {
	t.Log("Prueba de iterador externo con insertar y borrar primero")
	lista := TDAlista.CrearListaEnlazada[int]()
	maxIter := 100000

	for i := 1; i <= maxIter; i++ {
		lista.InsertarUltimo(i)
	}

	iter := lista.Iterador()
	for iter.HaySiguiente() {
		require.EqualValues(t, iter.VerActual(), lista.BorrarPrimero())
		iter.Siguiente()
	}

	require.True(t, lista.EstaVacia())
}

func TestCantidadOperacionesIterador_IteradorExternoConInsertarYBorrar(t *testing.T) {
	t.Log("Prueba de iterador externo con insertar y borrar")
	lista := TDAlista.CrearListaEnlazada[int]()
	maxIter := 100000

	for i := 1; i <= maxIter; i++ {
		lista.InsertarUltimo(i)
	}

	iter := lista.Iterador()
	for iter.HaySiguiente() {
		require.EqualValues(t, iter.VerActual(), lista.BorrarPrimero())
		iter.Siguiente()
	}

	require.True(t, lista.EstaVacia())
}

func TestCantidadOperacionesIteradorInsertarPrimero(t *testing.T) {
	t.Log("Prueba de volumen: solo insertar primero")
	lista := TDAlista.CrearListaEnlazada[int]()
	maxIter := 100000

	for i := 1; i <= maxIter; i++ {
		lista.InsertarPrimero(i)
	}
	require.EqualValues(t, maxIter, lista.Largo())
}

func TestCantidadOperacionesIteradorInsertarUltimo(t *testing.T) {
	t.Log("Prueba de volumen: solo insertar último")
	lista := TDAlista.CrearListaEnlazada[int]()
	maxIter := 100000

	for i := 1; i <= maxIter; i++ {
		lista.InsertarUltimo(i)
	}
	require.EqualValues(t, maxIter, lista.Largo())
}
