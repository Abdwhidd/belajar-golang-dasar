package main

import "fmt"

func main() {

	// Note : panjang array(3) sudah tidak bisa lagi ditambah atau diubah jika sudah ditetapkan
	var friendName = [3]string{
		"Budi",
		"Joko",
		"Jokowi",
	}
	fmt.Println(len(friendName))
	fmt.Println(friendName[0])
	fmt.Println(friendName[1])
	fmt.Println(friendName[2])
	
}
