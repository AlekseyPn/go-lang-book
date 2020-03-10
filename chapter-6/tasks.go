package main

import "fmt"

func main() {
	x := [5]int{1, 2, 3, 4, 5}
	fmt.Println(x[4])

	// вырезать 3 элемента из массива интов длинной 9
	a := make([]int, 3, 9)
	fmt.Println(len(a))
	fmt.Println(a)

	b := [6]string{"a", "b", "c", "d", "e", "f"}

	// вернет 2,3,4 элемент по индексу не включая 5
	fmt.Println(b[2:5])

	v := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	res := v[0]

	// находим минимальное число из массива интов
	for _, value := range v {
		if value < res {
			res = value
		}
	}

	fmt.Println(res)
}
