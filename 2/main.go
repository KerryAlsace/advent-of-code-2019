package main

import "fmt"

var officialInput = []int{1, 82, 26, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 10, 1, 19, 2, 9, 19, 23, 2, 13, 23, 27, 1, 6, 27, 31, 2, 6, 31, 35, 2, 13, 35, 39, 1, 39, 10, 43, 2, 43, 13, 47, 1, 9, 47, 51, 1, 51, 13, 55, 1, 55, 13, 59, 2, 59, 13, 63, 1, 63, 6, 67, 2, 6, 67, 71, 1, 5, 71, 75, 2, 6, 75, 79, 1, 5, 79, 83, 2, 83, 6, 87, 1, 5, 87, 91, 1, 6, 91, 95, 2, 95, 6, 99, 1, 5, 99, 103, 1, 6, 103, 107, 1, 107, 2, 111, 1, 111, 5, 0, 99, 2, 14, 0, 0}

var testInput = []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50} // 3500,9,10,70,2,3,11,0,99,30,40,50
var testInput2 = []int{1, 0, 0, 0, 99}                          // 2,0,0,0,99
var testInput3 = []int{2, 3, 0, 3, 99}                          // 2,3,0,6,99
var testInput4 = []int{2, 4, 4, 5, 99, 0}                       // 2,4,4,5,99,9801
var testInput5 = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}             // 30,1,1,4,2,5,6,0,99

var partTwoOutput = 19690720

func main() {
	partOneInput := make([]int, len(officialInput))
	copy(partOneInput, officialInput)
	partOne(partOneInput)

	partTwoInput := make([]int, len(officialInput))
	copy(partTwoInput, officialInput)
	partTwo(partTwoInput)
	bruteForce()
}

func bruteForce(givenInput []int) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			input := make([]int, len(givenInput))
			copy(input, officialInput)

			input[1] = noun
			input[2] = verb

			newInput := compute(input)

			if newInput[0] != partTwoOutput {
				continue
			}

			fmt.Printf("\nPart 2 Answer: %v\n", formatAnswer(noun, verb))
		}
	}
}

// // Functions for Part 2
// func partTwo(input []int) {
// 	lastZeroIndex := findIndexOfLastZero(input)

// 	// cont, set, lastIndex := getPrevSet(lastZeroIndex, input)

// 	// for cont {
// 	// 	makeReverseCalcs(set, input)
// 	// 	cont, set, lastIndex = getPrevSet(lastIndex, input)
// 	// }

// 	array := convertToArray(input)

// }

// func makeReverseCalcs(lastIndex int, array map[int]Thing, needsToEqual int) map[int]Thing {
// 	firstIndex := lastIndex - 3
// 	commandNode := array[firstIndex]
// 	nounNode := array[firstIndex + 1]
// 	verbNode := array[firstIndex + 2]
// 	addressNode := array[firstIndex + 3]

// 	// set the new value for needs to equal
// 	addressNode.new = needsToEqual

// 	switch commandNode.old {
// 	case 1:

// 	case 2:
// 	default:
// 		panic("whoooo")
// 	}
// }

// Thing is a thing
type Thing struct {
	old int
	new int
}

// Holder hodor
type Holder struct {
	old      int
	computed int
	new      int
}

func convertToMap(input []int, computed []int) map[int]Holder {
	newMap := make(map[int]Holder)

	for i := 0; i < len(input); i++ {
		newMap[i] = Holder{
			old:      input[i],
			computed: computed[i],
		}
	}

	return newMap
}

func convertToArray(input []int) map[int]Thing {
	newArray := make(map[int]Thing)
	for i := 0; i < len(input); i++ {
		newArray[i] = Thing{
			old: input[i],
		}
	}

	return newArray
}

// returns the index of the last time the output is put in index 0 before the program ends with 99
func findIndexOfLastZero(input []int) int {
	lastZeroIndex := 1
	for i := 0; input[i] != 99; i += 4 {
		if input[i+3] == 0 {
			lastZeroIndex = i + 3
		}
	}

	return lastZeroIndex
}

// formats the answer as directed
func formatAnswer(noun int, verb int) int {
	return (noun * 100) + verb
}

// Functions for Part 1
func partOne(input []int) {
	newInput := compute(input)

	fmt.Println(newInput[0])
}

func compute(input []int) []int {
	cont, set, lastIndex := getNextSet(-1, input)

	for cont {
		makeCalcs(set, input)
		cont, set, lastIndex = getNextSet(lastIndex, input)
	}

	return input
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

// returns false if program should halt, otherwise returns true, the next 4 ints, and the index of the last int
func getNextSet(lastIndex int, input []int) (bool, []int, int) {
	firstIndex := lastIndex + 1

	if input[firstIndex] == 99 {
		return false, []int{input[lastIndex-3], input[lastIndex-2], input[lastIndex-1]}, lastIndex
	}

	return true, []int{input[firstIndex], input[firstIndex+1], input[firstIndex+2], input[firstIndex+3]}, firstIndex + 3
}
