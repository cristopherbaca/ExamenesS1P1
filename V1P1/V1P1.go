package main

import "fmt"

func main() {
	var c float64

	fmt.Print("Enter the temperature in Celsius: ")
	fmt.Scanln(&c)

	f := (c * 9 / 5) + 32
	k := c + 273.15

	fmt.Printf("Celsius: %f, Fahrenheit: %f, Kelvin: %f\n", c, f, k)
}