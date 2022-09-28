La manera de definir una variable, dado un nombre y un tipo de dato, es:
``` var <nombre> <tipo> [= <valor>]```

Por ejemplo:
```go
var dias int
var meses int = 12
```
Si no se provee un valor inicial, las variables serán inicializadas automáticamente con el “valor cero” de cada tipo de datos: 0 para tipos numéricos, false para bool, y la cadena vacía ”” para string.
- En aras de la brevedad, se puede omitir tanto la palabra var como su tipo si se usa el operador de inicializacion ":="

``` meses := 12 ``` 
### Importante:
 Es recomendable declarar variables  con el operador := siempre que se pueda, aunque las variables globales y las que no tengan valor inicial deberán declararse mediante var. 