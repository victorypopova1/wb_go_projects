package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// выбираем опорный элемент (середина)
	pivot := arr[len(arr)/2]

	var left, right, equal []int

	// разделяем элементы относительно опорного
	for _, num := range arr {
		switch {
		case num < pivot:
			left = append(left, num)
		case num == pivot:
			equal = append(equal, num)
		case num > pivot:
			right = append(right, num)
		}
	}

	// рекурсивно сортируем левую и правую части
	left = quickSort(left)
	right = quickSort(right)

	// объединяем результаты
	return append(append(left, equal...), right...)
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90, 88}
	fmt.Println("Исходный массив:", arr)

	sorted := quickSort(arr)
	fmt.Println("Отсортированный массив:", sorted)
}
