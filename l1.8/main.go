package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMap struct {
	mu   sync.Mutex
	data map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	val, exists := sm.data[key]
	return val, exists
}

func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}

func (sm *SafeMap) Len() int {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return len(sm.data)
}

type SafeMapRW struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewSafeMapRW() *SafeMapRW {
	return &SafeMapRW{
		data: make(map[string]int),
	}
}

func (sm *SafeMapRW) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMapRW) Get(key string) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	val, exists := sm.data[key]
	return val, exists
}

func (sm *SafeMapRW) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}

func syncMapExample() {
	var sm sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key_%d", id)
			sm.Store(key, id*100)
			fmt.Printf("Записано: %s -> %d\n", key, id*100)
		}(i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key_%d", id)
			if value, ok := sm.Load(key); ok {
				fmt.Printf("Прочитано: %s -> %v\n", key, value)
			}
		}(i)
	}

	wg.Wait()
}

func main() {
	fmt.Println("=== Безопасная работа с map ===")
	fmt.Println()

	fmt.Println("1. Тестирование SafeMap с sync.Mutex:")
	safeMap := NewSafeMap()
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key_%d", id%5)
			safeMap.Set(key, id)
			time.Sleep(10 * time.Millisecond)
		}(i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key_%d", id%5)
			if val, exists := safeMap.Get(key); exists {
				fmt.Printf("  Получено: %s -> %d\n", key, val)
			}
			time.Sleep(15 * time.Millisecond)
		}(i)
	}

	wg.Wait()
	fmt.Printf("  Размер map: %d\n", safeMap.Len())

	fmt.Println("\n2. Тестирование SafeMap с sync.RWMutex:")
	safeMapRW := NewSafeMapRW()

	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("rw_key_%d", id%3)
			safeMapRW.Set(key, id*2)
		}(i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("rw_key_%d", id%3)
			if val, exists := safeMapRW.Get(key); exists {
				fmt.Printf("  Прочитано RWMutex: %s -> %d\n", key, val)
			}
		}(i)
	}

	wg.Wait()

	fmt.Println("\n3. Тестирование sync.Map:")
	syncMapExample()

	fmt.Println("\n4. Демонстрация UNSAFE подхода (гонки данных!):")
	unsafeMap := make(map[string]int)
	var unsafeWg sync.WaitGroup

	for i := 0; i < 5; i++ {
		unsafeWg.Add(1)
		go func(id int) {
			defer unsafeWg.Done()
			unsafeMap[fmt.Sprintf("unsafe_%d", id)] = id
		}(i)
	}

	unsafeWg.Wait()
	fmt.Println("Небезопасная запись завершена (возможны гонки)")

	fmt.Println("\n=== Тестирование завершено ===")
	fmt.Println("Запустите с флагом -race для проверки: go run -race main.go")
}

func (sm *SafeMap) Increment(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key]++
}

func (sm *SafeMap) GetAll() map[string]int {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	result := make(map[string]int)
	for k, v := range sm.data {
		result[k] = v
	}
	return result
}

func benchmarkMaps() {
	fmt.Println("\n5. Бенчмарк разных подходов:")

	// Тест с обычным map + mutex
	sm := NewSafeMap()
	start := time.Now()
	for i := 0; i < 1000; i++ {
		sm.Set(fmt.Sprintf("test_%d", i), i)
	}
	fmt.Printf("  SafeMap (Mutex): %v\n", time.Since(start))

	// Тест с RWMutex
	smRW := NewSafeMapRW()
	start = time.Now()
	for i := 0; i < 1000; i++ {
		smRW.Set(fmt.Sprintf("test_%d", i), i)
	}
	fmt.Printf("  SafeMap (RWMutex): %v\n", time.Since(start))

	// Тест с sync.Map
	var smSync sync.Map
	start = time.Now()
	for i := 0; i < 1000; i++ {
		smSync.Store(fmt.Sprintf("test_%d", i), i)
	}
	fmt.Printf("  sync.Map: %v\n", time.Since(start))
}
