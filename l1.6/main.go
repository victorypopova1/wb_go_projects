package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	fmt.Println("1. Завершение по условию:")
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Горутина запущена")
		counter := 0
		for counter < 3 {
			fmt.Printf("Работает, счетчик: %d\n", counter)
			counter++
			time.Sleep(300 * time.Millisecond)
		}
		fmt.Println("Завершена по условию")
	}()
	wg.Wait()

	fmt.Println("\n2. Завершение через канал уведомления:")
	done := make(chan bool)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Горутина запущена")
		for {
			select {
			case <-done:
				fmt.Println("Получен сигнал завершения")
				return
			default:
				fmt.Println("Работает...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()
	time.Sleep(1 * time.Second)
	done <- true
	wg.Wait()

	fmt.Println("\n3. Завершение через закрытие канала:")
	stopChan := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Горутина запущена")
		for {
			select {
			case _, ok := <-stopChan:
				if !ok {
					fmt.Println("Канал закрыт, завершаюсь")
					return
				}
			default:
				fmt.Println("   Работает...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()
	time.Sleep(1 * time.Second)
	close(stopChan)
	wg.Wait()

	fmt.Println("\n4. Завершение через контекст (cancel):")
	ctx1, cancel1 := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Горутина запущена")
		for {
			select {
			case <-ctx1.Done():
				fmt.Println("Контекст отменен")
				return
			default:
				fmt.Println("Работает...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()
	time.Sleep(1 * time.Second)
	cancel1()
	wg.Wait()

	fmt.Println("\n5. Завершение через контекст с таймаутом:")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel2()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Горутина запущена")
		for {
			select {
			case <-ctx2.Done():
				fmt.Println("Время контекста истекло")
				return
			default:
				fmt.Println("Работает...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()
	wg.Wait()

	fmt.Println("\n6. Завершение через контекст с дедлайном:")
	deadline := time.Now().Add(1 * time.Second)
	ctx3, cancel3 := context.WithDeadline(context.Background(), deadline)
	defer cancel3()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Горутина запущена")
		for {
			select {
			case <-ctx3.Done():
				fmt.Println("Дедлайн достигнут")
				return
			default:
				fmt.Println("Работает...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()
	wg.Wait()

	fmt.Println("\n7. Завершение через runtime.Goexit():")
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Горутина запущена")

		go func() {
			fmt.Println("Внутренняя горутина запущена")
			time.Sleep(300 * time.Millisecond)
			fmt.Println("Внутренняя горутина вызывает Goexit()")
			runtime.Goexit()
			fmt.Println("Этот код не выполнится")
		}()

		time.Sleep(1 * time.Second)
		fmt.Println("Основная горутина завершена")
	}()
	wg.Wait()

	fmt.Println("\n8. Завершение через панику (с recovery):")
	wg.Add(1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Восстановлено после паники: %v\n", r)
			}
			wg.Done()
		}()
		fmt.Println("Горутина запущена")
		time.Sleep(300 * time.Millisecond)
		fmt.Println("Вызываю панику")
		panic("экстренное завершение")
	}()
	wg.Wait()

	fmt.Println("\n9. Естественное завершение (return):")
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Горутина запущена")
		for i := 0; i < 3; i++ {
			fmt.Printf("Итерация %d\n", i)
			time.Sleep(300 * time.Millisecond)
		}
		fmt.Println("Естественное завершение")
		return
	}()
	wg.Wait()

	fmt.Println("\n10. Комбинированный способ (контекст + канал):")
	ctx4, cancel4 := context.WithCancel(context.Background())
	combinedDone := make(chan bool)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Горутина запущена")
		for {
			select {
			case <-ctx4.Done():
				fmt.Println("Контекст отменен")
				return
			case <-combinedDone:
				fmt.Println("Сигнал из канала")
				return
			default:
				fmt.Println("Работает...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()
	time.Sleep(1 * time.Second)
	combinedDone <- true
	cancel4()
	wg.Wait()

	fmt.Println("\n11. Завершение через таймер:")
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Горутина запущена")
		timer := time.NewTimer(1 * time.Second)
		defer timer.Stop()

		for {
			select {
			case <-timer.C:
				fmt.Println("Таймер сработал")
				return
			default:
				fmt.Println("Работает...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()
	wg.Wait()
}
