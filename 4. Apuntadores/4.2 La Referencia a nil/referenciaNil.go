package main

import "fmt"

func main() {
	var pi *int
	pi = nil
	if pi == nil {
		fmt.Print("¡No puedo hacer nada con este apuntador!")
		fmt.Println("porque no apunta a nada")
	}
}
