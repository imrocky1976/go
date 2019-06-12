package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从c中接收

	fmt.Println(x, y, x+y)
}

/*

带缓冲的信道
信道可以是 带缓冲的。将缓冲长度作为第二个参数提供给 make 来初始化一个带缓冲的信道：

ch := make(chan int, 100)
仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。


package main

import "fmt"
import "time"

func consumer(ch chan int) {
	for {
		fmt.Println("consumer", <-ch)
		time.Sleep(time.Second)
	}

}
func main() {
	ch := make(chan int, 1)
	go consumer(ch)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	//fmt.Println(<-ch)
}

*/
