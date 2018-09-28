package main

func zero() func(x, y int) int {

	return func(x, y int) int {
		return x + y
	}

}
