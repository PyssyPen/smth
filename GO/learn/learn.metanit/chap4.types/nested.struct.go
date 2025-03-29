package main

/*
type contact struct {
	email string
	phone string
}

type man struct {
	name string
	age  int
	contact
}

func main() {

	var tom = man{
		name: "Tom",
		age:  24,
		contact: contact{
			email: "tom@gmail.com",
			phone: "+1234567899",
		},
	}
	tom.email = "supertom@gmail.com"

	fmt.Println(tom.name)
	fmt.Println(tom.age)
	fmt.Println(tom.email) // supertom@gmail.com
	fmt.Println(tom.phone) // +1234567899
}





///////////////////////////////////////////////

type node struct {
	value int
	next  *node
}

// рекурсивный вывод списка
func printNodeValue(n *node) {

	fmt.Println(n.value)
	if n.next != nil {
		printNodeValue(n.next)
	}
}

func main() {

	first := node{value: 4}
	second := node{value: 5}
	third := node{value: 6}

	first.next = &second
	second.next = &third

	var current *node = &first
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
	fmt.Println(node{})
}
*/
