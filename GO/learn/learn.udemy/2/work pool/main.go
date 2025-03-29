package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	const jobsCount, workerCount = 15, 3
	jobs := make(chan int, 15)
	results := make(chan int, 15)
	//defer close(jobs)
	//defer close(results)

	for i := 0; i < workerCount; i++ {
		go worker(i, jobs, results)
	}

	for i := 0; i < jobsCount; i++ {
		jobs <- i + 1
	}

	for i := 0; i < jobsCount; i++ {
		fmt.Printf("result #%d : value = %d\n", i+1, <-results)
	}

	fmt.Println("time elapsed:", time.Since(t).String())
}

func worker(id int, jobs <-chan int, results chan<- int) { // если написано <-chan, то это только для чтиния,
	for j := range jobs { // а вот так chan<-, только писать
		time.Sleep(time.Second)
		fmt.Printf("worker #%d finished\n", id+1)
		results <- j * j
	}
}
