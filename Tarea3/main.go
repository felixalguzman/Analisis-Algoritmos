package main

import (
	"time"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)



type aristaDesdeUnVertex struct {
	ID   int
	cost float64
}


type Vertex struct {
	lat  float64
	long float64
}

func (v Vertex) String() string {
	return fmt.Sprintf("\nLatitud:\t%f\nLongitud:\t%f\n\n", v.lat, v.long)
}

var VertexMap = make(map[int]*Vertex)
var VerticesActuales = make(map[int]bool)
var n int

func main() {

	start := time.Now()

	archivos := []string{"archivos/bm33708.txt", "archivos/ar9152.txt", "archivos/ca4663.txt", "archivos/ch71009.txt"}

	calcularDistancias(archivos)
	elapsed := time.Since(start)

	fmt.Println("Tiempo de corrida: ", elapsed)
}

func calcularDistancias(archivos []string) {

	for archivo := range archivos {

		leerArchivo(archivos[archivo])

		var costoDelViaje float64

		aristaActual := 1
		delete(VerticesActuales, 1)
		aristaMasCorta := aristaDesdeUnVertex{0, math.Inf(1)}

		for len(VerticesActuales) > 1 {

			aristaMasCorta.cost = math.Inf(1)

			delete(VerticesActuales, aristaActual)

			for k := range VerticesActuales {

				estaDistancia := distancia(VertexMap[aristaActual], VertexMap[k])

				if estaDistancia < aristaMasCorta.cost {
					aristaMasCorta.ID = k
					aristaMasCorta.cost = estaDistancia
				} else if estaDistancia == aristaMasCorta.cost {
					if k < aristaMasCorta.ID {
						aristaMasCorta.ID = k
						aristaMasCorta.cost = estaDistancia
					}
				}
			}

			costoDelViaje += aristaMasCorta.cost

			if aristaMasCorta.ID == aristaActual {
				break
			} else {
				aristaActual = aristaMasCorta.ID
			}

		}

		fmt.Println("\nEl costo del archivo " + archivos[archivo] + " es ")
		fmt.Print(int(costoDelViaje + distancia(VertexMap[1], VertexMap[aristaMasCorta.ID])))
	}
}

func distancia(a, b *Vertex) float64 {
	return math.Sqrt(math.Pow(a.lat-b.lat, 2) + math.Pow(a.long-b.long, 2))
}

func leerArchivo(filename string) {

	file, _ := os.Open(filename)

	defer file.Close()

	lector := bufio.NewScanner(file)
	
	if lector.Scan() {
		n, _ = strconv.Atoi(lector.Text())
		
	}

	for lector.Scan() {

		thisLine := strings.Fields(lector.Text())

		esteID, _ := strconv.Atoi(thisLine[0])
		estaLatitud, _ := strconv.ParseFloat(thisLine[1], 64)
		estaLongitud, _ := strconv.ParseFloat(thisLine[2], 64)

		VertexMap[esteID] = &Vertex{estaLatitud, estaLongitud}
		VerticesActuales[esteID] = true
	}

	
}
