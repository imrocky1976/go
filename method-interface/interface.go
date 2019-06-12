package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat 实现了 Abser

	fmt.Println(a.Abs())

	a = &v // a *Vertex 实现了 Abser

	fmt.Println(a.Abs())

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	//a = v
}

type MyFloat float64

func (v MyFloat) Abs() float64 {
	if v < 0 {
		return float64(-v)
	}
	return float64(v)

}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
