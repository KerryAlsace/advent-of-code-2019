package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fileName := "input"
	intInput := 1
	partOne(fileName, intInput)
}

func formatInput(n string) []int {
	fileName := fmt.Sprintf("%s.txt", n)
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	s := strings.Split(string(b), ", ")

	m := make([]int, len(s))

	for i, r := range s {
		x, err := strconv.Atoi(r)
		if err != nil {
			panic(err)
		}

		m[i] = x
	}

	return m
}

func partOne(fileName string, intInput int) {
	input := formatInput(fileName)

	fmt.Printf("Part 1 Answer: %v\n", compute(input, intInput))
}

// Instruction is an instruction
type Instruction struct {
	Opcode     Opcode
	Parameters []Parameter
}

// Parameter is a parameter of the instruction
type Parameter struct {
	Value        int
	PositionMode PositionMode
}

// Opcode is the code that indicates what do to do the input
type Opcode int

// The possible opcodes
const (
	OpcodeOne        Opcode = 01
	OpcodeTwo               = 02
	OpcodeThree             = 03
	OpcodeFour              = 04
	OpcodeNinetyNine        = 99
)

// PositionMode indicates how to use the param
type PositionMode int

// The possible position modes
const (
	PositionModeZero PositionMode = 0
	PositionModeOne               = 1
)

func compute(input []int, intInput int) int {
	cont, opcode, params, lastIndex, intOutput := getNextSet(-1, input, intInput)

	for cont {
		makeCalcs(opcode, params, input, intOutput)
		cont, opcode, params, lastIndex, intOutput = getNextSet(lastIndex, input, intOutput)
	}

	return intOutput
}

func getNextSet(lastIndex int, input []int, intInput int) (cont bool, opcode Opcode, params []int, newLastIndex int, output int) {
	cont = true
	firstIndex := lastIndex + 1

	opcode, pms, length := getOpcodeAndPositionModes(input[firstIndex])
	params, newLastIndex = getParams(opcode, firstIndex, input)

	if opcode == OpcodeNinetyNine {
		cont = false
		output = intInput
	}


	return cont, opcode, params, newLastIndex, output
}

func getOpcodePositionModesAndLength(codes int) (opcode Opcode, pms int, length int) {
	codesString := strconv.Itoa(codes)

	opcodeString := codeString[len(codeString)-2:]
	opcodeInt, err := strconv.Atoi(opcodeString)
	if err != nil {
		panic(err)
	}

	opcode = getOpcode(opcodeInt)
	positionModesString := codeString[0 : len(codeString)-2]

	pms, err = strconv.Atoi(positionModesString)
	if err != nil {
		panic(err)
	}

	length = len(positionModesString)

	return opcode, pms, length
}

func getParams(opcode Opcode, firstIndex int, input []int) ([]int, int) {
	switch opcode {
	case OpcodeOne:
	case OpcodeTwo:
		return []int{input[firstIndex], input[firstIndex+1], input[firstIndex+2], input[firstIndex+3]}, firstIndex+3
	case OpcodeThree:
		return []int{}
	case OpcodeFour:
		return []int{}
	case OpcodeNinetyNine:
		return []int{}, firstIndex - 1
	default:
		panic(fmt.Sprintf("Unknown opcode: %v", opcode))
	}
}

func getOpcode(opcode int) Opcode {
	switch opcode {
	case 01:
		return OpcodeOne
	case 02:
		return OpcodeTwo
	case 03:
		return OpcodeThree
	case 04:
		return OpcodeFour
	case 99:
		return OpcodeNinetyNine
	default:
		panic(fmt.Sprintf("Unknown opcode: %v", opcode))
	}
}

func enactOpcode(opcode Opcode, input []int, set) {
	switch opcode {
	case OpcodeOne:
		opcodeOne()
	}
}

func opcodeOne(trigger int, input []int) {

}

func opcodeTwo(trigger int, input []int) {

}

func opcodeThree(trigger int, input []int, index int) {
	input[index] = input
}

func opcodeFour(input []int, index int) int {
	return input[index]
}

func opcodeNinetyNine() {

}

func getPositionModesOpcodesAndParameters(input []int) (int, []Param) {
	codes := input[0]
	params := input[1:]

	codeString := strconv.Itoa(codes)

	opcodeString := codeString[len(codeString)-2:]
	opCodeInt, err := strconv.Atoi(opcodeString)
	if err != nil {
		panic(err)
	}

	fmt.Printf("op %v\n", opcodeString)

	positionModesString := codeString[0 : len(codeString)-2]

	fmt.Printf("pms %v\n", positionModesString)

	// fill in omitted positionModes
	l := len(params) - len(positionModesString)
	if l != 0 {
		for r := 0; r < l+1; r++ {
			positionModesString = strings.Join([]string{"0", positionModesString}, "")
		}
	}

	var formattedParams []Parameter
	formattedParams = make([]Parameter, len(params), len(params))

	for i, p := range params {
		positionModeIdx := (len(positionModesString) - i) - 1
		positionMode := positionModesString[positionModeIdx : positionModeIdx+1]

		positionModeInt, err := strconv.Atoi(positionMode)
		if err != nil {
			panic(err)
		}

		f := Parameter{
			Parameter:    formatParameterByMode(positionModeInt, p, thing),
			PositionMode: positionModeInt,
		}

		formattedParams[i] = f
	}

	return opCodeInt, formattedParams
}

func formatParameterByMode(positionMode int, parameter int, input []int) int {
	switch positionMode {
	case 0:
		return input[parameter]
	case 1:
		return parameter
	default:
		panic(fmt.Sprintf("unknown position mode %v\n", positionMode))
	}
}
