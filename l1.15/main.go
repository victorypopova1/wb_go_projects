package main

import (
	"fmt"
	"strings"
)

// имитация функции создания большой строки
func createHugeString(size int) string {
	return strings.Repeat("x", size)
}

// корректная реализация
func someFuncCorrect() string {
	v := createHugeString(1 << 10) // 1024 символа

	// создаем новую строку с копированием только нужных данных
	result := make([]byte, 100)
	copy(result, v[:100])

	return string(result)
}

func main() {
	justString := someFuncCorrect()
	fmt.Printf("Длина: %d\n", len(justString))
	fmt.Printf("Первые 10 символов: %s\n", justString[:10])
}
