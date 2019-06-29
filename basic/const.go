package main

import (
	"fmt"
	"reflect"
)

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	a := [...]string{"no error", "Eio", "invalid argument"}
	s := []string{"no error", "Eio", "invalid argument"}
	m := map[int]string{1: "no error", 2: "Eio", 3: "invalid argument"}
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(s), reflect.TypeOf(m))
	fmt.Println(a, s, m)
}
