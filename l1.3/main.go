package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Использование: go run main.go <количество_воркеров>")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Ошибка: количество воркеров должно быть положительным числом")
		return
	}

	dataChan := make(chan string)
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup

	fmt.Printf("Запуск %d воркеров...\n", numWorkers)
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChan, &wg)
	}

	counter := 1
	go func() {
		for {
			select {
			case <-stopChan:
				fmt.Println("\nПолучен сигнал остановки, закрываю канал...")
				close(dataChan)
				return
			default:
				message := fmt.Sprintf("Сообщение #%d", counter)
				dataChan <- message
				counter++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
	fmt.Println("Все воркеры завершили работу")
}

func worker(id int, dataChan <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range dataChan {
		fmt.Printf("[Воркер %d] %s\n", id, data)
	}
	fmt.Printf("Воркер %d завершил работу\n", id)
}
