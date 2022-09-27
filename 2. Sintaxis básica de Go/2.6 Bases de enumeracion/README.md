Ademas de lo clásico que es del 0 al 9, Go permite usar otros tipos de bases de numeración según el prefijo que añadamos a cada numero.
Base | Prefijo | Descripcion | Ejemplo
---  | ---     | ---         | ---
Binario | 0b | Cada digito puede ser 0 o 1 | 0b10001101
Octal | 0 | Cada digito puede tomar valores del 0 al 7 | 012072
Hexadecimal | 0x | Cada digito puede tomar valores del 0 al 9, y de la A a la F representando los valores del 10 al 15 | 0xAF32

Cuando queremos agrupar bloques de dígitos, podemos insertar un guion bajo “_” entre estos. Go no lo tendrá en cuenta a la hora de establecer el valor. Ejemplo:
```go
productoInternoBruto := 1_419_000_000_000
bitsAgrupados := 0b_1000_1001_0110
```
- Ver codigo en: 

``` basesDeEnumeracion.go ```


---
 esto me da una respuesta en la terminal:

``` 1429000000000 2198```