package main

import (
	"fmt"
)

func main() {
	switch letra {
	case 'A':
		fmt.Println("Mayuscula")
		fallthrough
	case 'a':
		fmt.Println("Primera del abecedario")
	}
}
