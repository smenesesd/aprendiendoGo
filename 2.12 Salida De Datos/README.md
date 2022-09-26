El paquete fmt, permite mostrar datos en su terminal de linea de comandos(tambien llamado salida estandar) mediante las siguientes funciones:
```go
fmt.Print
fmt.Println
fmt.Printf
```

```fmt.Print y fmt.Println``` enviara a la salida estandar los datos que situe en parentesis, y separados por comas.

- Para los datos que no sean cadenas de texto, Go hara una conversion generica a texto antes de enviarlos a la salida estandar. Ademas, añadira un espacio entre los diferentes datos dentro de una misma invocacion a ```fmt.Print y fmt.Println```