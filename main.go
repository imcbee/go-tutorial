package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	var myName string = "John Doe"
	const myOtherName string = "Jane Doe"

	myThirdName := "Bob Doe"

	fmt.Println(myName)
	fmt.Println(myOtherName)
	fmt.Println(myThirdName)
}