//go:build goexperiment.arenas
// +build goexperiment.arenas

package main

import (
	"arena"
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

const (
	numberOfArrays = 100
	arrayLen       = 2000
)

func main() {
	// Запись в trace файл
	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()

	mem := arena.NewArena()
	for j := 0; j < numberOfArrays; j++ {
		o1 := arena.MakeSlice[float64](mem, arrayLen, arrayLen)
		for i := 0; i < arrayLen; i++ {
			o1[i] = float64(i)
		}
		time.Sleep(10 * time.Millisecond)
	}
	t := arena.New[int](mem)
	value := 42
	t = &value
	fmt.Println(*t)
	mem.Free()
}
