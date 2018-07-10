package main

import (
	"fmt"
	"math/rand"
	"time"
)

var colores = []string{"Verde", "Azul", "Negro", "Blanco"}
var paises = []string{"sp", "fr", "be", "de", "it", "at",
	"hr", "mk", "bg", "rs", "ba", "si",
	"sk", "ua", "ro", "fi", "no",
	"lv", "ru", "tr", "ge", "pt", "an", "lu", "ch", "nl", "dk",
	"pl", "cz", "li", "hu", "me",
	"al", "el", "by", "md", "mo", "se",
	"ee", "az", "am"}

var vecinos = map[string][]string{
	"sp": {"pt", "fr", "an"},
	"fr": {"it", "be", "lu", "ch", "de", "an"},
	"be": {"lu", "nl"},
	"de": {"be", "nl", "lu", "dk", "pl", "cz", "at", "ch", "li"},
	"it": {"ch", "si", "at"},
	"at": {"ch", "li", "si", "hu", "sk"},
	"hr": {"hu", "si", "ba", "rs", "me"},
	"mk": {"al", "bg", "el", "rs"},
	"bg": {"el", "ro", "tr"},
	"rs": {"me", "al", "ro", "hu", "ba"},
	"ba": {"me"},
	"si": {"hu"},
	"sk": {"hu", "cz", "ua", "pl"},
	"pl": {"ru", "lt", "by", "ua", "cz"},
	"ua": {"ro", "by", "ru", "md"},
	"ro": {"mo", "hu"},
	"fi": {"se", "no", "ru"},
	"no": {"se", "ru"},
	"lv": {"ru", "lt", "ee", "bg"},
	"ru": {"ee", "lt", "bg", "ge", "az"},
	"tr": {"am", "ge"},
	"ge": {"az"},
}
var colorEstado = map[string]string{
	"sp": "", 	"fr": "", 	"be": "", 	"de": "", 	"it": "", 	"at": "", 	"hr": "", 	"mk": "", 	"bg": "", 	"rs": "", "ba": "", "si": "", "sk": "", "pl": "", "ua": "", "ro": "", "fi": "", "no": "", "lv": "", "ru": "", "tr": "", "ge": "", "pt": "", "an": "", "lu": "", "ch": "", "nl": "", "dk": "", "cz": "", "li": "", "hu": "", "me": "", "al": "", "el": "", "by": "", "md": "", "mo": "", "se": "", "ee": "", "az": "", "am": "",
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	start := time.Now()
	asignarColores()

	elapsed := time.Since(start)

	fmt.Println("Tiempo de corrida: ", elapsed)
}

func ponerColor(estado string, color string) bool {
	for _, v := range vecinos[estado] {
		colorVecino := colorEstado[v]
		if colorVecino == color {
			return false
		}
	}
	return true
}

func asignarColores() {

	for i := 0; i < 3; i++ {
		for _, estado := range paises {
			colorEstado[estado] = obtenerColor(estado)
		}
	}

	fmt.Println(colorEstado)
}

func obtenerColor(es string) string {
	for _, c := range colores {
		if ponerColor(es, c) == true {
			return c
		}
	}
	return " "
}
