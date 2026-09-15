package main

import "fmt"

func main() {
    a := 252
    b := 105

    for b != 0 {
        resto := a % b
        a = b
        b = resto
    }

    fmt.Println("MCD:", a)
}

