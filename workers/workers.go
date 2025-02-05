package main

import "fmt"

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	
	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for {
		select {
		case msg1 := <-results:
			fmt.Println(msg1)
		}
	}

}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
