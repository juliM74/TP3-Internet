package cola_prioridad_test

import (
	"math/rand"
	"testing"

	cola "tp3/tdaGrafo/tdas/cola_prioridad"

	"github.com/stretchr/testify/require"
)

func mezclarSlice(slice []int) {
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func compararInt(a, b int) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

func max(slice []int) int {
	m := slice[0]
	for _, v := range slice {
		if v > m {
			m = v
		}
	}
	return m
}

func TestHeapVolumenGrande(t *testing.T) {
	const n = 50000
	elementos := make([]int, n)
	for i := 0; i < n; i++ {
		elementos[i] = i
	}
	mezclarSlice(elementos)
	heap := cola.CrearHeap(compararInt)
	maximo := -1
	for i, v := range elementos {
		heap.Encolar(v)
		if v > maximo {
			maximo = v
		}
		require.Equal(t, maximo, heap.VerMax())
		require.Equal(t, i+1, heap.Cantidad())
	}
	for i := n - 1; i >= 0; i-- {
		require.Equal(t, i, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
}

func TestHeapConNegativos(t *testing.T) {
	heap := cola.CrearHeap(compararInt)
	elementos := []int{-10, -20, -5, -30, -25}
	for _, v := range elementos {
		heap.Encolar(v)
	}
	require.Equal(t, -5, heap.VerMax())
	require.Equal(t, -5, heap.Desencolar())
	require.Equal(t, -10, heap.Desencolar())
	require.Equal(t, -20, heap.Desencolar())
	require.Equal(t, -25, heap.Desencolar())
	require.Equal(t, -30, heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func TestHeapSortNegativosYRepetidos(t *testing.T) {
	arr := []int{3, -1, 2, -1, 3, 0}
	cola.HeapSort(arr, compararInt)
	require.Equal(t, []int{-1, -1, 0, 2, 3, 3}, arr)
}

func TestEncolarElementos(t *testing.T) {
	heap := cola.CrearHeap(compararInt)
	elementos := []int{10, 20, 5, 30, 25}
	for i, elem := range elementos {
		heap.Encolar(elem)
		require.False(t, heap.EstaVacia())
		require.Equal(t, i+1, heap.Cantidad())
		require.Equal(t, max(elementos[:i+1]), heap.VerMax())
	}
}

func TestDesencolarElementos(t *testing.T) {
	heap := cola.CrearHeap(compararInt)
	elementos := []int{10, 20, 5, 30, 25}
	for _, elem := range elementos {
		heap.Encolar(elem)
	}
	esperados := []int{30, 25, 20, 10, 5}
	for _, esperado := range esperados {
		require.Equal(t, esperado, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
}

func TestAlternarEncolarDesencolar(t *testing.T) {
	heap := cola.CrearHeap(compararInt)
	heap.Encolar(10)
	require.Equal(t, 10, heap.VerMax())
	require.Equal(t, 10, heap.Desencolar())
	heap.Encolar(20)
	require.Equal(t, 20, heap.VerMax())
	heap.Encolar(5)
	require.Equal(t, 20, heap.VerMax())
	require.Equal(t, 20, heap.Desencolar())
	require.Equal(t, 5, heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func TestHeapifyTodosIguales(t *testing.T) {
	elementos := []int{7, 7, 7, 7, 7}
	heap := cola.CrearHeapArr(elementos, compararInt)
	require.Equal(t, 5, heap.Cantidad())
	for i := 0; i < 5; i++ {
		require.Equal(t, 7, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
}

func TestHeapVacioYUnElemento(t *testing.T) {
	heap := cola.CrearHeap(compararInt)
	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })

	heap.Encolar(99)
	require.False(t, heap.EstaVacia())
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 99, heap.VerMax())
	require.Equal(t, 99, heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func TestHeapifyCasosCompletos(t *testing.T) {
	heapVacio := cola.CrearHeapArr([]int{}, compararInt)
	require.True(t, heapVacio.EstaVacia())
	require.Equal(t, 0, heapVacio.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapVacio.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heapVacio.Desencolar() })

	heapUno := cola.CrearHeapArr([]int{7}, compararInt)
	require.False(t, heapUno.EstaVacia())
	require.Equal(t, 1, heapUno.Cantidad())
	require.Equal(t, 7, heapUno.VerMax())
	require.Equal(t, 7, heapUno.Desencolar())
	require.True(t, heapUno.EstaVacia())

	elementos := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mezclarSlice(elementos)
	heap := cola.CrearHeapArr(elementos, compararInt)
	require.Equal(t, len(elementos), heap.Cantidad())
	require.Equal(t, 10, heap.VerMax())
	for i := 10; i >= 1; i-- {
		require.Equal(t, i, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
}

func TestHeapConElementosRepetidos(t *testing.T) {
	heap := cola.CrearHeap(compararInt)
	elementos := []int{5, 1, 5, 3, 5}
	for _, v := range elementos {
		heap.Encolar(v)
	}
	require.Equal(t, 5, heap.VerMax())
	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 3, heap.Desencolar())
	require.Equal(t, 1, heap.Desencolar())
	require.True(t, heap.EstaVacia())
}
