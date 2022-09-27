El paquete fmt, permite mostrar datos en su terminal de linea de comandos(tambien llamado salida estandar) mediante las siguientes funciones:
```go
fmt.Print
fmt.Println
fmt.Printf
```

```fmt.Print y fmt.Println``` enviara a la salida estandar los datos que situe en parentesis, y separados por comas.

- Para los datos que no sean cadenas de texto, Go hara una conversion generica a texto antes de enviarlos a la salida estandar. Ademas, añadira un espacio entre los diferentes datos dentro de una misma invocacion a ```fmt.Print y fmt.Println```
Ejemplo:
```go
x := 33
fmt.Println("Hola, tu numero es: ", x, "!")
``` 
Y la salida estandar a dicho ejemplo seria:
```go
Hola, tu numero es 33 !
```
la diferencia entre ```fmt.Print y fmt.Println``` es que ```fmt.Println``` añade un salto de linea al final. Es decir, sucesivas invocaciones a ```fmt.Print``` serian mostradas una detras de otra, en la misma linea; mientras que sucesivas invocaciones a ```fmt.Println``` serian mostradas en diferentes lineas, una debajo de la otra.

Cuando requiera un control mas exhaustivo de como se deben mostrar los datos, la funcion  ```fmt.Printf``` admite una cadena de texto en la que puede colocar unas “marcas de formato”, conocidas como verbos. A continuacion, separados por comas, se colocan los datos que Go debe introducir en el lugar de cada uno de los verbos.
```go
cosa := "deposito"
x := 36
y := 84
fmt.Printf("Coordenadas de %v: (%v, %v)", cosa, x, y)
```
Salida estandar: 
```go
Coordenadas de deposito: (36, 84)
```
fmt.Printf no añade ninguna nueva linea al final, por lo que si necesita que el siguiente texto aparezca en la linea siguiente, debe finalizar la cadena de formato de caracter especial de nueva linea: '\n'

Verbo | Descripción
----- | ---
%v    | Valor en el formato por defecto (Tal como se mostraria en fmt.Print)
%T    | Tipo de dato
%d    | Valor numerico en base decimal
%b    | Valor numerico en base binaria
%o    | Valor numerico en base octal
%x    | Valor numerico en base hexadecimal (con letras a-f en minusculas)
%X    | Valor numerico en base hexadecimal (con letras A-F en mayusculas)
%e    | Notacion cientifica (Ej: -1.234e+10)
%f    | Numero en coma flotante, sin exponente(Ej: -13.34)
%c    | Un caracter individual
%s    | Cadena de texto
%q    | Cadena de texto delimitada por comillas dobles, con todos los caracteres especiales substituidos por marcas de Go (Ej: en vez de mostrar un salto de linea, mostrar \n)