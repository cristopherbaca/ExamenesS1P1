package main

import "fmt"

func main() {

	segundos := 9385

	horas := segundos / 3600
	resto := segundos % 3600

	minutos := resto / 60
	segundosRestantes := resto % 60

	fmt.Printf("Horas: %d, Minutos: %d, Segundos: %d\n", horas, minutos, segundosRestantes)
}
