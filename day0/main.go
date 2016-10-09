package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	s = scan()
	fmt.Println("Hello, World.")
	fmt.Println(s)
}

func scan() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}
