package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	exRat := exRate()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Select valuta:  ")
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	for i, j := range exRat {
		if i == s {
			fmt.Println(i, j)
		}
	}
}
