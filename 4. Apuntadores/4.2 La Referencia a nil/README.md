Los apuntadores definidos en la seccion anterior no han sido inicializados a ningun valor. Apuntan al valor ```nil```, que podria interpretarse como "a ninguna parte".

El valor ```nil``` se puede utilizar tanto para reiniciar el valor de un apuntador "a nada" como para comprobar si un apuntador apunta a algún lugar valido:
```go
var pi *int
pi = nil
if pi == nil{
    fmt.Print("¡No puedo hacer nada con este apuntador!")
    fmt.Println("porque no apunta a nada!")
}
```
- Ver codigo:
```referenciaNil.go``` 

Cuando un apuntador hace referencia a la direccion nil, este no se puede utilizar para leer o modificar valores apuntados, ya que dicho valor apuntado no existe. Si tratáramos de hacerlo, el sistema de memoria segura de Go abortaria la ejecución del programa, mostrando un error similar al siguiente:

```panic: runtime error: invalid memory address or nil pointer dereference```

```[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x49169f] ```