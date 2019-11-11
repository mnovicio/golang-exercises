package main

import (
	"fmt"
	"testing"
)

func Test_reverse(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// expected := []int{7, 6, 5, 4, 3, 2, 1}

	fmt.Println("input (before reverse) = ", input)
	reverse(input)
	fmt.Println("input (after reverse) = ", input)
	// fmt.Println("expected = ", expected)

}
