package main

import "fmt"

var hora, minuto, segundos int

func main() {

	fmt.Print("HH:MM:SS? ")
	fmt.Scanf("%d:%d:%d", &hora, &minuto, &segundos)
	fmt.Println("%d horas, %d minutos, %d segundos", hora, minuto, segundos)
}
