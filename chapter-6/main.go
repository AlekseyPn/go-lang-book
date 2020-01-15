package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	b := [5]float64{
		90,
		93,
		77,
		11,
		22,
	}

	var total float64 = 0
	for _, value := range b {
		total += value
	}

	fmt.Println(total / float64(len(b)))

	// срезы
	c := make([]float64, 5, 10)
	fmt.Println(c)

	arr := b[1:3]

	fmt.Println(arr)

	// функция срезов для добавления элементов в массив
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	// функция срезов для копирования массивов
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice4, slice3)
}
