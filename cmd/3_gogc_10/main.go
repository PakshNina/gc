package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"runtime/trace"
	"sync"
	"time"
)

const (
	NumWorkers    = 4     // Количество воркеров
	NumTasks      = 500   // Количество задач
	MemoryIntense = 10000 // Размер память затратной задачи (число элементов)
)

func main() {
	// Запись в trace файл
	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()

	// Установка целевого процента сборщика мусора GOGC=10%
	debug.SetGCPercent(10)

	// Очередь задач и очередь результата
	taskQueue := make(chan int, NumTasks)
	resultQueue := make(chan int, NumTasks)

	// Запуск воркеров
	var wg sync.WaitGroup
	wg.Add(NumWorkers)
	for i := 0; i < NumWorkers; i++ {
		go worker(taskQueue, resultQueue, &wg)
	}

	// Отправка задач в очередь
	for i := 0; i < NumTasks; i++ {
		taskQueue <- i
	}
	close(taskQueue)

	// Получение результатов из очереди
	go func() {
		wg.Wait()
		close(resultQueue)
	}()

	// Обработка результатов
	for result := range resultQueue {
		fmt.Println("Результат:", result)
	}

	fmt.Println("Готово!")
}

// Функция воркера
func worker(tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		result := performMemoryIntensiveTask(task)
		results <- result
	}
}

// performMemoryIntensiveTask функция требующая много памяти
func performMemoryIntensiveTask(task int) int {
	// Создание среза большого размера
	data := make([]int, MemoryIntense)
	for i := 0; i < MemoryIntense; i++ {
		data[i] = i + task
	}

	// Имитация временной задержки
	time.Sleep(10 * time.Millisecond)

	// Вычисление результата
	result := 0
	for _, value := range data {
		result += value
	}
	return result
}
