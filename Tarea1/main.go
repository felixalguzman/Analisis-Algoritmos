package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"time"
	"math/rand"
)

type paises struct {
	lista []string
	mapa map[string]string
}

 var paise = []*paises{}
 //rojo, verde, amarillo, negro, morado
 var colores = []string{"rojo", "verde", "azul", "negro", "morado"}


func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	start := time.Now()
	leerArchivo("archivo/paises.txt")

	var m map[string]string

	for  i := range paise  {

		 m = make(map[string]string)

		for j := range paise[i].lista{

			pais := paise[i].lista[j]


			//nuevoColor := colores[0]
			//m[pais] = nuevoColor



			if buscarColor(pais, m[pais], i, j) {

				nuevoColor := obtenerColor()
				m[pais] = nuevoColor

			}


		}
		paise[i].mapa = m



	}

	for  i := range paise  {
		fmt.Println(paise[i].lista)
		for j := range paise[i].lista{

			if val, ok := paise[i].mapa[paise[i].lista[j]]; ok {

				fmt.Print( " ", val, " ")
			}
		}

		fmt.Println("")

	}

	elapsed := time.Since(start)

	fmt.Println("Tiempo de corrida: ", elapsed )

}


func obtenerColor() string  {

	var nuevoColor string

	nuevoColor = colores[random(0,4)]

	return nuevoColor
}

func random(min, max int) int {
	return min + rand.Intn(max-min)
}

func buscarColor(pais string, color string, posF int, posC int)bool {
	for i := range paise{

		colorPrimero := paise[i].mapa[paise[i].lista[0]]




			colorPais := paise[i].mapa[paise[i].lista[0]]

			if color == "" {

				return true
			}

			if colorPrimero == color {

				return true
			}
			if i != posF {

				if color == colorPais && pais == paise[i].lista[0] {

					return true
				}
			}







	}

	return false
}


func leerArchivo(ruta string)  {
	archivo, _ := os.Open(ruta)
	defer archivo.Close()
	scanner := bufio.NewScanner(archivo)
	scanner.Split(bufio.ScanLines)

	//n := new(paises)
	for scanner.Scan(){
		//fmt.Println(scanner.Text())

		pais := strings.Split(scanner.Text(), " ")
		n := new(paises)
		n.lista = pais

		paise = append(paise, n)
	}
	fmt.Println("Paises agregados!")


}




