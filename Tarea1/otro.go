package main

import "fmt"

var colores = []string{"Red", "Blue", "Green", "Yellow", "Black"}
var estados = []string{"sp", "fr", "be", "de", "it", "at",
	"hr", "mk", "bg", "rs", "ba", "si",
	"sk", "ua", "ro", "fi", "no",
	"lv", "ru", "tr", "ge","pt", "an", "lu", "ch", "nl", "dk",
	"pl", "cz", "li", "hu", "me",
	"al", "el", "by", "md", "mo", "se",
	"ee", "az", "am",}
var vecinos = map[string][]string{
	"sp" : {"pt", "fr", "an"},
	"fr" : {"it", "be", "lu", "ch", "de", "an"},
	"be" : {"lu","nl"},
	"de" : {"be", "nl", "lu", "dk", "pl", "cz", "at", "ch", "li"},
	"it" : {"ch", "si", "at"},
	"at" : {"ch", "li", "si","hu","sk"},
	"hr" : {"hu", "si", "ba", "rs", "me"},
	"mk" : {"al", "bg", "el", "rs"},
	"bg" : {"el","ro", "tr"},
	"rs" : {"me", "al", "ro", "hu", "ba"},
	"ba" : {"me"},
	"si" : {"hu"},
	"sk" : {"hu", "cz", "ua", "pl"},
	"pl" : {"ru", "lt", "by", "ua", "cz"},
	"ua" : {"ro", "by", "ru" ,"md"},
	"ro" : {"mo", "hu"},
	"fi" : {"se", "no", "ru"},
	"no" : {"se", "ru"},
	"lv" : {"ru", "lt", "ee", "bg"},
	"ru" : {"ee", "lt", "bg", "ge", "az"},
	"tr" : {"am", "ge"},
	"ge" : {"az"},
}
var color_del_estado = map[string]string{
	"sp": "",
	"fr": "",
	"be": "",
	"de": "",
	"it": "",
	"at": "",
	"hr": "",
	"mk": "",
	"bg": "",
	"rs": "",
	"ba": "",
	"si": "",
	"sk": "",
	"pl": "",
	"ua": "",
	"ro": "",
	"fi": "",
	"no": "",
	"lv": "",
	"ru": "",
	"tr": "",
	"ge": "",
	"pt": "",
	"an": "",
	"lu": "",
	"ch": "",
	"nl": "",
	"dk": "",
	"cz": "",
	"li": "",
	"hu": "",
	"me": "",
	"al": "",
	"el": "",
	"by": "",
	"md": "",
	"mo": "",
	"se": "",
	"ee": "",
	"az": "",
	"am": "",
}

func promesa(estado string, color string) bool {
	for i,v := range vecinos[estado] {
		colorVecino := color_del_estado[v]
		if(colorVecino == color) {
			fmt.Println(i)
			return false
		}
	}
	return true
}

func optenrColorEstado(estado string) string {
	for i, c := range colores {
		if promesa(estado, c) == true {
			fmt.Println(i)
			return c
		}
	}
	return " "
}

func ponerColores() {
	for i, estado := range estados {
		color_del_estado[estado] = optenrColorEstado(estado)
		fmt.Println(i)
	}
	fmt.Println(color_del_estado)
}

func main() {
	ponerColores()
}