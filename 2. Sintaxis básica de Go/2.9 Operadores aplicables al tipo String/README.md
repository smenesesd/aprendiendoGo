El operador + aplicado entre string, retorna un nuevo string con las cadenas concatenadas:
```go
a := "el cocherito"
b := "lere"
concat := a + "" + b
```
Ejemplo del codigo real:
- Ver el codigo:
```stringConcatenadas.go ```

En el ejemplo anterior, la variable concat contendria la cadena ```“el cocherito lere”```.
Los operadores de comparacion tambien pueden usarse para comparar cadenas de texto.
### IMPORTANTE: 
- Go hara distincion entre mayusculas y minusculas, por lo que “HOLA” == “hola” retornara false y “ZZZ” < “aaa” tambien retornara false, ya que el valor numerico de la Z mayuscula es menor al de la a minuscula.