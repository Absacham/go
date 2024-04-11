package main

import "fmt"

func main(){
	//Defer
	defer fmt.Println("Before Defer")
	
	// Continue y break
	for i := 0; i < 10; i++{
		fmt.Println(i)

		// Continue
		if i == 4 {
			fmt.Println("Es 4")
			continue
		}

		// Break
		if i == 6 {
			fmt.Println("Break")
			break
		}
	}
}