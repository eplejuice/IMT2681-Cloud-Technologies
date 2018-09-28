package main

import "testing"

func Test_swap(t *testing.T) {

	a := 10
	b := 20

	swap(&a, &b)

	if a != 20 || b != 10 {
		t.Error("The swap did not work")
	}
}

func Test_hex(t *testing.T) {
	testCases := map[string]string{
		"a":      "10",
		"aa":     "1010",
		"abc":    "101112",
		"aabbcc": "101011111212",
	}
	for input, expected := range testCases {
		actual := hex(input)

		if actual != expected {
			t.Error("Failure, Expected ", expected, " but got ", actual)
		}
	}
}

func main() {
}
