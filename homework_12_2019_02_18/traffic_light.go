package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	m := map[int]string{0: "green", 1: "yellow", 2: "red"}
	for i := 0; i < 3; i++ {
		c := make(chan string)
		go sendColor(m[rand.Intn(len(m))], c)
		go receiveColor(c)
	}
	<-time.After(10 * time.Second)
}

func sendColor(color string, ch chan string) {
	for range time.NewTicker(time.Duration(rand.Uint32() / 1000)).C {
		switch {
		case color == "green":
			time.Sleep(time.Second)
		case color == "yellow":
			time.Sleep(2 * time.Second)
		case color == "red":
			time.Sleep(3 * time.Second)
		}
		ch <- color + fmt.Sprintf(" address of channel is %v", ch)
	}
}

func receiveColor(ch chan string) {
	for {
		if color, ok := <-ch; ok {
			fmt.Println(color)
		}
	}
}
