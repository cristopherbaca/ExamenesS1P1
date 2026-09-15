/*
	El algoritmo recorre el arreglo, buscando X numero, cuando lo encuentra le aumenta uno a su frecuencia, una vez
	terminada ese frecuencia la compara con la anterior para ver si hay otra mayor y remplazarla.
*/

package main

func main() {

	a := [7]int{7, 8, 7, 9, 7, 8, 10} // Creamos el arreglo

	mejorValor := a[0]
	mejorFrecuencia := 0

	for i := 0; i < len(a); i++ { // Recorremos el arreglo
		frecuencia := 0

		for j := 0; j < len(a); j++ { // Buscamos cuantas veces aparece a[i]
			if a[i] == a[j] { // Si son iguales, contamos la aparición
				frecuencia++
			}
		}

		if frecuencia > mejorFrecuencia { // Si encontramos una frecuencia mayor
			mejorFrecuencia = frecuencia
			mejorValor = a[i]
		}
	}

	println("Mejor valor:", mejorValor)
	println("Frecuencia:", mejorFrecuencia)
}
