package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
		fmt.Printf("%d\n", i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
