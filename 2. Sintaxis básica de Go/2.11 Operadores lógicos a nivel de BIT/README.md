Las variables numéricas tambien pueden verse como agrupaciones de dígitos binarios (bits) con los que podemos realizar las siguientes operaciones, que se aplican tras comparar los bits de igual posición en dos operandos.
Símbolo | Descripción | Ejemplo
---     | ---         | ---
&       | AND. Resulta 1 cuando si - y solo si- ambos bits son 1              | 0b_1010 & 0b_1100 == 0b_1000
Or      | OR. Resulta 1 cuando al menos uno de los bits es 1            | 0b_1010 Or 0b_1100 == 0b_1110 
^       | XOR (o exclusivo). Resulta 1 si los bits comparados son diferentes, o 0 si son iguales |0b_1010^0b_1100==0b_0110     
^(unario)| NOT. Cuando se aplica delante de un solo operador, invierte el valor de cada dígito |^0b_100110 == 0b_011001         
<< | Desplazamiento a la izquierda. Desplaza cada bit del operador izquierdo tantas posiciones como el numero del operador derecho. Los espacios libres a la derecha se rellenan de ceros | 0b_1001 << 2 == 0b_100100
">>" | Desplazamiento a la derecha. Desplaza cada bit del operador izquierdo tantas posiciones como el numero del operador derecho. Los bits mas a la derecha son eliminados | 0b_1001 >> 2 == 0b_0010  

Or: || 
      
