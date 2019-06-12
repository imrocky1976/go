package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, ss := range strings.Fields(s) {
		m[ss] += 1
	}
	return m
}

func main() {
	m := WordCount("this is a test, this is right.")
	fmt.Printf("word count: %v\n", m)

	fmt.Println(s)
}

// 映射的文法
type Vertex struct {
	Lat, Long float64
}

var s = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

/*
修改映射

在映射 m 中插入或修改元素：
m[key] = elem

获取元素：
elem = m[key]

删除元素：
delete(m, key)

通过双赋值检测某个键是否存在：
elem, ok = m[key]

若 key 在 m 中，ok 为 true ；否则，ok 为 false。

若 key 不在映射中，那么 elem 是该映射元素类型的零值。

同样的，当从映射中读取某个不存在的键时，结果是映射的元素类型的零值。

注 ：若 elem 或 ok 还未声明，你可以使用短变量声明：
elem, ok := m[key]
*/
