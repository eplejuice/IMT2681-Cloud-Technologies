package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter gender (male/female): ")
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	fmt.Print("Enter country: ")
	c, _ := reader.ReadString('\n')
	c = strings.TrimSpace(c)

	fmt.Print("Enter date of birth (YYYY-MM-DD): ")
	d, _ := reader.ReadString('\n')
	d = strings.TrimSpace(d)

	fmt.Println(s, c, d)

	le := getLifeEx(s, c, d)

	fmt.Println("Got life expectancy ", le)
}
