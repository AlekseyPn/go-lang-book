package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	for k := 1; k <= 100; k++ {
		if k%3 == 0 {
			if k%5 == 0 {
				fmt.Println("FizzBuzz")
			}
			fmt.Println("Fizz")
		}

		if k%5 == 0 {
			fmt.Println("Buzz")
		}
	}
}
