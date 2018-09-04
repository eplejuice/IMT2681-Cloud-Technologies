package modules

import "fmt"

func Factorialloop(arg int) int {

	if arg <= 1 {
		return 1
	} else {
		var temp int = 1
		for i := arg; i > 1; i-- {
			temp *= i
			fmt.Println(temp)
		}
		return temp
	}
}

func Factorialrecursive(arg int) int {
	if arg <= 1 {
		return 1
	} else {
		return (arg * factorialrecursive(arg-1))
	}
}
