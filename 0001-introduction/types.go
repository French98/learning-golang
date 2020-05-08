package main

import "fmt"

func add(x, y float64) float64 {
	return x + y
}

func main() {
	//go will figure out the type
	num1, num2 := 7.2, 9.35

	fmt.Println(add(num1, num2))
}
