La entrada estandar permite a un programa en Go obtener datos desde el exterior, a traves del teclado en la linea de comandos.
Go tiene varios modos para entrada de datos de parte del usuario. 
### Funciones:
- Scan(): Lee los datos del teclado y los guarda en las variables pasadas en la invocacion. Cada variable debe ir precedida por el simbolo ampersand "&" 
- Ver codigo:
```go
EntradaDeDatos.go
```
- Scanln(): Es similar al comando de Scan(), pero con este se para se escanear cuando el usuario hunde enter.
- Scanf(): Permite especificar con mas detalle el formato de entrada, tomando como primer parametro una cadena de texto en la que se pueden introducir los diversos verbos.
- Ver codigo:
```go
- EntradaDeDatosScanf.go
- EntradaDatosScanf.go
```