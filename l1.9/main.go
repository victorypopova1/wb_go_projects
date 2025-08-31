package main

import (
	"fmt"
	"sync"
)

// generateNumbers генерирует числа и отправляет их в канал
func generateNumbers(numbers chan<- int, data []int) {
	defer close(numbers)
	for _, x := range data {
		numbers <- x
	}
}

// doubleNumbers умножает числа на 2 и отправляет результат в другой канал
func doubleNumbers(numbers <-chan int, doubled chan<- int) {
	defer close(doubled)
	for x := range numbers {
		doubled <- x * 2
	}
}

// printResults выводит результаты в stdout
func printResults(doubled <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range doubled {
		fmt.Println(result)
	}
}

func main() {
	// Создаем каналы
	numbers := make(chan int, 3) // Буферизованный канал для лучшей производительности
	doubled := make(chan int, 3)

	// Массив чисел для обработки
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var wg sync.WaitGroup

	// Запускаем генерацию чисел
	go generateNumbers(numbers, data)

	// Запускаем обработку чисел
	go doubleNumbers(numbers, doubled)

	// Запускаем вывод результатов
	wg.Add(1)
	go printResults(doubled, &wg)

	// Ожидаем завершения вывода
	wg.Wait()

	fmt.Println("Конвейер завершил работу")
}
