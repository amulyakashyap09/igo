package chapter2

import "fmt"

func Run() {
	var i int = 10
	var i8 int8 = 127
	var i64 int64 = 9999999
	var f32 float32 = 3.14
	var f64 float64 = 3.14159265358979
	var boolean bool = true
	var b byte = 'a'
	var r rune = '世'
	var greeting string = "Hello, 世界"

	fmt.Printf("i is %d, i8 is %d, i64 is %d\n", i, i8, i64)
	fmt.Printf("f32 is %f, f64 is %f\n", f32, f64)
	fmt.Println("boolean ", boolean)
	fmt.Println("byte ", b)
	fmt.Println("rune ", r)
	fmt.Println("String ", greeting)

	var zeroInt int
	var zeroStr string
	var zeroBool bool
	var zeroFloat float64

	fmt.Println(zeroInt, zeroStr, zeroBool, zeroFloat)
}
