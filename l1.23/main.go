package main

import (
	"fmt"
	"runtime"
)

func removeElementGeneric[T any](slice []T, index int) []T {
	if index >= len(slice) {
		return slice
	}

	copy(slice[index:], slice[index+1:])

	var zero T
	slice[len(slice)-1] = zero

	return slice[:len(slice)-1]
}

func main() {
	var m runtime.MemStats

	runtime.GC()

	numbers := []int{100, 200, 300}
	fmt.Println("Исходный:", numbers)

	runtime.ReadMemStats(&m)
	fmt.Printf("Используется %d байт\n", m.Alloc)

	numbers = removeElementGeneric(numbers, 1)
	fmt.Println("После удаления:", numbers)

	runtime.ReadMemStats(&m)
	fmt.Printf("Используется %d байт\n", m.Alloc)

	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Printf("После GC: %d байт\n", m.Alloc)

}
