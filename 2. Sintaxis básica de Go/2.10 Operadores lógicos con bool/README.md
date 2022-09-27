Cualquier variable o expresion que retorne bool puede combinarse con otros booleanos mediante los siguientes operadores logicos, resultando en otros valores booleanos. Cuando se mezclan con las operaciones de comparacion del apartado anterior, las operaciones de logica son las de menor precedencia.
Operador | Descripcion | Ejemplo
---      |---          |---
&&       | AND. Retorna true si -y solo si- ambos operadores booleanos son true | 6 == 6 && 3 > 2 (Resultado = true && true, entonces true)
OR | OR. Retorno true en caso de que uno de los operadores booleanos sea true |  false OR 3 < 8 (Resultado: false OR true, entonces true) 
!(unario) | NOT. Invierte true por false y false por true en el operador situado a la derecha | !(7 < 1)(Resultado: !false, entonces true)

OR: || 