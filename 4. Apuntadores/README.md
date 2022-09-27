### Definicion de APUNTADORES:
- La memoria de un ordenador podria abstraerse, como si fuera un conjunto de cajones colocados en una gigantesca estanteria. Cada cajon tiene un numero identificativo unico, que permite al ordenador referirse a el, a la hora de ir a buscar o guardar datos. Dicho número identificativo se conoce como "Direccion de memoria".
- Los apuntadores (o punteros) son variables que no guardan valores útiles como tales, sino las direcciones de memoria donde se encuentras dichos valores. Su nombre hace referencia al hecho de que no guardan una variable sino que "apuntan" a su direccion.  
- En go, los apuntadores tienen un tipo asociado, es decir, solo pueden apuntar a variables de un tipo en concreto: un puntero int solo podrá guardar la direccion de memoria de una variable del tipo int, y asi entre los demas. 