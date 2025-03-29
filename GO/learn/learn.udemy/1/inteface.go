// package main
package main

// import (
// 	"fmt"
// 	"math"
// )

// // неявная имплементация, не нужно присваивать каждой структуре интерфейс

// type Shape interface {
// 	ShapeWithArea
// 	ShapeWithPerimeter
// }

// type Square struct {
// 	sideLenght float32
// }

// type Circle struct {
// 	radius float32
// }

// func (s Square) Area() float32 {
// 	return s.sideLenght * s.sideLenght
// }

// func (c Circle) Area() float32 {
// 	return c.radius * c.radius * math.Pi
// }

// func PrintShapeArea(shape Shape) {
// 	// fmt.Println("qwerty", 1, true, 2312312.2233, shape)
// 	// здесь нет привязки к какому-либо типу, поскольку интерфейс пустой

// 	fmt.Println(shape.Area())
// }

// func main() {
// 	square := Square{5}
// 	circle := Circle{8}

// 	PrintShapeArea(square)
// 	PrintShapeArea(circle)

// 	//PrintInteface(square)
// 	//PrintInteface(circle)
// 	PrintInteface("qwerty")
// 	//PrintInteface(true)
// }

// // тут про пустой интерфейс

// func PrintInteface(i interface{}) {
// 	// switch value := i.(type) {  // тут тип неопределен
// 	// case int:
// 	// 	fmt.Println("int", value)
// 	// case bool:
// 	// 	fmt.Println("bool", value)
// 	// default:
// 	// 	fmt.Println("unknown type", value)
// 	// }
// 	// //fmt.Printf("%+v\n", i)

// 	str, ok := i.(string) // приведение интерфейса к конкретному типу
// 	if !ok {
// 		fmt.Println("interface is not string")
// 		return
// 	}
// 	fmt.Println(len(str))
// }

// type ShapeWithArea interface {
// 	Area() float32
// }

// type ShapeWithPerimeter interface {
// 	Area() float32
// }

// // по идее без этого работать не должно, но работает
// func (s Square) Perimeter() float32 {
// 	return s.sideLenght * 4
// }
