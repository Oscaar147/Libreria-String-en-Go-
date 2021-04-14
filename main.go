package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(strings.Contains("distribuidos", "bui"))  //existe bui en la cadena
	//fmt.Println(strings.Count("distribuidos", "i"))   //numero de veces se repite la i
	//fmt.Println(strings.HasPrefix("distribuidos", "di")) //inicia con di
	//fmt.Println(strings.HasSuffix("distribuidos", "os")) // termina os
	//fmt.Println(strings.Index("distribuidos", "as"))  //retorna un boleano si existe o no la subcadena as

	// fmt.Println(strings.Join([]string{"Sistemas", "Distribuidos", "CUCEI"}, "-")) //une las cadenas puede ser con un guion etc espacio si no pones nada al ultimo queda pegada
	//fmt.Println(strings.Repeat("distribudos", 2)) //repite la cadena las veces que pongas
	//fmt.Println(strings.Replace("aaaaabbb", "a", "b", 2)) // remplaza la segunda a por b
	fmt.Println(strings.Split("Mi mama me mima", "-")) //separa el slice por espacios
	// fmt.Println(strings.ToLower("Mi mama me mima"))
	// fmt.Println(strings.ToUpper("Mi mama me mima"))

	b := []byte("test")
	fmt.Println(b)

	s := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(s)
}
