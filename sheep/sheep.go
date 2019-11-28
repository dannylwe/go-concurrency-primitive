package main

import (
	"fmt"
	"time"
)

func main() {
	go count("sheep")
	go count("fish")

	fmt.Scanln()
}

func count(s string) {
	for i := 1; true; i++ {
		fmt.Println(i, s)
		time.Sleep(time.Millisecond * 500)
	}
}
