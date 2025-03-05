package main

import (
	"fmt"
	"time"
)

func main() {
	chn1 := make(chan string)
	chn2 := make(chan string)

	go func() {
		for {
			chn1 <- "канал 1. прошло 200 мс"
			time.Sleep(time.Millisecond * 200)
		}
	}()

	go func() {
		for {
			chn2 <- "канал 2. прошло 1 с"
			time.Sleep(time.Second)
		}
	}()

	for {
		select { //если писать вывод каналов без селекта,
		// то произойдет блокировка до получения сообщения в канале
		case msg := <-chn1:
			fmt.Println(msg)
		case msg := <-chn2:
			fmt.Println(msg)
		default:
		}
	}
}
