package main

import "fmt"

func isPair(num int) bool {
	return num % 2 == 0
}

func isValidUser(userName, pass string) bool {
	return userName == "Alpha" && pass == "MyPassword"
}

func main()  {

	valor1 := 1
	valor2 := 2

	if valor1 == valor2 {
		fmt.Println("verdadero")
	} else {
		fmt.Println("Falso")
	}

	if isPair(6) {
		fmt.Println("Number is pair")
	} else {
		fmt.Println("Number is odd")
	}

	if isValidUser("Alpha5", "MyPassword") {
		fmt.Println("Credentials are valid")
	} else {
		fmt.Println("Credentials aren't valid")
	}

}
