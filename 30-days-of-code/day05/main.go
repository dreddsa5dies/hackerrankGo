package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int

	fmt.Scanf("%d", &n)

	for i := 0; i < 10; i++ {
		fmt.Printf("%v x %v = %v\n", n, i+1, n*(i+1))
	}
}
