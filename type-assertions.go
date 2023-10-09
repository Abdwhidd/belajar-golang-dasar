package main

import "fmt"

func random() interface{} {
	return 3.4
}

func main() {
	var result interface{} = random()

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is a String")
	case int:
		fmt.Println("Value", value, "is a int")
	default:
		fmt.Println("Ukhnown")
	}

}
