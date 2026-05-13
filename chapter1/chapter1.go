package chapter1

import "fmt"

func Run() {
	fmt.Println("Hello, Go :)")

	var name = "Amulya"
	age := 34
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	const pi = 3.149
	fmt.Println("Pi : ", pi)
}
