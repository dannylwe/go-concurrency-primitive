package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("sheep", c)

	for msg := range c {
		fmt.Println(msg)
	}
	
}

func count(s string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- s
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
