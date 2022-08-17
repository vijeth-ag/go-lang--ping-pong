package main

import (
	"fmt"
	"time"
)

func helloBack(ch chan string) {
	for {
		select {
		case data := <-ch:
			fmt.Println(data)
			time.Sleep(time.Second * 1)
			ch <- "ping"
		}
	}
}

func main() {
	ch := make(chan string)

	go helloBack(ch)
	ch <- "serve"
	for {
		select {
		case data := <-ch:
			fmt.Println(data)
			time.Sleep(time.Second * 1)
			ch <- "pong"
		}
	}
}
