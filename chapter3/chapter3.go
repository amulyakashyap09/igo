package chapter3

import "fmt"

type Celsius float64
type Fahrenheit float64

func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func Run() {

	c := Celsius(25)
	f := c.ToFahrenheit()

	fmt.Println("Temp in Fah: ", f)
}
