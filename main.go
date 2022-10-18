package main

import "fmt"

var dogName string = "Sheldun"

func main() {
	frst()
	scnd()
}

func frst() {
	fmt.Println("My dog's name is", dogName)
}

func scnd() {
	fmt.Println("My friend's dog", dogName, "is awesome!")
}
