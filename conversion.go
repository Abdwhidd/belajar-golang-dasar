package main

import "fmt"

func main() {
	name := "Wahid"        // Resultnya adalah "Wahid"
	e := name[0]           // Resultnya adalah 87 karena btye "W" adalah 87
	eToString := string(e) // Resultnya adalah "W" karena value dari variabel e dikonversi menjadi string

	fmt.Println(e)
	fmt.Println(eToString)
}
