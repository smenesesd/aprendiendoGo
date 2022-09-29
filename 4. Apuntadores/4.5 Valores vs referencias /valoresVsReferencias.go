package main

import "fmt"

func main() {
	i := 0
	j := 0
	p1 := &i
	p2 := &j

	if p1 == p2 {
		fmt.Println("p1 y p2 apuntan a la misma direccion")
	} else {
		fmt.Println("p1 y p2 apuntan a direcciones distintas")
	}
	if *p1 == *p2 {
		fmt.Println("p1 y p2 apuntan a valores iguales")
	} else {
		fmt.Println("p1 y p2 apuntan a valores distintos")
	}

}
