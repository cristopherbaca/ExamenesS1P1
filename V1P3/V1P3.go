package main

func main() {

	a := [8]int{7,8,7,9,7,8,10}


	mejorValor := a[1]
	mejorFrecuencia := 0

	for i := 0; i < len(a); i++ {
		frecuencia := 0
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] {
				frecuencia++;
			}
		}
	}

	println(true)

}
