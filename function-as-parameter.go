package main

import "fmt"

type Filter func(string) string

func getSayhelloFiltered(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Gila" {
		return "*****"
	} else {
		return name
	}
}

func main() {
	getSayhelloFiltered("Wahid", spamFilter)
}
