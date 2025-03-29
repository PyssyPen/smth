package main

import (
	"fmt"
)

func main() {
	num := [5]int{1, 2, 3, 4, 5}
	im, valm := 0, 0
	fmt.Println("Ввеедите номер, который хотите изменить (от 1 до 5): ")
	fmt.Scanln(&im)
	fmt.Println("Ввеедите то, на что хотите изменить: ")
	fmt.Scanln(&valm)
	r := SafeWrite(num, im, valm)
	fmt.Println(r)
}

func SafeWrite(nums [5]int, i, val int) [5]int {
	if i <= len(nums) && i > -1 {
		nums[i-1] = val
	} else {
		fmt.Println("Выбранный вами элемент не принадлежит массиву!")
	}
	return nums
}
