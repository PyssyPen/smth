package main

/*
// создание пустой мапы
var m map[int]string
// сокращенное создание пустой мапы
m := map[int]string{}
// рекомендуемое создание с обозначением размера
m := make(map[int]string, 10)
// создание мапы с элементами
m := map[int]string{1: "hello", 2: "world"}
// добавление элемента
m[3] = "!" // map[1:hello, 2:world, 3:!]
// чтение элемента
word := m[1] // "hello"


elements := map[int64]bool{1: true, 2: false}
element, elementExists := elements[1] // true, true
element, elementExists := elements[2] // false, true
element, elementExists := elements[225] // false, false


// пустая структура struct{} — это тип данных, который занимает 0 байт
// используется, когда нужно проверять в мапе только наличие ключа
cache := make(map[string]struct{})
// проверяем есть ли ключ `key` в мапе
_, ok := cache["key"]
fmt.Println(ok)  // false
// добавим ключ и проверим вновь
cache["key"] = struct{}{}
_, ok := cache["key"]
fmt.Println(ok)  // true


engToRus := map[string]string{"hello":"привет", "world":"мир"}
delete(engToRus, "world")
fmt.Println(engToRus) // map[hello:привет]


package main
import (
    "fmt"
)
func main() {
    m := map[int]string{1: "hello", 2: "world"}
    modifyMap(m)
    fmt.Println(m) // вывод: map[1:changed 2:world 200:added]
}
func modifyMap(m map[int]string) {
    m[200] = "added"
    m[1] = "changed"
}



m := map[string]int{"foo": 1, "bar": 2}
elem := m["baz"]
fmt.Println(elem * 42)  //0






OK		m := make(map[int]string)
NOT		m := make(map[int]string{1: "foo", 2: "bar"})
OK		m := map[int]string{}
NOT		m := map[int]string
OK		m := map[string]int{"foo": 1, "bar": 2}
NOT		var m map[int]{1, 2, 3}string{"foo", "bar", "baz"}

*/
