package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	var (
		n   int
		sum int64
	)
	fmt.Scanf("%d", &n)

	nums := make([]int64, 0)
	var i int64
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

	for k := 0; k < len(nums); k++ {
		sum = sum + nums[k]
	}

	fmt.Println(sum)
}
