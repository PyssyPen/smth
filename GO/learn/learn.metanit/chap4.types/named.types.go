package main

import "fmt"

type BinaryOp func(int, int) int

func action(n1 int, n2 int, op BinaryOp) {

	result := op(n1, n2)
	fmt.Println(result)
}

func add(x int, y int) int {

	return x + y
}

/*
func main() {

	var myOperation BinaryOp = add
	action(10, 35, myOperation) // 45
}
*/
