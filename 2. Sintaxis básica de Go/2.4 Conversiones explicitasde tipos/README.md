En go a lo contrario de lo que ocurre en otro lenguajes de tipado estático, si se quiere asignar una variable numérica a otra variable  de otro tipo numérico, se deberá explicitar el tipo de destino:
- Ver codigo: 
```ConversionExplTipos.go````
```go
var segundos int8 = 30
var horas int
horas = int(segundos)
```
- En el codigo anterior, se indica a Go que el valor de la variable segundos, del tipo int8, se va a copiar en otra variable del tipo int. Aunque es obvio que cualquier valor del tipo int8 (8 bits) cabe en una variable del tipo int( 32 o 64 bits), Go obliga a hacer explicita esta conversion.

- Si la conversion se hace desde un tipo entero de tamaño superior al tipo de destino, se usan los bits menos significativos del tipo de origen que caben en el tipo de destino.

- Si la conversion se hace desde un tipo de coma flotante, se trunca la parte de decimales y se asigna la parte entera. Por ejemplo:
```go
distancia := 12.78
kms := int(distancia)
```
- En el ejemplo anterior, la variable kms tomara el valor 12( eliminando, sin redondeo alguno, el 0.78 de la variable original)