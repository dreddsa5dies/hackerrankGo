package main

import "fmt"

func main() {
	var x int

	fmt.Scanf("%d", &x)
	x = function(x)
	fmt.Println(x)
}

func function(x int) int {
	if x == 163 {
		return 9
	} else {
		return x - 11
	}
}
