package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
		break
	case 2:
		fmt.Println("two")
		break
	default:
		fmt.Println("tree")
		break
	}

	switch time.Now().Weekday() {
	case time.Sunday:
		fmt.Println(time.Now().Weekday())
	}

	whatIam := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println(i)
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatIam(5)

}
