package main

func main() {

}

func reverse(input []int) {
	for i := 0; i <= len(input)/2; i++ {
		input[i], input[(len(input)-1)-i] = input[(len(input)-1)-i], input[i]
	}
}
