package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

var testInput1Answer = 159

var testInput2Answer = 135

func main() {
	inputA := formatInput("testInput1A")
	inputB := formatInput("testInput1B")
	partOne(inputA, inputB)
}

// copied from https://github.com/brianstarke/aoc-2019/blob/master/day2-2/main.go#L31
func formatInput(fileName string) []string {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(b), ",")
}

// Part 1 Functions
func partOne(inputA []string, inputB []string) {
	// get wire paths
	wirePathA := getPaths(inputA)
	wirePathB := getPaths(inputB)

}

// WirePath is the path of a wire
type WirePath struct {
	Ports []Port
}

// Port is a Port
type Port struct {
	Column int
	Row    int
}

func getPaths(input []string) WirePath {
	currentPort := Port{
		Column: 0,
		Row:    0,
	}

	wirePath := WirePath{
		Ports: []Port{
			currentPort,
		},
	}

	for _, instruction := range input {
		newPorts, nextPort := findNextPorts(currentPort, instruction)

		for _, p := range newPorts {
			wirePath.Ports = append(wirePath.Ports, p)
		}
		currentPort = nextPort
	}

	return wirePath
}

func findNextPorts(currentPort Port, instructions string) ([]Port, Port) {
	var ports []Port
	var nextPort Port

	split := strings.SplitN(instructions, "", 2)
	direction := split[0]
	distance, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}

	// add all ports inbetween current and next port
	switch direction {
	case "R":
		currentColumn := currentPort.Column
		for i := 0; i < distance+1; i++ {
			n := Port{
				Row:    currentPort.Row,
				Column: currentColumn + 1,
			}
			ports = append(ports, n)
		}

		nextPort.Row = currentPort.Row
		nextPort.Column = currentPort.Column + distance
	case "L":
		currentColumn := currentPort.Column
		for i := 0; i < distance+1; i++ {
			n := Port{
				Row:    currentPort.Row,
				Column: currentColumn - 1,
			}
			ports = append(ports, n)
		}

		nextPort.Row = currentPort.Row
		nextPort.Column = currentPort.Column - distance
	case "U":
		currentRow := currentPort.Row
		for i := 0; i < distance+1; i++ {
			n := Port{
				Row:    currentRow - 1,
				Column: currentPort.Column,
			}
			ports = append(ports, n)
		}

		nextPort.Row = currentPort.Row - distance
		nextPort.Column = currentPort.Column
	case "D":
		currentRow := currentPort.Row
		for i := 0; i < distance+1; i++ {
			n := Port{
				Row:    currentRow + 1,
				Column: currentPort.Column,
			}
			ports = append(ports, n)
		}

		nextPort.Row = currentPort.Row + distance
		nextPort.Column = currentPort.Column
	default:
		panic("unknown direction")
	}

	return ports, nextPort
}
