package main

import "os"

func main() {

	n := 97


	if n < 2 {
		println(false)
		os.Exit(0)
	}



	for i := 2; i <= n-1; i++ {
		if n%i == 0 {
			println(false)
			os.Exit(0)
		}
	}

	println(true)

}
