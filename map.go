package main

import "fmt"

func main() {
	person := map[string]string{
		"name":  "Wahid",
		"noKtp": "0303030",
	}

	// Note: dibawah adalah code untuk menambah key dan value baru
	person["hoby"] = "run"

	fmt.Println(person)
	fmt.Println(len(person))

}
