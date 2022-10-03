Un ejemplo sencillo de como se definiría una funcion, seria de este modo: 
```go
func <nombre Funcion>(){
    //codigo a ejecutar en cada invocacion
}
```
Por ejemplo, el siguiente código definiría una funcion llamada Hola que, al invocarse, muestra por pantalla el mensaje ¡Hola!:
```go
func.Hola(){
    fmt.Println("Hola")
}
```
Las funciones se definen fuera del cuerpo de la funcion main, y suelen invocarse desde el codigo de las propias funciones, mediante el nombre de la funcion seguido de dos parentesis ():
```go
func main(){
    fmt.Println("Invocando una funcion: ")
    Hola()
    fmt.Println("Invocando la misma funcion, otra vez: ")
    Hola()
}
```
Si se añade la definicion de la funcion Hola al mismo archivo que contiene la funcion main anterior y se ejecuta el programa, la salida estandar seria:
```go
Invocando una funcion:
¡Hola!
Invocando la misma funcion, otra vez:
¡Hola!
```
Aunque una funcion ejecuta siempre el mismo codigo, se puede modificar ligeramente su comportamiento si se definen **argumentos** (o **parametros**) entre los parentesis de la cabecera de la funcion. Cada argumento consiste en un nombre seguido de un tipo (como si de una variable se tratara). Los diferentes argumentos se separan por comas.

Por ejemplo, se podria modificar el comportamiento de la anterior funcion Hola para especificar el nombre y apellidos de la persona a saludar mediante argumentos:
```go
func Hola(nombre string, apellido string){
    fmt.Printf("¡Hola, %s %s!\n", nombre, apellido)
}
```