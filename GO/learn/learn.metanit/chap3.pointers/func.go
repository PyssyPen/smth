package main

import "fmt"

func changeValue(x *int) {
	*x = (*x) * (*x)
}

func createPointer(x int) *int {
	p := new(int)
	*p = x
	return p
}

func main() {

	d := 5
	fmt.Println("d before:", d) // 5
	changeValue(&d)             // изменяем значение
	fmt.Println("d after:", d)  // 25 - значение изменилось!

	p1 := createPointer(7)
	fmt.Println("p1:", *p1, "pointer:", p1) // p1: 7
	p2 := createPointer(10)
	fmt.Println("p2:", *p2, "pointer:", p2) // p2: 10
	p3 := createPointer(28)
	fmt.Println("p3:", *p3, "pointer:", p3) // p3: 28
}
