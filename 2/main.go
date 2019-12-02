package main

import "fmt"

var officialInput = []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 10, 1, 19, 2, 9, 19, 23, 2, 13, 23, 27, 1, 6, 27, 31, 2, 6, 31, 35, 2, 13, 35, 39, 1, 39, 10, 43, 2, 43, 13, 47, 1, 9, 47, 51, 1, 51, 13, 55, 1, 55, 13, 59, 2, 59, 13, 63, 1, 63, 6, 67, 2, 6, 67, 71, 1, 5, 71, 75, 2, 6, 75, 79, 1, 5, 79, 83, 2, 83, 6, 87, 1, 5, 87, 91, 1, 6, 91, 95, 2, 95, 6, 99, 1, 5, 99, 103, 1, 6, 103, 107, 1, 107, 2, 111, 1, 111, 5, 0, 99, 2, 14, 0, 0}

var testInput = []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50} // 3500,9,10,70,2,3,11,0,99,30,40,50
var testInput2 = []int{1, 0, 0, 0, 99}                          // 2,0,0,0,99
var testInput3 = []int{2, 3, 0, 3, 99}                          // 2,3,0,6,99
var testInput4 = []int{2, 4, 4, 5, 99, 0}                       // 2,4,4,5,99,9801
var testInput5 = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}             // 30,1,1,4,2,5,6,0,99

func main() {
	input := testInput
	cont, set, lastIndex := getNextSet(-1, input)

	for cont {
		makeCalcs(set, input)
		cont, set, lastIndex = getNextSet(lastIndex, input)
	}

	fmt.Println(input)
}

// returns new input with edits
func makeCalcs(set []int, input []int) []int {
	command := set[0]
	indexA := set[1]
	indexB := set[2]
	indexC := set[3]

	switch command {
	case 1:
		input[indexC] = input[indexA] + input[indexB]
	case 2:
		input[indexC] = input[indexA] * input[indexB]
	case 99:
		fmt.Println("end")
	default:
		panic("whaaat")
	}

	return input
}

// returns true if program should halt, otherwise returns false, the next 4 ints, and the index of the last int
func getNextSet(lastIndex int, input []int) (bool, []int, int) {
	firstIndex := lastIndex + 1

	if input[firstIndex] == 99 {
		return false, nil, lastIndex
	}

	return true, []int{input[firstIndex], input[firstIndex+1], input[firstIndex+2], input[firstIndex+3]}, firstIndex + 3
}
