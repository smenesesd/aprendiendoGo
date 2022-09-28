Go permite definir cadenas de texto explicitamente, insertando un texto cualquiera entre comillas dobles:
```texto : "- Hola, Como estas?"```

Si una cadena ha de contener un texto mostrado en diversas lineas, puede introducir un el caracter especial de nueva linea \n alla donde quiera que termine una linea y empiece otra:
```texto := "Hola, como estas \n - Estoy bien, Gracias"  ```

Si una cadena ha de contener comillas dobles en su interior, estas deben especificarse como un caracter especial \ “, para que go no las confunda como el final de una cadena de texto:
```texto := "Podria decirse que estoy \"bien\"..." ```

Cuando un texto contiene multiples lineas o comillas dobles, puede resultar mas limpio substituir el delimitador de comillas dobles  por el de “acento grave”; esto le permitira escribir cadenas en multiples lineas, tomando los saltos de linea como literales. El equivalente a la cadena anterior, seria: 
```go 
texto := Hola, ¿Como estas?
- Estoy bien", gracias 
 ```