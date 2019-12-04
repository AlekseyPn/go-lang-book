package main

import "fmt"

func main() {
	var x string
	x = "Hello, World"
	fmt.Println(x)
	x = x + "first"
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)
	fmt.Println(x == "second")

	b := "Hell is exist"
	fmt.Println(b)

	dogName := "Frenky"
	fmt.Println("My dog's name is ", dogName)
	//scope of variable limited by literal and visible for nested literals
	// for example

	// constants

	const c = "Hello"
	// throw error
	// c = "World"

	// declare many variables
	var (
		a = "c"
		v = "v"
		g = "G"
	)
	fmt.Println(a, v, g)
}

func f() {
	// is throw error x is undefined
	// fmt.Println(x)
}
