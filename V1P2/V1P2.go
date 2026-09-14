package main

import "fmt"

func main() {
	n := 97
	primo := true

	if n < 2 {
		primo = false
	} else {
		for i := 2; i*i <= n-1; i++ { // Podemos poner i*i <= n-1 pero seria lo mismo que i*i < n
			if n%i == 0 {
				primo = false
				break
			}
		}
	}

	fmt.Println(primo)
}