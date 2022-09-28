Una constante es un valor literal al que se le asigna un nombre, que no puede cambiar durante la vida del programa. Su definicion es similar a la de una variable, pero remplazando la palabra var por la palabra const. Por ejemplo:
```go
const Pi = 3.1416
```
ó
```go
const Pi float64 = 3.1416 
```
Cuando se definen multiples constantes, se puede agrupar semanticamente bajo la misma directiva const. Por ejemplo:
```go
//configuracion de una tipografia 
const (
    TipoFuente = "Times New Roman"
    AlturaFuente = 12
    Subrayado = false
    Negrita = false
)
```