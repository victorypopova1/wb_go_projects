package main

import "fmt"

func Intersection[T comparable](a, b []T) []T {
	set := make(map[T]bool)
	for _, item := range a {
		set[item] = true
	}

	var result []T
	for _, item := range b {
		if set[item] {
			result = append(result, item)
			set[item] = false // предотвращаем дубликаты
		}
	}

	return result
}

func main() {
	nums1 := []int{1, 2, 3, 4, 2}
	nums2 := []int{2, 3, 5, 6}
	fmt.Printf("Числа: %v\n", Intersection(nums1, nums2))

	strings1 := []string{"apple", "banana", "orange"}
	strings2 := []string{"banana", "grape", "orange"}
	fmt.Printf("Строки: %v\n", Intersection(strings1, strings2))
}
