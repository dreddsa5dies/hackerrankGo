// example for roman_ragimoff

package main

import "fmt"
import "bufio"
import "os"
import "strconv"

func main() {
	var (
		n         int
		s         string
		phonebook map[string]string = map[string]string{}
	)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		s = scanner.Text()
		scanner.Scan()
		phonebook[s] = scanner.Text()
	}

	for scanner.Scan() {
		s = scanner.Text()
		if v, ok := phonebook[s]; ok {
			fmt.Printf("%s=%s\n", s, v)
		} else {
			fmt.Println("Not found")
		}
	}
}
