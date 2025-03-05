package main

import (
	"context"
	"fmt"
	"time"
)

func parse(ctx context.Context) { // она породила, только она и может отменить
	id := ctx.Value("id")
	fmt.Println(id.(int)) // приведение типа
	for {
		select {
		case <-time.After(time.Second * 2):
			fmt.Println("parsing completed")
			return
		case <-ctx.Done():
			fmt.Println("deadline exceded")
			return
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*3) // _ = cancel
	ctx = context.WithValue(ctx, "id", 1)

	// go func() {
	// 	time.Sleep(time.Millisecond * 100) // эта строка через 100мс отправляет отмену, но если ее не будет,
	//     // то парсинг успеет произойти
	// 	cancel()
	// }()

	parse(ctx)
}

// 1) context.Background() - использовать только на самом высоком уровне
// 2) context.TODO - когда не уверен, что нужно использовать (но потом надо убирать)
// 3) context.Value - использовать как можно реже, и передавать только не обязательные параметры
// 4) context всегда передается первым аргументом функции
// 5) в качестве context`а не пердавать nil
// 6) только функция, порождающая context может его отменить
