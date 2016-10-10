package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var (
		mealCost   float64
		tipPercent int
		taxPercent int
	)
	fmt.Scanf("%f\n%d\n%d\n", &mealCost, &tipPercent, &taxPercent)
	tip := mealCost * float64(tipPercent) / 100
	tax := mealCost * float64(taxPercent) / 100
	totalCost := mealCost + tip + tax
	fmt.Printf("The total meal cost is %v dollars.\n", int(totalCost))
}
