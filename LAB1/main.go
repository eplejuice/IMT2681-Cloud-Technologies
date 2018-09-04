package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg, _ := strconv.Atoi(os.Args[1])
	fmt.Println(Fibonacciloop(arg))
	fmt.Println(Fibonaccirecursive(arg))
	fmt.Println(Factorialloop(arg))
	fmt.Println(Factorialrecursive(arg))
}
