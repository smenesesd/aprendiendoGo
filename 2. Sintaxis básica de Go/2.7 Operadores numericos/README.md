Go permite hacer las siguientes operaciones con tipos de datos numéricos. Por orden de precedencia: 

1. Agrupaciones de operaciones, delimitadas por paréntesis.
2. Multiplicaciones, divisiones (operadores * y /), así como el resto de la división entera (o módulo, operador %).
3. Sumas y restas (operadores + y -). 

El orden de procedencia hace referencia a que operaciones se evalúan primero, cuando una expresión compleja engloba múltiples operaciones. Primero, se evalúan las operaciones de mayor precedencia. En este caso de múltiples operaciones con la misma precedencia, estas de evalúan según su posición en la expresión, de izquierda a derecha.

Por ejemplo, dada la siguiente expresión:
```go
 a := 8 + 3 * (1 + 2) % 5
```
1. 1. Primero evaluaria la expresion entre parentesis, ya que es la de mayor precedencia: 
    
    a := 8 + 3 * 3 % 5
    
2. Luego evaluaria la multiplicacion y el modulo. Al ser de la misma precedencia, primero evaluaria la multiplicacion, ya que está más a la izquierda:
    
    a := 8 + 9 % 5
    
3. Y continuaria por el resto de la division entera:
    
    a := 8 + 9
    
4. Siendo la suma la operacion de menos precedencia, seria la ultima en evaluarse:
    
    a := 12

Ademas los anteriores operadores matematicos, Go provee los operadores de incremento(++) y decremento(--), que van detras de una variable que se quiere incrementar o decrementar, respectivamente:
```go
a := 10
b := 20
a++
b--
```
Despues de ejecutar el programa anterior, la variable a contendria el valor de 11 y la variable b el valor de 19.