package main

import "fmt"

func main() {
	var p *int
	// fmt.Println(*p) error: invalid memory address or nil pointer dereference

	i := 42
	p = &i
	fmt.Println(*p) // 通过指针 p 读取 i

	*p = 21 // 通过指针 p 设置 i
	fmt.Println(*p)

	// 这也就是通常所说的“间接引用”或“重定向”
	// 与 C 不同，Go 没有指针运算
}
