package main

import "sync"


type MyInt struct {
	mu sync.Mutex
	val int
}

func demo1() {
	var i MyInt

	// i.mu is usable without explicit initialisation.
	i.mu.Lock()
	i.val++
	i.mu.Unlock()
}


func main() {
	demo1()
}
