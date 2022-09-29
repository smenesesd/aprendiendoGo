Las variables de Go, por defecto, son operadas "por valor". Esto significa que:
- El operador de asignanción =``` copia ```el valor de la derecha hacia el valor de la variable de la izquierda. Tras las intrucción a = b, a y b serán dos variables con el mismo valor, pero no son la misma variable: ocupan zonas distintas de la memoria y la modificación de una no afecta al valor anterior de la otra.

- El operador de igualdad == entre dos variables resulta en ```true``` si ambas variables tienen un valor igual, aunque sean dos variables distintas.

Los punteros de Go nos permitirán operar "por referencia", es decir: 
- El operador de asignación = apunta (o referencia) el apuntador de la izquierda hacia el apuntador de la derecha. Tras la instruccion p = &i, el apuntador p apunta hacia la misma direccion de memoria donde esta i, Cualquier modificacion de p afectara a tal valor y a cómo se accede a traves de i. 
- El operadro de igualdad == entre dos apuntadores resulta en true si ambos apuntan a la misma direccion de memoria, o en false si apuntan a direcciones distintas (aunque esas direcciones contengan el mismo valor). Si se quisiera comparar la igualdad de dos valores apuntados por los apuntadores p1 y p2, se deberia usar la operador asterisco(*) para obtener el valor de ambos y, asi, poder comparar valores en vez de direcciones:
- Ver codigo:
```valoresVsReferencias.go```

El ejemplo anterior mostraria la siguiente salida estándar:

```p1 y p2 apuntan a direcciones distintas```

```p1 y p2 apuntan a valores iguales```

Ya que p1 y p2 apuntan a dos variables con el mismo valor, pero en distintas zonas de memoria (figura 4.5), tan solo mostraria el mensaje p1 y p2 apuntan a valores iguales, ya que apuntan a dos variables iguales, pero que no son la misma.

![Figura 4.5](https://raw.githubusercontent.com/smenesesd/aprendiendoGo/master/img/Figura4.5.png)
```Figura4.5``` p1 y p2 apuntan a valores iguales, pero distintas variables.

La primera comprobacion seria "por referencia", ya que se comparan direcciones de memoria; mientras que la segunda comprobacion seria "por valor", ya que se comparan dos valores concretos.