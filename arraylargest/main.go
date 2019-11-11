package main

import (
	"sort"
)

func main() {

}

func getIndicesOfLargest(input []int, num int) []int {
	// input := []int{1, 7, 2, 3, 8, 4, 5, 9, 6}
	// num := 4

	type KeyValue struct {
		key   int
		value int
	}

	list := []KeyValue{}
	for idx, value := range input {
		list = append(list, KeyValue{
			key:   idx,
			value: value,
		})
	}

	// fmt.Printf("checkpoint 1 = %+v\n", list)

	sort.Slice(list, func(i, j int) bool {
		return list[i].value > list[j].value
	})

	// fmt.Printf("checkpoint 2 = %+v\n", list)

	output := []int{}
	for _, v := range list[:num] {
		output = append(output, v.key)
	}

	// fmt.Printf("checkpoint 3 = %+v\n", output)

	return output
}
