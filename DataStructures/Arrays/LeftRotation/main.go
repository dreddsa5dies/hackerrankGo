package main

import (
	"fmt"
	"strconv"
)

func main() {
	// stdin: uint64 - for all testcase
	// get stdin
	var (
		n, a, d uint64
	)

	fmt.Scanf("%d %d", &n, &d)

	arr := []uint64{}

	var i uint64

	for i = 0; i < n; i++ {
		fmt.Scan(&a)
		arr = append(arr, a)
	}

	// its algorithm
	for i = 0; i < d; i++ {
		arr = append(arr, arr[i])
	}

	// copy arr
	tempArr := make([]uint64, len(arr)-int(d))
	copy(tempArr, arr[d:])

	pP(tempArr)

}

func pP(arr []uint64) {
	// stdout
	str := ""

	for k, v := range arr {
		if k == len(arr)-1 {
			s := strconv.FormatUint(v, 16)
			str += s
		} else {
			s := strconv.FormatUint(v, 16)
			str += s + " "
		}
	}

	fmt.Println(str)
}
