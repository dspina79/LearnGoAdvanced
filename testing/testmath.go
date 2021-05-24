package main

import (
	"fmt"
	"testing"
)

func add(x, y int) int {
	return x + y
}

func TestAddBasic(t *testing.T) {
	result := add(3, 4)

	if result != 7 {
		t.Errorf("The answer should have been 7 but was %d", result)
	}
}

func TestAddTable(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6}
	for _, n := range numbers {
		testname := "Add_" + fmt.Sprint(n)
		t.Run(testname, func(t *testing.T) {
			ans := add(n, n)
			if ans != (n * 2) {
				t.Errorf("%d + %d should equal %d but we got %d.", n, n, (n * 2), ans)
			}
		})
	}
}
