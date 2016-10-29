package main

import "fmt"
import "strings"

func main() {
	var (
		n  int
		s  string
		ss []string
	)

	fmt.Scanf("%d\n", &n)

	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		ss = append(ss, s)
	}

	for i := 0; i < n; i++ {
		if nextPermutation(ss[i]) {
			// fmt.Println(reverse(ss[i]))
			fmt.Println(reverse(ss[i]))
		} else {
			fmt.Println("no answer")
		}
	}
}

func nextPermutation(str string) bool {
	// Find non-increasing suffix
	i := len(str) - 1
	for i > 0 && str[i-1] >= str[i] {
		i--
	}

	if i <= 0 {
		return false
	}

	return true
}

func reverse(str string) string {
	i := len(str) - 1
	for i > 0 && str[i-1] >= str[i] {
		i--
	}

	// Find successor to pivot
	j := len(str) - 1
	for str[j] <= str[i-1] {
		j--
	}

	strArr := []string{}
	for l := 0; l < len(str); l++ {
		strArr = append(strArr, string(str[l]))
	}

	strArr[i-1], strArr[j] = strArr[j], strArr[i-1]

	// Reverse suffix
	strArrRes := []string{}
	for k := len(strArr) - 1; k > i-1; k-- {
		strArrRes = append(strArrRes, strArr[k])
	}

	copy(strArr[i:], strArrRes)

	return strings.Join(strArr, "")
}
