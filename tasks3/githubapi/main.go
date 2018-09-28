package main

import "fmt"

func main() {
	nLines := githubapi()

	fmt.Println("Number of lines in C++ code: ", nLines)
}
