package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("sheep")
		wg.Done()
	}()
	wg.Wait()
}

func count(s string) {
	for i := 1; i<=5; i++ {
		fmt.Println(i, s)
		time.Sleep(time.Millisecond * 500)
	}
}
