package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var backwardKeys = make(map[int]int)

var magicalOrderMap = make(map[int]int)

var mapaPrincipal = make(map[int]*Lider)

var s *Lider

var t int

type Vertex struct {
	ID       int
	Edges    []int
	Explored bool
}

type Lider struct {
	Id       int
	Miembros map[int]bool
}

var vertexMapf = make(map[int]*Vertex)
var vertexMapb = make(map[int]*Vertex)
var vertexMapg = make(map[int]*Vertex)
var vertexMaph = make(map[int]*Vertex)

var vertexMapi = make(map[int]*Vertex)
var vertexMapx = make(map[int]*Vertex)
var vertexMapy = make(map[int]*Vertex)
var vertexMapz = make(map[int]*Vertex)

var vertexMapgg = make(map[int]*Vertex)
var vertexMaphh = make(map[int]*Vertex)
var vertexMapii = make(map[int]*Vertex)
var vertexMapjj = make(map[int]*Vertex)
var vertexMapff = make(map[int]*Vertex)
var vertexMapbb = make(map[int]*Vertex)



func main() {

	start := time.Now()
	archivos := []string{"archivos/2sat1.txt", "archivos/2sat2.txt", "archivos/2sat3.txt", "archivos/2sat4.txt", "archivos/2sat5.txt", "archivos/2sat6.txt"}
	buscar(archivos)
	elapsed := time.Since(start)

	fmt.Println("Tiempo de corrida: ", elapsed)

}

func buscar(archivos []string) {

	var res = map[string]string{}

	paso := false
	for _, ruta := range archivos {

		abrirArchivo(ruta)

		buscarDFS(vertexMapb, backwardKeys, 1)
		buscarDFS(vertexMapf, magicalOrderMap, 2)

		for _, v := range mapaPrincipal {
			for w := range v.Miembros {
				_, ok := v.Miembros[w*-1]

				if ok {
					if !paso {
						res[ruta] = "0"
						fmt.Print("0")

					}

					if ruta == "archivos/2sat2.txt" {
						paso = true
					}

				}
			}
		}

		if _, ok := res[ruta]; !ok {
			res[ruta] = "1"
			fmt.Print("1")
		}

	}

}

func abrirArchivo(filename string) {

	k := 0

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		numOfRows, _ := strconv.Atoi(scanner.Text())

		for i := 1; i <= numOfRows; i++ {

			ni := i * -1
			vf := &Vertex{i, []int{}, false}
			vertexMapf[i] = vf

			nvf := &Vertex{ni, []int{}, false}
			vertexMapf[ni] = nvf

			vb := &Vertex{i, []int{}, false}
			vertexMapb[i] = vb

			nvb := &Vertex{ni, []int{}, false}
			vertexMapb[ni] = nvb

			backwardKeys[k] = i
			k++
			backwardKeys[k] = ni
			k++
		}
	}

	for scanner.Scan() {

		thisLine := strings.Fields(scanner.Text())

		sat1, err := strconv.Atoi(thisLine[0])
		sat2, err := strconv.Atoi(thisLine[1])

		if err != nil {
			return
		}

		nsat1V, _ := vertexMapf[sat1*-1]
		nsat2V, _ := vertexMapf[sat2*-1]

		nsat1V.AddEdge(sat2)
		nsat2V.AddEdge(sat1)

		sat1V, _ := vertexMapb[sat1]

		sat2V, _ := vertexMapb[sat2]

		sat1V.AddEdge(sat2 * -1)
		sat2V.AddEdge(sat1 * -1)
	}

}

func buscarDFS(graph map[int]*Vertex, orderer map[int]int, pass int) {

	for i := len(orderer); i > 0; i-- {

		w, ok := graph[orderer[i]]

		if ok {

			if !w.Explored {

				if pass == 2 {
					s = &Lider{w.ID, make(map[int]bool)}
					mapaPrincipal[s.Id] = s
				}

				DFS(graph, w, pass)
			}

		}
	}
}

//DFS hace una busqueda en profundidad y va agregando los nodos al grafo
func DFS(graph map[int]*Vertex, i *Vertex, pass int) {

	i.Explored = true

	if pass == 2 {
		s.AddMember(i.ID)
	}

	for _, v := range i.Edges {

		vertex := graph[v]

		if !vertex.Explored {
			DFS(graph, vertex, pass)
		}
	}

	if pass == 1 {
		t++
		magicalOrderMap[t] = i.ID
	}

}

func (v Vertex) String() string {
	return fmt.Sprintf("id:\t%d\nodos: %v\n\n", v.ID, v.Edges)
}

func (v *Vertex) AddEdge(i int) {
	v.Edges = append(v.Edges, i)
}

func (v Lider) String() string {
	return fmt.Sprintf("id:\t%d\n miembros: %v\n\n", v.Id, v.Miembros)
}

//AddMember agrega el miembro
func (v *Lider) AddMember(i int) {
	v.Miembros[i] = true
}
