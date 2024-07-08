package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

func main() {
	res := plus(3, 4)
	fmt.Println(res)

	res1 := plusPlus(3, 4, 5)
	fmt.Println(res1)

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

}
