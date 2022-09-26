En cada instruccion condicional o cada bucle, cada vez que se define un nuevo contexto entre un par de llaves { y } es posible definir nuevas variables cuya vida y visibilidad se limitaran a ese contexto:
- Ver codigo:
```sh
llaves.go
```
Si, en el ejemplo anterior, intentara mostrar el contenido de la variable i desde fuera del bloque if dondeha sido definida, el compilador le mostraria un mensaje de error.

Una caracteristica de Go, que puede ser util pero tambien puede ser una fuente de problemas, es el poder definir nuevas variables con nombres que ya existen en los contextos mas globales. A pesar de tener el mismo nombre, seran variables distintas. Este concepto es conocido como "ocultacion" o,en ingles, shadowing.
Ejemplo:
```sh
a := 0
b := 0
if true{
    a := 1
    b = 1
    a++
    b++
}
fmt.Printf("a = %d, b = %d\n", a ,b)
```
En un despiste, se podria pensar que la salida estandar de este programa seria a = 2, b = 2. Sin embargo, es a = 0, b = 2, ya que la variable a sido redeclarada y ocultada (eclipsada) dentro de la condicion (observe el sutil detalle del operador de declaracion := usado con a, frente al de asignacion = usado con b). Por tanto, la variable a mostrada en fmt.Printf no ha sido modificada en ningun momento.