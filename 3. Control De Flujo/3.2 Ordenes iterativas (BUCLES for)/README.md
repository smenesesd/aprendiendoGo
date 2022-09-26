Un bucle es un conjunto de ordenes que se repite. El bucle más sencillo que Go permite especificar es:
```go
for {
	//Instrucciones a repetir
}
```
EL bucle anterior es un “bucle infinito”. Si un programa llega a ese punto, el programa jamas continuara mas alla del bucle for, a no ser que el bucle se rompa con la instruccion break:
- Ver el codigo:
```sh
for_break.go
```
El ejemplo anterior repite un bucle en el que el programa pide un caracter al usuario, y se repite hasta que el usuario introduce el caracter ‘S’ o ‘s’, momento en el que el bucle se rompe mediante la instruccion break.

La instruccion continue rompe el flujo de cada iteracion de un for. En este caso, el flujo del programa salta de nuevo al inicio del bucle, sin ejecutar las instrucciones restantes de la iteracion en que se invoca:
- Ver el codgio:
```sh
for_continue.go
``` 
De acuerdo con el código del ejemplo:
 - Si el usuario introduce ‘N’ o ‘n’, se ejecutara la instruccion continue, por lo que la ejecuccion volvera al inicio del for.
 - Si el usuario introduce ‘S’ o ‘s’, se ejecutara la instruccion break, por lo que el for finaliza y la ejecuccion continua por el mensaje adios!
 - Solo en el caso de que el usuario introdujera cualquier otro caracter, se mostraria el mensaje caracter no reconocido.

 Las ordenes break y continue son totalmente validas y aceptadas en las convenciones sobre estilo de Go. Sin embargo, a menudo resulta mas limpia y util la forma condicional de for:
 ```go
 for <condicion>{
	//Instrucciones a ejecutar mientras 
	//la condicion sea cierta
}
 ```
 Por ejemplo: 
 - Ver codigo
 ```sh
 for_condicion.go
 ```

 - El programa anterior repetirá el código dentro del bucle mientras el usuario no introduzca 'S' ni 's'.

 Si usted esta familiarizado con otros lenguajes de programacion, se habra dado cuenta de que esta forma de for "condicional" suele llamarse while en otros lenguajes.

 Hay una tercera forma de describir un for; lo más similar al bucle for de los demas lenguajes:

 ```go
 for <inicio>; <condicion>; <actualizacion>{
    //Ordenes que se ejecutaran mientras la condicion
    //sea cierta
 }
 ```

 Este tipo de for: 
 1. Ejecuta la roden de inicio una sola vez, antes de ejecutar la primera iteración.
 2. Antes de cada iteracion, se comprobara si la condicion es cierta; en caso de que no lo sea, el bucle acabará.
 3. Despues de cada iteracion, antes de volver a comprobar si la condicion sigue siendo cierta, se ejecutará la orden de actualizacion. Por ejemplo, el siguiente bucle for mostrara una cuenta del 1 al 10:
 
 ```go
 for i := 1; i <= 10; i++{
    fmt.Println(i)
 }
 ``` 

 4. La variable i se inicia al valor 1 (observe que el alcance de esta se limita al bucle)
 5. Se muestra el valor de la variable i.
 6. Se incrementa la variable i.
 7. Se comprueba si el valor de i es menor o igual que 10. Si es el caso, se vuelve al paso 2. Si no es el caso, se sale del bucle for.