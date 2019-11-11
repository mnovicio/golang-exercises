package main

import (
	"fmt"
)

func main() {

}

func insertAsc(dst []int, a int) []int {
	fmt.Printf("input = %+v, insert %d\n", dst, a)
	if len(dst) == 0 {
		dst = append(dst, a)
	}
	for idx, val := range dst {
		if val > a {
			dst = append(dst[:idx], append([]int{a}, dst[idx:]...)...)
			break
		}
	}

	return dst
}
