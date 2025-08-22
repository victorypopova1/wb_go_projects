package main

import (
	"context"
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dataChan := make(chan string)
	var wg sync.WaitGroup

	fmt.Printf("Запуск %d воркеров...\n", numWorkers)
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, dataChan, &wg)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	counter := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				message := fmt.Sprintf("Сообщение #%d", counter)
				select {
				case dataChan <- message:
					counter++
					time.Sleep(500 * time.Millisecond)
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	<-signalChan
	fmt.Println("\nПолучен сигнал завершения, начинаем graceful shutdown...")

	cancel()

	close(dataChan)

	wg.Wait()
	fmt.Println("Все воркеры завершили работу. Программа завершена.")
}

func worker(ctx context.Context, id int, dataChan <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case data, ok := <-dataChan:
			if !ok {
				fmt.Printf("Воркер %d: канал закрыт, завершаю работу\n", id)
				return
			}
			fmt.Printf("[Воркер %d] %s\n", id, data)

		case <-ctx.Done():
			fmt.Printf("Воркер %d: получен сигнал завершения\n", id)
			return
		}
	}
}
