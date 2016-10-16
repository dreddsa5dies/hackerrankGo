package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		valueAlice, valueBob   string
		allValAlice, allValBob []string
		alicePoint, bobPoint   int
	)
	scanner := bufio.NewScanner(os.Stdin)
	// scanning valueAlice
	scanner.Scan()
	valueAlice = scanner.Text()

	// scanning valueBob
	scanner.Scan()
	valueBob = scanner.Text()

	allValAlice = strings.Split(valueAlice, " ")
	allValBob = strings.Split(valueBob, " ")

	for i := 0; i < len(allValAlice); i++ {
		a, _ := strconv.Atoi(allValAlice[i])
		b, _ := strconv.Atoi(allValBob[i])
		if a > b {
			alicePoint++
		} else if a < b {
			bobPoint++
		}
	}

	fmt.Printf("%v %v", alicePoint, bobPoint)
}
