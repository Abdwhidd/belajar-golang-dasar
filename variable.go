package main

import "fmt"

func main() {
	// INI ADALAH MULTIPLE VARIABLE
	var (
		firtsName = "Abdul"
		lastName  = "Wahid"
	)
	fmt.Println(firtsName)
	fmt.Println(lastName)

	// VARIABLE NORMAL
	var friendName string
	friendName = "Budi"
	fmt.Println(friendName)

	// VARIABLE YANG LEBIH SINGKAT
	countryName := "Makassar"
	fmt.Println(countryName)
}
