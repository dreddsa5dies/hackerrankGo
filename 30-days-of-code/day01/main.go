package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var i uint32 = 4
	var d float32 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewReader(os.Stdin)

	// Declare second integer, double, and String variables.
	var (
		i2 uint32
		d2 float32
	)
	// Read and save an integer, double, and String to your variables.
	fmt.Scanf("%d\n%f\n", &i2, &d2)
	s2, _ := scanner.ReadString('\n')
	// Print the sum of both integer variables on a new line.
	i = i + i2
	fmt.Printf("%v\n", i)
	// Print the sum of the double variables on a new line.
	d = d + d2
	fmt.Printf("%1.1f\n", d)
	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	s = s + s2
	fmt.Printf("%s\n", s)
}
