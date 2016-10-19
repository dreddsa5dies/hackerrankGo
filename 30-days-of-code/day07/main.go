package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	var (
		n int
	)
	fmt.Scanf("%d\n", &n)

	nums := make([]int, 0)
	var i int
	for m := 0; m < n; m++ {
		_, err := fmt.Scanf("%v", &i)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		nums = append(nums, i)
	}

	for k := len(nums) - 1; k >= 0; k-- {
		fmt.Printf("%v ", nums[k])
	}
	fmt.Println()
}
