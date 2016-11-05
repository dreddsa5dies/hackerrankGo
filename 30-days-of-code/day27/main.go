package main

import "fmt"

func main() {
	fmt.Println("5")

	fmt.Println("4 3") //YES
	fmt.Println("-1 3 0 2")

	fmt.Println("5 1") //NO
	fmt.Println("0 -1 2 1 4")

	fmt.Println("8 3") //YES
	fmt.Println("-3 0 2 3 4 5 6 7")

	fmt.Println("7 5") //NO
	fmt.Println("0 -1 -3 -4 -2 3 4 ")

	fmt.Println("6 4") //YES
	fmt.Println("-1 0 4 2 9 10")
}
