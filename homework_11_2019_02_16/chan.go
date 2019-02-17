package main

import (
	"fmt"
	"time"
)

func main() {
	msgChannel := make(chan string, 1)
	msgChannel <- "ping"
	done := make(chan struct{})
	for i := 0; i < 2; i++ {
		go player(msgChannel, done)
	}
	time.Sleep(5 * time.Second)
	done <- struct{}{}
	fmt.Println("Exit from program")
}

func player(msg chan string, done chan struct{}) {
	for {
		select {
		case <-done:
			return
		case m := <-msg:
			if m == "ping" {
				time.Sleep(300 * time.Millisecond)
				msg <- "pong"
			}
			if m == "pong" {
				time.Sleep(300 * time.Millisecond)
				msg <- "ping"
			}
			fmt.Println(m)
		}
	}
}
