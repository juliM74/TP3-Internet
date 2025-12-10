package utils

import (
	"fmt"
	"strconv"
	"strings"
	biblioteca "tp3/biblioteca"
	"tp3/internet/estado"
)

// ============================= EL DESPACHADOR PRINCIPAL  =============================

func EjecutarLinea(est estado.Estado, comando string, parametros string) {

	switch comando {

	case "listar_operaciones":
		ejecutarListarOperaciones()

	case "camino":
		ejecutarCamino(est, parametros)

	case "mas_importantes":
		ejecutarMasImportantes(est, parametros)

	case "conectados":
		ejecutarConectados(est, parametros)

	case "ciclo":
		ejecutarCiclo(est, parametros)

	case "lectura":
		ejecutarLectura(est, parametros)

	case "diametro":
		ejecutarDiametro(est)

	case "rango":
		ejecutarRango(est, parametros)

	case "comunidad":
		ejecutarComunidad(est, parametros)

	case "navegacion":
		ejecutarNavegacion(est, parametros)

	case "clustering":
		ejecutarClustering(est, parametros)

	default:
		fmt.Println("Comando desconocido:", comando)
	}
}

// ======================================== listar_operaciones (O(1)) ========================================

func ejecutarListarOperaciones() {
	fmt.Println("camino")
	fmt.Println("mas_importantes")
	fmt.Println("conectados")
	fmt.Println("ciclo")
	fmt.Println("lectura")
	fmt.Println("diametro")
	fmt.Println("rango")
	fmt.Println("comunidad")
	fmt.Println("navegacion")
	fmt.Println("clustering")
}

// ======================================== camino (BFS) ========================================

func ejecutarCamino(est estado.Estado, parametros string) {

	parts := strings.Split(parametros, ",")
	if len(parts) != 2 {
		fmt.Println("No se encontro recorrido")
		return
	}

	origen := strings.TrimSpace(parts[0])
	destino := strings.TrimSpace(parts[1])

	if !est.Grafo().Pertenece(origen) || !est.Grafo().Pertenece(destino) {
		fmt.Println("No se encontro recorrido")
		return
	}

	camino := biblioteca.CaminoMinimoBFS(est.Grafo(), origen, destino, strings.EqualFold)

	if camino == nil {
		fmt.Println("No se encontro recorrido")
		return
	}

	// imprimir camino
	for i := range camino {
		if i > 0 {
			fmt.Print(" -> ")
		}
		fmt.Print(camino[i])
	}
	fmt.Println()
	fmt.Println("Costo:", len(camino)-1)
}

// ======================================== mas_importantes (PageRank) ========================================

func cmpFloat(a, b float64) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func ejecutarMasImportantes(est estado.Estado, parametros string) {
	n, err := strconv.Atoi(strings.TrimSpace(parametros))
	if err != nil {
		return
	}

	if !est.TienePagerank() {
		pr := biblioteca.PageRank(est.Grafo(), 0.85, 15, strings.EqualFold)
		pr.Iterar(func(pag string, val float64) bool {
			est.GuardarPagerank(pag, val)
			return true
		})
		est.MarcarPagerankCalculado()
	}

	top := biblioteca.TopN(est.IterarPagerank, n, cmpFloat)

	for i := range top {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(top[i])
	}
	fmt.Println()
}

// ======================================== conectados (CFC) ========================================

func ejecutarConectados(est estado.Estado, parametros string) {

	pagina := strings.TrimSpace(parametros)

	if !est.Grafo().Pertenece(pagina) {
		// Si no existe, no imprimimos nada
		return
	}

	if est.TieneCFC(pagina) {
		for _, v := range est.ObtenerCFC(pagina) {
			fmt.Println(v)
		}
		return
	}

	cfc := biblioteca.CFCSoloDe(est.Grafo(), pagina, strings.EqualFold)
	est.GuardarCFC(pagina, cfc)

	for _, v := range cfc {
		fmt.Println(v)
	}
}

// ======================================== ciclo ========================================

