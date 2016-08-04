package main
import (
	"fmt"
	"time"
)

func print_message(id int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: %d\n", id, i)
	}
}

func producer(message chan string) {
	for i:= 0; ; i++ {
		message <- "Hello world!"
	}
}

func producer1(message chan string) {
	for i := 0; ; i++ {
		message <- "Ping"
	}
}

func consumer(message chan string) {
	for i := 0; ; i++ {
		fmt.Printf("%s\n", <- message)
		time.Sleep(time.Second / 100)		
	}
}

func main() {
	go print_message(1)
	go print_message(2)
	fmt.Scanln()
	chn := make(chan string, 100)
	go producer(chn)
	go producer1(chn)
	go consumer(chn)
	var input string
	fmt.Scanln(&input)
}