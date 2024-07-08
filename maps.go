package main

import "fmt"

func main() {
	var s []string
	fmt.Println(s)

	s = make([]string, 3)
	fmt.Println(s, cap(s))

}
