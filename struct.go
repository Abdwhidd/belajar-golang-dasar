package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {

	var newUser Customer
	newUser.Name = "Wahid"
	newUser.Address = "Makassar"
	newUser.Age = 20

	fmt.Println(newUser)
	//user := Customer{
	//	"Wahid",
	//	"Makassar",
	//	20,
	//}
	//fmt.Println(user)

}
