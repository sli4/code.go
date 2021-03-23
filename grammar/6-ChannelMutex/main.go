package main

import (
	"fmt"
	"sync"
	"time"
)

func chanDemo(ch chan string, worker string) {
	for {
		select {
		case v:=<-ch:
			fmt.Println("收到老板指令：",v)
		default:
			fmt.Println("无事可做，好无聊")
		}
		time.Sleep(time.Second)
	}
}
var m =  sync.Mutex{}
func mutexDemo(m sync.Mutex, worker string) {
	for {
		m.Lock()
		fmt.Println("我是",worker,",我在打黑工")
		m.Unlock()
		time.Sleep(time.Second)
	}
}

func main() {
	go mutexDemo(m, "小马哥")
	go mutexDemo(m, "周润发")
	time.Sleep(time.Minute)
}
