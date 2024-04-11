package main

import (
	"fmt"
	"strings"
)


func isPalindrome(text string) bool {
	var textReverse string
	/* Reto: validar que si el texto contiene 
	mayúsculas igual siga validando el palíndromo. */
	text = strings.ToLower(text) // Pasamos el texto a minúsculas
	text = strings.ReplaceAll(" ", text, text) // Eliminamos los espacios
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	
	if text == textReverse {
		return true
	} else {
		return false
	}

}

func main()  {

	/* Recorrer slices */
	slice2 := []string{"cero", "uno"}

	for indice, dato := range slice2 {
		fmt.Println(indice, dato)
	}
	
	slice := []string{"hola", "golang"}

	for _, v := range slice {
		fmt.Println(v)
	}

	
	fmt.Println(isPalindrome("Anita lava la tina"))
}