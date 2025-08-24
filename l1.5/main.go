package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Использование: go run main.go <N_секунд>")
		return
	}

	seconds, err := strconv.Atoi(os.Args[1])
	if err != nil || seconds <= 0 {
		fmt.Println("Ошибка: количество секунд должно быть положительным числом")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(seconds)*time.Second)
	defer cancel()

	dataChan := make(chan int)

	go func() {
		for {
			select {
			case value, ok := <-dataChan:
				if !ok {
					fmt.Println("Читатель: канал закрыт")
					return
				}
				fmt.Printf("← Прочитано: %d\n", value)
			case <-ctx.Done():
				fmt.Println("Читатель: контекст завершен")
				return
			}
		}
	}()

	fmt.Printf("Отправка данных в течение %d секунд...\n", seconds)

	counter := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("\nВремя вышло, завершаем отправку...")
			close(dataChan)
			return
		default:
			select {
			case dataChan <- counter:
				fmt.Printf("→ Отправлено: %d\n", counter)
				counter++
				time.Sleep(400 * time.Millisecond)
			case <-ctx.Done():
				continue
			}
		}
	}
}
