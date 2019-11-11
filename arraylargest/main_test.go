package main

import (
	"fmt"
	"testing"
)

func Test_getIndicesOfLargest(t *testing.T) {
	input := []int{1, 7, 2, 3, 8, 4, 5, 9, 6}
	num := 4

	output := getIndicesOfLargest(input, num)

	fmt.Printf("output = %+v\n", output)
}

func Test_swap(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// expected := []int{7, 6, 5, 4, 3, 2, 1}

	fmt.Println("input (before swap) = ", input)
	swap(input)
	fmt.Println("input (after swap) = ", input)
	// fmt.Println("expected = ", expected)

}
