package main

import "fmt"

// Definicion de variables, llamadas dias y meses
// Y a meses se le asigna un valor entero = 12
var dias int
var meses = 12

/* Si no se provee un valor inicial, las variables serán inicializadas
automáticamente con el “valor cero” de cada tipo de datos:
 0 para tipos numéricos,false para bool, y la cadena vacía ”” para string.*/

func main() {
	fmt.Println(dias, "Y", meses)

}
