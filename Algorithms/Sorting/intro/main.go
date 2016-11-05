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
	v := 0
	fmt.Scanln(&n)
	fmt.Scanln(&v)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := strings.Split(scanner.Text(), " ")
	for i := 0; i < len(str); i++ {
		ls, _ := strconv.Atoi(str[i])
		if n == ls {
			fmt.Println(i)
		}
	}
}
