package main

import "fmt"

func main() {
	//for condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//for while
	counter:= 0
	for counter < 5 {
		fmt.Println(counter)
		counter ++
	}
}