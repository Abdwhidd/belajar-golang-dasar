package main

import "fmt"

type CustomerService struct {
	Name  string
	Title int
}

// Note: ini adalah method/function yang seakan-akan milik dari struct Customers
func (customer CustomerService) sayHelloCustomer(customerService string) {
	fmt.Println("Hello", customerService, "my name is", customer.Name)
}

func (cs CustomerService) getStarted() {
	fmt.Println("Let's go", cs.Name)
}

func main() {
	customers := CustomerService{Name: "Wahid"}
	customers.sayHelloCustomer("Joko")
	customers.getStarted()
}
