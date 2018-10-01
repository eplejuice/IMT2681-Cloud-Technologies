package main

import (
	"os"
)

func main() {

	arg1 := os.Args[1]
	arg2 := os.Args[2]
	githubapi(arg1, arg2)

}
