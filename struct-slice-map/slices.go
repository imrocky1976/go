package main

import (
	"fmt"
	"strings"
)

func main() {
	primes := [6]int{1, 2, 3, 4, 5, 6}
	var p []int = primes[1:4]
	fmt.Println(primes, p)

	// 切片就像数组的引用
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// 切片文法，类似没有长度的数组
	// 构造一个数组
	i := [3]bool{false, true, false}
	// 构造一个类似上面的数组，然后构造一个引用它的切片
	ii := []bool{false, true, false}
	fmt.Println(i, ii)

	// 结构体切片
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
	}
	fmt.Println(s)

	// 切片默认行为
	// 上界默认为该切片长度,下界默认为0
	//var a [10]int
	// 下列切片等价
	/*
		a[0:10]
		a[:10]
		a[0:]
		a[:]
	*/

	lenAndCap()

	nilTest()

	makeSlice()

	slice2()

	appendToSlice()
}

// 切片的长度和容量
// 切片的长度就是它所包含的元素个数。
// 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
func lenAndCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func nilTest() {
	//切片的零值是 nil
	//nil 切片的长度和容量为 0 且没有底层数组
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func makeSlice() {
	// 用 make 创建切片
	// 切片可以用内建函数 make 来创建，这也是你创建动态数组的方式。

	// make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：

	a := make([]int, 5) // len(a)=5

	// 要指定它的容量，需向 make 传入第三个参数：

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5

	b = b[:cap(b)] // len(b)=5, cap(b)=5
	b = b[1:]      // len(b)=4, cap(b)=4
}

// 切片的切片
func slice2() {
	// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// 向切片追加元素
func appendToSlice() {
	var s []int
	printSlice(s)

	// 添加一个空切片
	s = append(s, 0)
	printSlice(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	printSlice(s)
}

/*
实现 Pic。它应当返回一个长度为 dy 的切片，其中每个元素是一个长度为 dx，元素类型为 uint8 的切片。
当你运行此程序时，它会将每个整数解释为灰度值（好吧，其实是蓝度值）并显示它所对应的图像。

图像的选择由你来定。几个有趣的函数包括 (x+y)/2, x*y, x^y, x*log(y) 和 x%(y+1)。

（提示：需要使用循环来分配 [][]uint8 中的每个 []uint8；
请使用 uint8(intValue) 在类型之间转换；你可能会用到 math 包中的函数。）
*/
func Pic(dx, dy int) [][]uint8 {

	pic := make([][]uint8, dy)

	for y := range pic {

		pic[y] = make([]uint8, dx)

		for x := range pic[y] {
			pic[y][x] = uint8((x + y) / 2)
		}
	}

	return pic
}
