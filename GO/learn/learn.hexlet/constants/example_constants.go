package main

/*

const StatusOk int = 200

const (
    StatusOk = 200
    StatusNotFound = 404
)

package main

type Person struct {
}

func main() {
    // такие константы допустимы
    const (
        num = 20
        str = "hey"
        isValid = true
    )

    // нельзя объявить структуру как константу
    const p = Person{} // ошибка компиляции: const initializer Person{} is not a constant
}

const (
    // публичная константа, которую можно использовать во внешних пакетах
    StatusOk = 200

    // приватная константа, доступная только в рамках текущего пакета
    statusInternalError = 500
)

package main

import "fmt"

const defaultStatus = 200

func main() {
    const status = 404

    fmt.Println("default status:", defaultStatus) // default status: 200
    fmt.Println("current status:", status) // current status: 404
}

package main

import "fmt"

const (
    zero = iota
    one
    two
    three
)

const (
	a = iota
	b = 42
	c = iota
	d
)

func main() {
	fmt.Println(zero, one, two, three) // 0 1 2 3
	fmt.Println(a, b, c, d)            // 0, 42, 2, 3
}

*/
