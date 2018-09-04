package modules

func Fibonacciloop(arg int) int {

	if arg <= 1 {
		return 1
	} else {

		x := 0
		y := 1
		var fib int

		for i := 1; i < arg; i++ {
			fib = y + x
			x = y
			y = fib
		}
		return fib
	}
}

func Fibonaccirecursive(arg int) int {

	if arg <= 2 {
		return 1
	} else {

		return fibonaccirecursive(arg-1) + fibonaccirecursive(arg-2)
	}

}
