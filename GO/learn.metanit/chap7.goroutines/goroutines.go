package main

// import "fmt"

// func main() {

// 	for i := 1; i < 7; i++ {
// 		go factorial(i)
// 	}
// 	fmt.Scanln() // ждем ввода пользователя
// 	fmt.Println("The End")
// }

// func factorial(n int) {
// 	if n < 1 {
// 		fmt.Println("Unvalid input number")
// 		return
// 	}
// 	result := 1
// 	for i := 1; i <= n; i++ {
// 		result *= i
// 	}
// 	fmt.Println(n, "-", result)
// }

/*
Горутины также могут представлять вызовы анонимных функций:

package main
import "fmt"

func main() {

    for i := 1; i < 7; i++{

        go func(n int){
            result := 1
            for j := 1; j <= n; j++{
                result *= j
            }
            fmt.Println(n, "-", result)
        }(i)
    }
    fmt.Scanln()
    fmt.Println("The End")
}
*/
