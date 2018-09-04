package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/eplejuice/IMT2681-Cloud-Technologies/LAB1/modules"
)

func main() {
	arg, _ := strconv.Atoi(os.Args[1])
	fmt.Println(modules.Fibonacciloop(arg))
	fmt.Println(modules.Fibonaccirecursive(arg))
	fmt.Println(modules.Factorialloop(arg))
	fmt.Println(modules.Factorialrecursive(arg))
}
