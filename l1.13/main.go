package main

import "fmt"

func main() {
	a, b := 5, 10
	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	a = a ^ b // a = 15 (5 ^ 10), b = 10
	b = a ^ b // a = 15, b = 5 (15 ^ 10)
	a = a ^ b // a = 10 (15 ^ 5), b = 5

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)
}
