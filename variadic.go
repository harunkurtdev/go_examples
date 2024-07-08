package main

import "fmt"

func sums(nums ...int) int {
	fmt.Println(nums, "")
	total := 0

	for _, num := range nums {
		total += num
	}

	return total

}
func main() {
	fmt.Println(sums(1, 2, 3))
	fmt.Println(sums(([]int{1, 2, 3})...))
}
