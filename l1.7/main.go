package main

import (
	"fmt"
)

func SetBit(num int64, i uint, bit int) int64 {
	if bit == 1 {
		// в 1 с помощью OR
		return num | (1 << i)
	} else {
		// в 0 с помощью AND NOT
		return num &^ (1 << i)
	}
}

func main() {
	var num int64 = 5 // 0101
	var i uint = 1    // 1-й бит

	// 1-й бит в 0
	result := SetBit(num, i, 0)
	fmt.Printf("Исходное число: %d (%b)\n", num, num)
	fmt.Printf("После установки %d-го бита в 0: %d (%b)\n", i, result, result)

	// 1-й бит обратно в 1
	result = SetBit(result, i, 1)
	fmt.Printf("После установки %d-го бита в 1: %d (%b)\n", i, result, result)

	// установка 0-го бита в 0 для числа 7 (0111)
	num = 7
	result = SetBit(num, 0, 0)
	fmt.Printf("%d (%b) -> установка 0-го бита в 0 -> %d (%b)\n", num, num, result, result)

	// установка 2-го бита в 1 для числа 3 (0011)
	num = 3
	result = SetBit(num, 2, 1)
	fmt.Printf("%d (%b) -> установка 2-го бита в 1 -> %d (%b)\n", num, num, result, result)
}
