Las variables de Go, por defecto, son operadas "por valor". Esto significa que:
- El operador de asignanción =``` copia ```el valor de la derecha hacia el valor de la variable de la izquierda. Tras las intrucción a = b, a y b serán dos variables con el mismo valor, pero no son la misma variable: ocupan zonas distintas de la memoria y la modificación de una no afecta al valor anterior de la otra.

- El operador de igualdad == entre dos variables resulta en ```true``` si ambas variables tienen un valor igual, aunque sean dos variables distintas.

