package main

import "fmt"

func FeetToMeters(feet float64) float64 {
	return feet * 0.3048
}

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := FeetToMeters(input)
	fmt.Println(output)
}
