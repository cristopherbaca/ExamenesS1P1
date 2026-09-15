/*

 El algorito toma una temperatura Celcius y la convierte en Fahnereit y Kelvin, y los retorna.

*/

package main

import "fmt"

func main() {

	var c float64

	print(": ")
	fmt.Scanln(&c)

	f := (c * 9 / 5) + 32 // Conversion a Fahnereit
	k := c + 273.15       // Conversion a Kelvin

	fmt.Printf("Celsius: %f, Fahrenheit: %f, Kelvin: %f\n", c, f, k)
}
