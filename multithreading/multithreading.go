package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n, end int)  {
	for i := 0; i < end; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i, 5)
	}
	var input string
	fmt.Scanln(&input)
}