El operador ampersand(&) delante de una variable retorna la direccion de memoria de esta. Este valor se puede asignar directamente a un puntero: 
```go
i := 10
var p *int
p = &i
```

En el ejemplo anterior, el apuntador p apuntará a la variable i. El codgio anterior se puede abreviar de la siguiente manera:
```go
i := 10
p := &i
```
Ya que &i retorna un valor del tipo *int, p será declarado como *int(apuntador a int) y desde el principio apuntará al int i.

img:
