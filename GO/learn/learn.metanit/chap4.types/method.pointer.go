package main

import "fmt"

type person struct {
	name string
	age  int
	smth int
}

func (p *person) updateAge(newAge int) {
	p.age = newAge
}

// я поставил указатель на входе в функцию, и это сработало,
// но предложенный вариант решения внизу в комменте

func main() {

	var tom = person{name: "Tom", age: 24, smth: 111}
	fmt.Println("before", tom.age)
	tom.updateAge(33)
	fmt.Println("after", tom.age)
	fmt.Println(tom)
}

/*
type person struct{
    name string
    age int
}
func (p *person) updateAge(newAge int){
    (*p).age = newAge
}

func main() {

    var tom = person { name: "Tom", age: 24 }
    var tomPointer *person = &tom
    fmt.Println("before", tom.age)
    tomPointer.updateAge(33)
    fmt.Println("after", tom.age)
}
*/
