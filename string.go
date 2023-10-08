package main

import "fmt"

func main() {
	firstName := "Abdul"
	lastName := "Abdul Wahid"

	fmt.Println(len(firstName)) // resultnya adalah 5
	fmt.Println(lastName[0])    // resultnya 65 karena bytenya "A" adalah 69
	fmt.Println(lastName[9])    // resultnya 105 karena bytenya "d" adalah 105
	fmt.Println("Wahid"[1])     // resultnya 97 karena bytenya "w" adalah 97

	fmt.Printf("halo %s", firstName) // resultnya halo Abdul%
}
