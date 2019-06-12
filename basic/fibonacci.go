package main

import (
	"fmt"
	"time"
)

/* version 1
// 返回一个“返回int的函数”
func fibonacci() func() int {
	//var a,b int = 0,1
	a, b := 0, 1
	return func() (rlt int) {
		rlt = a
		a, b = b, a+b
		return
	}
}*/

// version 2
func fibonacci() func() int {
	x, y, z := 0, 1, 0
	return func() int {
		z, x, y = x, y, x+y
		return z
	}
}

func main() {
	f := fibonacci()
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		//fmt.Println(f())
		f()
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("elapsed time: %v\n", elapsed)
}
