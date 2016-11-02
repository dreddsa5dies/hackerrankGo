package main

import "fmt"

func main() {
	var n, v int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &v)
		if isPrime(v) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not prime")
		}
	}
}

func isPrime(num int) bool {
	if num == 1 {
		return false
	}

	if num == 2 {
		return true
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
