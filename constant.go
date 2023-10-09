package main

import "fmt"

func main() {
	// Note: Constant WAJIB langsung dideklarasikan valueanya

	// Multiple constant
	const (
		firtsName = "Wahid"
		lastName  = "Kahar"
	)

	// Normal Constant
	const nameLengkap = "Tes"

	fmt.Println(nameLengkap)
	fmt.Println(firtsName)
	fmt.Println(lastName)
}
