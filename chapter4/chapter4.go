package chapter4

import (
	"fmt"
	"runtime"
)

const randomValue = 10

func testIfElse(a int) {
	if a > randomValue {
		fmt.Println("A is greater than 10")
	} else if a == 10 {
		fmt.Println("A is equal to 10")
	} else {
		fmt.Println("A is less than 10")
	}
}

func testForLoop() {
	for i := 0; i < randomValue; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func testWhileLoop() {
	n := 20
	for {
		n--
		if n <= 0 {
			break
		}
		fmt.Printf("%d ", n)
	}
}

func Run() {
	os := runtime.GOOS

	switch os {
	case "darwin":
		fmt.Println("MacOs")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("Default: MacOs")
	}

	testIfElse(5)
	testIfElse(15)
	testIfElse(10)

	testForLoop()
	testWhileLoop()

}
