package main

import (
	"fmt"
)

func main() {
	a := 3
	if a == 3 {
		i := 1
		fmt.Println("'i' solo es visibles en este contexto:", i)
	}
	fmt.Println("'a' es visible en todo 'main':", a)
}
