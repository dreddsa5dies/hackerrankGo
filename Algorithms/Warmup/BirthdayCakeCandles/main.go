package main

import (
	"fmt"
	"os"
)

func main() {
	// stdin: uint - for all testcase
	var (
		n, a uint
	)

	fmt.Scanf("%d", &n)
	arr := []uint{}
	var i uint
	for i = 0; i < n; i++ {
		fmt.Scan(&a)
		arr = append(arr, a)
	}

	// uncknown, but println testcase check stderr!
	fmt.Fprintf(os.Stdout, "%d", birthdayCakeCandles(n, arr))
}

func birthdayCakeCandles(n uint, arr []uint) uint {
	max := arr[0]
	var num uint
	num = 1
	var i uint
	for i = 1; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
			num = 1
		} else if arr[i] == max {
			num++
		}
	}
	return num
}
