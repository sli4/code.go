package main

import (
	"fmt"
	"time"
)

var ch = make(chan int, 1)
func read(i int, ch chan int) {
	j := <- ch
	fmt.Println(i, "第一次从channel中读取数据:", j)
	j = <- ch
	fmt.Println(i, "第二次从channel中读取数据:", j)
}
func write(j int, ch chan int) {
	ch <- j
}
func main() {
	i:=1
	for ;i<6;i++{
		go read(i,ch)
	}
	time.Sleep(2*time.Second)

	j:=10
	for ;j<20;j++ {
		write(j,ch)
	}
	for {

	}
}
