package main

import (
	"fmt"
)

func main() {
	fmt.Print("Escribe un caracter: ")
	var c int8
	fmt.Scanf("%c", &c)

	switch {
	case c >= 'A' && c <= 'Z':
		fmt.Println("Letra mayuscula")
	case c >= 'a' && c <= 'z':
		fmt.Println("Letra minuscula")
	case c >= '0' && c <= '9':
		fmt.Println("Digito")
	default:
		fmt.Println("Ni letra ni digito")
	}
}
