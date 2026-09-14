package main

import "fmt"

func main() {
	n := 97
	primo := true

	if n < 2 {
		primo = false
	} else {
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				primo = false
				break
			}
		}
	}

	fmt.Println(primo)
}