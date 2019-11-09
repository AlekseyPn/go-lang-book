package main

import "fmt"

func main() {
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println(len("Hello"))
	fmt.Println("Hello"[1])
	fmt.Println("He" + "Lo")
	fmt.Println(true && true)
	fmt.Println(false && true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(32123 * 42452)
	fmt.Println((true && false) || (false && true) || !(false && false))
}
