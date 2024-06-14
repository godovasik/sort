package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Int() % 100)
	}
}
