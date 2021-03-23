package main

import "fmt"

func main() {
	p := new([]int)
	fmt.Printf("P:%p,%v\n", p, p)
	for _, i := range *p {
		fmt.Println("I:",i)
	}


	m := make([]int, 10)
	fmt.Printf("M:%p,%v\n", m, m )
	for _, i := range m {
		fmt.Println("I:", i)
	}

	(*p)[0] = 2
	m[0] =2

	*p = make([]int,3)
	(*p)[0] = 2
	fmt.Printf("P:%p,%v\n", p, p)
}
