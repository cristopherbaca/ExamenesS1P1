package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 3, 4, 4, 5, 5, 5, 7}

	output := []int{}

	for i := 0; i < len(arr); i++ {
		encontrado := false

		for j := 0; j < len(output); j++ {
			if arr[i] == output[j] {
				encontrado = true
			}
		}

		if !encontrado {
			output = append(output, arr[i])
		}
	}

	for _, x := range output {
		fmt.Println(x)
	}
}
