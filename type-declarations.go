package main

import "fmt"

func main() {

	type noKTP string
	type merriedStatus bool

	var noKTPUser noKTP = "4994949499499449"
	var merriedStatuUser merriedStatus = true

	fmt.Println(noKTPUser)
	fmt.Println(merriedStatuUser)
}
