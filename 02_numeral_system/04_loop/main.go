package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 200; i++ {
		fmt.Printf("%d , %b , %x \n", i, i, i)
	}

}
