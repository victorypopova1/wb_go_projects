package main

import (
	"fmt"
	"sync"
)

func generateNumbers(numbers chan<- int, data []int) {
	defer close(numbers)
	for _, x := range data {
		numbers <- x
	}
}

func doubleNumbers(numbers <-chan int, doubled chan<- int) {
	defer close(doubled)
	for x := range numbers {
		doubled <- x * 2
	}
}

func printResults(doubled <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range doubled {
		fmt.Println(result)
	}
}

func main() {
	numbers := make(chan int, 3)
	doubled := make(chan int, 3)

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var wg sync.WaitGroup

	go generateNumbers(numbers, data)

	go doubleNumbers(numbers, doubled)

	wg.Add(1)
	go printResults(doubled, &wg)

	wg.Wait()

	fmt.Println("Конвейер завершил работу")
}
