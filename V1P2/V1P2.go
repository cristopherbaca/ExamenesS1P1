/*

	El algoritmo toma un numero(n) y le hace una comparacion-verificacion-divisibilidad de 2 a N-1 para checar si
	ese numero es divisible entre el rango 2 a N-1, si lo es entonces el numero no es primo.

*/

package main

func main() {

	n := 97
	primo := true
 
	if n < 2 { // Si el numero es menor a 2 pues no es primo ya que no contaria como tal
		primo = false
	} else {
		for i := 2; i*i <= n-1; i++ { // Compilamos de 2 a n-1 o < n, probando cada posible opcion de numero entero
			// Podemos poner i*i <= n-1 pero seria lo mismo que i*i < n
			if n%i == 0 { // Si en un caso llega a ser divisible eso quiere decir que no es primo
				primo = false
				break
			}
		}
	}

	println(primo)
}
