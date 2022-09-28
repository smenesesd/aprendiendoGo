- Crear un archivo cuya extension sea .go, por ejemplo, ```hola.go``` 
```go 
package main
/* fmt (format package):
es el paquete que se encargara de darle formato a cualquier tipo de entrada o salida de datos
- Imprimir sentencias en pantalla
- No sirve para desbuggear¡ERROR!*/

//funcion principal "main"
func main(){
	//muestra un saludo en la pantalla 
	fmt.println("¡Hola, saludo desde go!")
}
``` 
- Ver codigo:

```hola.go```
- El punto de entrada del programa marca mediante func main() {. El programa comenzara ejecutando el código contenido en el bloque delimitado por las llaves {y}.
- Cada fichero esta encabezado por una directiva package, que otorga un nombre común para la agrupación(paquete) de todos los archivos Go en un directorio. El punto de entrada del programa debe de estar encabezado por package main (paquete principal)
- Cuando se incorporan funcionalidades de la biblioteca estandar o de otras bibliotecas de terceros (por ejemplo para mostrar un mensaje en pantalla), deben incorporarse explicitamente los paquetes que las contienen mediante la directiva import. En el programa de ejemplo, la directiva import “fmt” permite usar funcionalidades de escritura de mensajes en pantalla(entre otras).
- Cualquier texto escrito despues de dos horas // es ignorado por el compilador, es usado para poner comentarios en el codigo.
    - Puede usar comentarios de multiples lineas si los situa entre los simbolos /* y */
- El comando fmt.println muestra un mensaje en pantalla. Esta compuesto por el nombre del paquete en el que esta guardado(fmt) y el nombre de la funcion Println. Entre parentesis, se escribe el texto a mostrar(que debe de ir entre comillas dobles).
- Para ejecutar el programa, vaya desde la linea de comandos al directorio que contiene el archivo y ejecute el siguiente comando:
``` go run hola.go```

Para generar un archivo ejecutable que pueda ser ejecutado en otro ordenador, debe utilizar el comando go.build:

```go build hola.go```

Si quisiera generar ejecutables para otros sistemas operativos, puede usar las variables de entorno GOOS y GOARCH.

```GOOS=windows GOARCH=386 go build hola.go```

```GOOS=linux GOARCH=arm64 go build hola.go ```