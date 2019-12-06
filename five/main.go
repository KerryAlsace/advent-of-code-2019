package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fileName := "input"
	partOne(fileName)
}

func partOne(fileName string) {

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

func getNextSet(lastIndex int, input []int) (bool, []int, int) {
	firstIndex := lastIndex + 1

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
	default:
		panic(fmt.Sprintf("Unknown opcode: %v", opcode))
	}
}

func enactOpcode(opcode Opcode) {

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
