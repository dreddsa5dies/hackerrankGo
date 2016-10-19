package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d\n", &n)

	if n%2 != 0 {
		fmt.Println("Weird")
	} else {
		switch {
		case n >= 2 && n <= 5:
			fmt.Println("Not Weird")
		case n >= 6 && n <= 20:
			fmt.Println("Weird")
		case n > 20:
			fmt.Println("Not Weird")
		default:
			fmt.Println("Weird")
		}
	}
}
