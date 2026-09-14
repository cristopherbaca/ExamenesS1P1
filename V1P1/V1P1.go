package V1P1

import "fmt"

func main() {
	var c float64

	fmt.Print("Enter the temperature in Celsius: ")
	fmt.Scanln(&c)

	f := (c * 9 / 5) + 32
	k := c + 273.15

	fmt.Printf("Celsius: %.2f, Fahrenheit: %.2f, Kelvin: %.2f\n", c, f, k)
}