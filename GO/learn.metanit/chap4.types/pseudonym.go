// Именные типы отличаются (нельзя использовать тип kilometer там,
// где используется mile) от псевдонимов (такое можно делать) инициацией и использованием
// ИМЕННЫЕ БЕЗ =, ПСЕВОДОНИМЫ С =

package main

import "fmt"

type mile = uint
type kilometer = uint

func distanceToEnemy(distance mile) {

	fmt.Println("расстояние для противника:")
	fmt.Println(distance, "миль")
}

/*
func main() {

	var distance mile = 5
	distanceToEnemy(distance)

	var distance1 uint = 5
	distanceToEnemy(distance1) // норм

	var distance2 kilometer = 5
	distanceToEnemy(distance2) // норм
}
*/
