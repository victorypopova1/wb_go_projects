package main

import "fmt"

func createSet(items []string) map[string]bool {
	set := make(map[string]bool)
	for _, item := range items {
		set[item] = true
	}
	return set
}

func setToSlice(set map[string]bool) []string {
	result := make([]string, 0, len(set))
	for item := range set {
		result = append(result, item)
	}
	return result
}

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	set := createSet(words)
	uniqueWords := setToSlice(set)

	fmt.Printf("Множество: %v\n", uniqueWords)
}
