package main

import "fmt"

func main() {
	// 当使用 for 循环遍历切片时，每次迭代都会返回两个值。
	// 第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
