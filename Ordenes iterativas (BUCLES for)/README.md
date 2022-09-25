Un bucle es un conjunto de ordenes que se repite. El bucle más sencillo que Go permite especificar es:
```ss
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

La instruccion continue rompe el flujo de cada iteracion de un for. En este caso, el flujo del programa salta de nuevo al inicio del bucle, sin ejecutar las instrucciones restantes de la iteracion en que se invoca

 -Ver el codgio:
```sh
for_continue.go
``` 
De acuerdo con el código del ejemplo:
 - Si el usuario introduce ‘N’ o ‘n’, se ejecutara la instruccion continue, por lo que la ejecuccion volvera al inicio del for.
 - Si el usuario introduce ‘S’ o ‘s’, se ejecutara la instruccion break, por lo que el for finaliza y la ejecuccion continua por el mensaje adios!
 - Solo en el caso de que el usuario introdujera cualquier otro caracter, se mostraria el mensaje caracter no reconocido.