func ejecutarCiclo(est estado.Estado, parametros string) {

	parts := strings.Split(parametros, ",")
	if len(parts) != 2 {
		fmt.Println("No se encontro recorrido")
		return
	}

	pagina := strings.TrimSpace(parts[0])
	n, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

	if !est.Grafo().Pertenece(pagina) {
		fmt.Println("No se encontro recorrido")
		return
	}

	ciclo := biblioteca.CicloLargoN(est.Grafo(), pagina, n, strings.EqualFold)

	if ciclo == nil {
		fmt.Println("No se encontro recorrido")
		return
	}

	for i := range ciclo {
		if i > 0 {
			fmt.Print(" -> ")
		}
		fmt.Print(ciclo[i])
	}
	fmt.Println()
}

// ======================================== lectura ========================================

func ejecutarLectura(est estado.Estado, parametros string) {

	parts := strings.Split(parametros, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
		if !est.Grafo().Pertenece(parts[i]) {
			fmt.Println("No existe forma de leer las paginas en orden")
			return
		}
	}

	orden := biblioteca.Lectura2am(est.Grafo(), parts, strings.EqualFold)

	if orden == nil {
		fmt.Println("No existe forma de leer las paginas en orden")
		return
	}

	fmt.Println(strings.Join(orden, ", "))
}

// ======================================== diametro (BFS desde todos) ========================================

func ejecutarDiametro(est estado.Estado) {
	if est.TieneDiametro() {
		camino := est.ObtenerDiametro()
		imprimirCaminoDiametro(camino)
		return
	}

	camino := biblioteca.Diametro(est.Grafo(), strings.EqualFold)
	est.GuardarDiametro(camino)

	imprimirCaminoDiametro(camino)
}

func imprimirCaminoDiametro(camino []string) {
	for i := range camino {
		if i > 0 {
			fmt.Print(" -> ")
		}
		fmt.Print(camino[i])
	}
	fmt.Println()
	fmt.Println("Costo:", len(camino)-1)
}

// ======================================== rango (BFS) ========================================

func ejecutarRango(est estado.Estado, parametros string) {

	parts := strings.Split(parametros, ",")
	if len(parts) != 2 {
		return
	}

	pagina := strings.TrimSpace(parts[0])
	n, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

	cant := biblioteca.CantidadEnRango(est.Grafo(), pagina, n, strings.EqualFold)

	fmt.Println(cant)
}

// ======================================== comunidad ========================================

func ejecutarComunidad(est estado.Estado, parametros string) {

	pagina := strings.TrimSpace(parametros)

	if !est.TieneComunidades() {
		comunidades := biblioteca.LabelPropagation(est.Grafo(), strings.EqualFold)
		comunidades.Iterar(func(p string, e int) bool {
			est.GuardarEtiqueta(p, e)
			return true
		})
		est.MarcarComunidadesCalculadas()
	}

	etiq, ok := est.ObtenerEtiqueta(pagina)
	if !ok {
		return
	}

	est.IterarEtiquetas(func(p string, e int) {
		if e == etiq {
			fmt.Println(p)
		}
	})
}

// ======================================== navegacion (Primer link) ========================================

func ejecutarNavegacion(est estado.Estado, parametros string) {
	origen := strings.TrimSpace(parametros)

	if !est.Grafo().Pertenece(origen) {
		fmt.Println(origen)
		return
	}

	cam := biblioteca.PrimerLink(est.Grafo(), origen, strings.EqualFold)

	if len(cam) == 0 {
		fmt.Println(origen)
		return
	}

	fmt.Println(strings.Join(cam, " -> "))
}

// ======================================== clustering (local y global) ========================================

func ejecutarClustering(est estado.Estado, parametros string) {
	param := strings.TrimSpace(parametros)

	// global
	if param == "" {

		if est.TieneClusteringPromedio() {
			fmt.Printf("%.3f\n", est.ObtenerClusteringPromedio())
			return
		}

		val := biblioteca.ClusteringPromedio(est.Grafo(), strings.EqualFold)
		est.GuardarClusteringPromedio(val)

		fmt.Printf("%.3f\n", val)
		return
	}

	// local por página
	pagina := param

	if est.TieneClusteringLocal(pagina) {
		fmt.Printf("%.3f\n", est.ObtenerClusteringLocal(pagina))
		return
	}

	val := biblioteca.ClusteringVertice(est.Grafo(), pagina, strings.EqualFold)
	est.GuardarClusteringLocal(pagina, val)

	fmt.Printf("%.3f\n", val)
}
