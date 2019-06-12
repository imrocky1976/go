package main

import (
	"fmt"
	"time"
)

func say(what string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(what)
	}
}

func main() {
	go say("world")
	say("hello")
}
