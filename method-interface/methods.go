package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 值接收者， 传入的值为原始值的副本
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 指针接收者方法可以修改接收者指向的值
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/*
使用指针接收者的原因有二：

首先，方法能够修改其接收者指向的值。

其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。

通常来说，所有给定类型的方法都应该有值或指针接收者，但并不应该二者混用。
*/

// 内置类型的接收者，仅对该包有效
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())

	p := &v
	p.Scale(0.01)
	fmt.Println(p.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
