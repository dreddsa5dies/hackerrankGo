package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	n := 0
	fmt.Scanln(&n)
	arr := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := strings.Split(scanner.Text(), " ")
	for i := 0; i < len(str); i++ {
		ls, _ := strconv.Atoi(str[i])
		arr = append(arr, ls)
	}

	for i := 1; i < len(arr); i++ {
		for j := i; j != 0 && arr[j] < arr[j-1]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			fmt.Println(arr)
		}
	}
}
