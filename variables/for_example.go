package main

import "fmt"

func main() {
	const count int = 5
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	for _, v := range 6 {
		fmt.Println(v)
	}
}
