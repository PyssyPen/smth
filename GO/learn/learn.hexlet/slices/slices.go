package main

import (
	"fmt"
)

/*func Remove(nums []int, i int) []int {
	if i <= len(nums) && i > 0 {
		l := len(nums)
		for f := i; f < l; f++ {
			nums[f-1] = nums[f]
		}
		m := l - 1
		n := make([]int, 0, m)
		for i := 0; i < m; i++ {
			n = append(n, i+1)
		}
		for l := 0; l < m; l++ {
			n[l] = nums[l]
		}
		return n
	}
	fmt.Println("Выбранный вами элемент не принадлежит массиву!")
	return nums
}*/

func Remove(nums []int, i int) []int {
	if i < 0 || i > len(nums)-1 {
		return nums
	}
	nums[i] = nums[len(nums)-1]
	return nums[:len(nums)-1]
}

/*
НА МЕТАНИТЕ ВСЕ ЕЩЕ ПРОЩЕ

users := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
//удаляем 4-й элемент
var n = 3
users = append(users[:n], users[n+1:]...)
fmt.Println(users)      //["Bob", "Alice", "Kate", "Tom", "Paul", "Mike", "Robert"]

*/

func main() {
	v := 0
	im := 0
	fmt.Println("Введите количество эллементов в массиве: ")
	fmt.Scanln(&v)
	num := make([]int, 0, v) // первое длина, второе вместимость (будет сделан массив, куда вместится
	for i := 0; i < v; i++ { // v эллементов, но их там не будет)
		num = append(num, i+1)
	}
	fmt.Println("Поллученный массив: ", num)
	fmt.Println("Ввеедите номер, который хотите удалить (от 0 до", v, "): ")
	fmt.Scanln(&im)
	r := Remove(num, im-1)
	fmt.Println(r)
}
