package main

/*
type Person struct {
    // [название поля] [тип данных]
    Name string
    Age int
}

func main() {
    p := Person{Name: "John", Age: 25}

    p.Name // "John"
    p.Age // 25
}

//////////////////////////////////////////////////

func main() {
    p := Person{}

    p.Name // ""
    p.Age // 0
}

/////////////////////////////////////////////////////////

type Person struct { // структура публична
    Name string // поле публично

    wallet wallet // поле приватно: можно обращаться только внутри текущего пакета
}

type wallet struct { // структура приватна: можно инициализировать только внутри текущего пакета
    id string
    moneyAmount float64
}

type User struct {
    ID int64 `json:"id" validate:"required"`
    Email string `json:"email" validate:"required,email"`
    FirstName string `json:"first_name" validate:"required"`
}

////////////////////////////////////////////////////////////////

package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    ID        int64  `json:"id"`
    Email     string `json:"email"`
    FirstName string `json:"first_name"`
}

func main() {
    u := User{}
    u.ID = 22
    u.Email = "test@test.com"
    u.FirstName = "John"

    bs, _ := json.Marshal(u)

    fmt.Println(string(bs)) // {"id":22,"email":"test@test.com","first_name":"John"}
}

///////////////////////////////////////////////////////////////////////

package main

import (
    "fmt"
    "github.com/go-playground/validator/v10"
)

type User struct {
    ID        int64  `validate:"required"`
    Email     string `validate:"required,email"`
    FirstName string `validate:"required"`
}

func main() {
    // создали пустую структуру, чтобы проверить валидацию
    u := User{}

    // создаем валидатор
    v := validator.New()

    // метод Struct валидирует переданную структуру и возвращает ошибку `error`, если какое-то поле некорректно
    fmt.Println(v.Struct(u))
}



Вывод программы:

Key: 'User.ID' Error:Field validation for 'ID' failed on the 'required' tag
Key: 'User.Email' Error:Field validation for 'Email' failed on the 'required' tag
Key: 'User.FirstName' Error:Field validation for 'FirstName' failed on the 'required'
*/
