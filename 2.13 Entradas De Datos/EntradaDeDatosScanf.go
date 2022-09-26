package main

import "fmt"

func main() {
	var i, j int

	fmt.Print("Type two numbers: ")
	fmt.Scanf("%v %v", &i, &j)
	fmt.Println("Your numbers are:", i, "and", j)
}
