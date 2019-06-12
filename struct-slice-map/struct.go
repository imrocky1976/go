package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func main() {
	v := Vertex{1, 2} // 短变量声明
	fmt.Println(v)
	fmt.Println(Vertex{1, 2})
	var i int // 变量声明
	j := i    // 类型推导
	fmt.Println(j)
	var t Vertex = Vertex{1, 2} // 变量初始化
	t.X = 5.1
	fmt.Println(t)

	// 结构体指针
	p := &v
	p.X = 1e9
	fmt.Println(v)

	// 结构体文法
	var (
		v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
		v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
		v3 = Vertex{}      // X:0 Y:0
		p1 = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
	)

	fmt.Println(v1, p1, v2, v3)

}
