package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/eplejuice/IMT2681-Cloud-Technologies/LAB1/modules"
)

func main() {
	// Takes input from user
	arg, _ := strconv.Atoi(os.Args[1])

	// Task 1
	fmt.Println(modules.Fibonacciloop(arg))
	fmt.Println(modules.Fibonaccirecursive(arg))

	// Task 2

	fmt.Println(modules.Factorialloop(arg))
	fmt.Println(modules.Factorialrecursive(arg))

	// Task 4
	fmt.Println(zero()(1, 2))

	// Task 5
	agey := person{age: 10}
	fmt.Println(agey.age)
	p := &agey
	p.addten()
	fmt.Println(p.age)
	add10(p)
	fmt.Println(p.age)
}
