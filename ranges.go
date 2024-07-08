package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	fmt.Println(nums)

	sum := 0
	for _, v := range nums {
		sum += v
	}

	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println(i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println(k, v)
	}

	for k := range kvs {
		fmt.Println(k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
