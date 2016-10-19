package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	var (
		n, number int
		s         string
		nums      []string
	)

	k := make(map[string]int)

	fmt.Scanf("%d", &n)

	for m := 0; m < n; m++ {
		_, err := fmt.Scanf("%s %d\n", &s, &number)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		k[s] = number
	}

	for m := 0; m < n; m++ {
		_, err := fmt.Scanf("%s\n", &s)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		nums = append(nums, s)
	}

	for i := 0; i < len(nums); i++ {
		_, prs := k[nums[i]]
		if prs {
			fmt.Printf("%v=%v\n", nums[i], k[nums[i]])
		} else {
			fmt.Println("Not found")
		}
	}
}
