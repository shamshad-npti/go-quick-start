package main

import (
	//"fmt"
	"sync"
	"time"
)
var lock sync.Mutex
/* Unsafe counter */
func increment(val *int) {
	for i := 0; i < 100000; i++ {
		*val++
	}
}

/* Safe counter */
func syncIncrement(val *int) {
	lock.Lock()
	for i := 0; i < 100000; i++ {
		*val++
	}
	defer lock.Unlock()
}

func main() {
	val := 0
	go increment(&val)
	go increment(&val)
	safeVal := 0
	go syncIncrement(&safeVal)
	go syncIncrement(&safeVal)
	time.Sleep(time.Millisecond)
	println(val)
	println(safeVal)
}