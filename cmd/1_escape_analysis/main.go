package main

import "runtime"

// Run with: go run -gcflags=-m cmd/1_escape_analysis/main.go

func main() {
	var arrayBefore10Mb [1310720]int // Меньше 10 МБ
	arrayBefore10Mb[0] = 1

	var arrayAfter10Mb [1310721]int // Больше 10 МБ
	arrayAfter10Mb[0] = 1

	sliceBefore64 := make([]int, 8192) // Меньше 64 КБ
	sliceOver64 := make([]int, 8193)   // Больше 64 КБ
	sliceOver64[0] = sliceBefore64[0]

	runtime.GC()
}
