package main

import (
	"fmt"
	"math/rand"
)

func main() {
	valor := rand.Int()
	if valor%2 == 0 {
		//bloque a ejecutar solo si la condicion es true
		fmt.Println("El numero", valor, "es par ")
	} else {
		//bloque a ejecutar solo si la condicion es false
		fmt.Println("El numero", valor, "es impar ")
	}
	fmt.Println("Adios!")

}
