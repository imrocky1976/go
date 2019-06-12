package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("hello, world, secends: %d\n", time.Now().UnixNano()/1e9)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Printf("My favorite number is %d\n", r1.Intn(int(time.Now().UnixNano()/1e9)))
}
