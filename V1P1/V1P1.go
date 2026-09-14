package main

import "fmt"

func main(){

	c := 0;

	fmt.Printf("Enter the temperature: ");
	fmt.Scanln(&c);

	f := (c * (9/5)) + 32;
	k := float64(c) + 275.13;

	fmt.Println("Celsius: ", c , ", Fahnreit: " , f , ", Kelvin: " , k) ;

}