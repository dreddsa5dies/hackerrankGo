package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	var (
		n int
	)

	fmt.Scan(&n)
	scanner := bufio.NewScanner(os.Stdin)
	strArr := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		str := strings.Split(scanner.Text(), " ")
		matched, _ := regexp.MatchString(".+@gmail\\.com$", str[1])
		if matched {
			strArr = append(strArr, str[0])
		}
	}

	sort.Strings(strArr)
	for i := n; i < len(strArr); i++ {
		fmt.Println(strArr[i])
	}
}
