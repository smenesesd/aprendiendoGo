A través de un apuntador, se puede leer o modificar el valor de la varibale apuntada. El operador asterisco * delante del nombre del apuntador nos permitirá acceder al valor de la variable apuntada, tanto para su lectura como para su modificacion:
```go
i := 10
p := &i
a := *p
*p = 21
```
En el ejemplo anterior, donde el apuntador p apunta a la variable i, operando a través de este apuntador se está leyendo y modificando el valor de la variable i. La tercera línea (a := *p) tendria un resultado equivalente a a := i, y la cuarta linea (*p = 21) tendira un resultado equivalente a i = 21. la figura 4.4 esquematiza la evolucion de la memoria según las anteriores instrucciones se van ejecutando: qué las variables se van creando y qué valores toman en cada momento.

![Figura 4.4](https://raw.githubusercontent.com/smenesesd/aprendiendoGo/master/img/Figura4.4.png)
``` Figura 4.4:``` Uso del operador asterisco sobre punteros

Viendolo desde el punto de vista del ejemplo anterior, podria parecer que los apuntadores son un paso intermedio innecesario. Más adelante, se demostrara la verdadera utilidad de los apuntadores cuando se usan para recorrer complejas estructuras de datos, para referirse a alguna de sus partes o para intercambiar datos con "Funciones".