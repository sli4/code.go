package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	text = strings.ReplaceAll(text, "\n","")
	go Receive()
	for i:=0; i<10; i++ {
		fmt.Println("***", i)
		fmt.Println(text=="a")
		fmt.Println(text=="b")
		if text=="a" && i%2==0 {
			time.Sleep(time.Second*10)
			Send(i)
		}
		if text=="b" && i%2==1 {
			time.Sleep(time.Second*10)
			Send(i)
		}

	}

	time.Sleep(2*time.Minute)
}

var ch = make(chan int,10)
func Send(i int) {
	fmt.Println("发送消息：",i)
	ch <- i
}

func Receive() {
	for {
		select {
		case i,_ := <- ch:
			fmt.Println("收到消息：",i)
		}
	}
}