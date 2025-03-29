package main

import (
	"fmt"
)

type Vehicle interface {
	move()
}

func drive(vehicle Vehicle) {
	vehicle.move()
}

type Car struct{}
type Aircraft struct{}

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}
func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}

/*
func main() {

	tesla := Car{}
	boing := Aircraft{}
	drive(tesla)
	drive(boing)
}


type vehicle interface {
	move()
}

type Car struct{}
type Aircraft struct{}

func (c Car) move() {
	fmt.Println("car is moving")
}

func (a Aircraft) move() {
	fmt.Println("plane is flying")
}

func drivecar(c Car) {
	c.move()
}

func driveaircraft(a Aircraft) {
	a.move()
}

func main() {

	var tesla Vehicle = Car{}
    var boing Vehicle = Aircraft{}
    tesla.move()
    boing.move()

	var mercedes Car = Car{}
	var boing Aircraft = Aircraft{}
	drivecar(mercedes)
	driveaircraft(boing)
}
*/
