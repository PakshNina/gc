package main

import (
	"fmt"
	"sync"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return []int{}
		},
	}

	// Получение объекта из пула
	obj := pool.Get().([]int)
	obj = []int{42}
	fmt.Println("Получен объект из пула:", obj)

	// Возврат объекта в пул
	pool.Put(obj)
	fmt.Println("Объект возвращен в пул")
}
