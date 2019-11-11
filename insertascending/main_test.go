package main

import (
	"fmt"
	"testing"
)

func Test_insertAsc(t *testing.T) {
	input := []int{1, 2, 4, 5}
	fmt.Println("input (before insert 3) = ", input)
	input = insertAsc(input, 3)
	fmt.Println("input (after insert 3) = ", input)

}
