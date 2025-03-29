package main

// import (
// 	"fmt"
// )

// type Age int

// func (a Age) isAdult() bool {
// 	return a >= 18
// }

// type User struct {
// 	name   string
// 	age    Age
// 	sex    string
// 	weight int
// 	height int
// }

// func (u *User) setName(name string) {
// 	u.name = name
// }

// func (u User) getName() string {
// 	return u.name
// }

// func NewUser(name, sex string, age, weight, height int) User {
// 	return User{
// 		name:   name,
// 		sex:    sex,
// 		age:    Age(age),
// 		weight: weight,
// 		height: height,
// 	}
// }

// func main() {
// 	user1 := NewUser("Vasya", "Male", 25, 60, 180)
// 	user2 := User{"Petya", 32, "Male", 87, 190}

// 	user1.setName("vlados")
// 	fmt.Println(user1.age.isAdult())

// 	fmt.Println(user1.getName())
// 	fmt.Println(user2.getName())
// }
