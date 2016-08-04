package main
import (
	"fmt"
	"sync"
)

var n, m int
var active int64 = 0
var lock sync.Mutex
func atomicAdd(v int64) {
	lock.Lock()
	active += v
	lock.Unlock()
}
func next(r, x, y, z, v int, chn chan int) {
	y = (y << 1) & m
	z >>= 1
	p := x | y | z
	for i, b:= 0, 1; i < n; i, b = i + 1, b << 1 {
		if (p & b) == 0 {
			if r == n - 1 {
				chn <- v
			} else {
				next(r + 1, x | b, y | b, z | b, v, chn)
			}
		}
	}
}

func next_parallel(r, x, y, z, v int, chn chan int) {
	y = (y << 1) & m
	z >>= 1
	p := x | y | z
	for i, b:= 0, 1; i < n; i, b = i + 1, b << 1 {
		if (p & b) == 0 {
			if r == n - 1 {
				chn <- v
			} else {
				next(r + 1, x | b, y | b, z | b, v, chn)
			}
		}
	}
	atomicAdd(-1)
}

func main() {
	mill := 1000000
	pipe := make(chan int, mill)
	fmt.Scanf("%d", &n)
	m = (1 << uint(n)) - 1
	b := 1
	for i := 0; i < n / 2; i, b = i + 1, b << 1 {
		atomicAdd(1)
		go next_parallel(1, b, b, b, 2, pipe)
	}
	if(n & 1 == 1) {
		atomicAdd(1)
		go next_parallel(1, b, b, b, 1, pipe)
	}
	counter := 0
	for ; active != 0; {
		select {
			case v := <- pipe:
				counter += v
			default:
		}
	}
	println(counter)
}