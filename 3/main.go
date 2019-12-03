package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var part1TestInput1Answer = 159
var part1TestInput2Answer = 135

var part2TestInput1Answer = 610
var part2TestInput2Answer = 410

func main() {
	// Part 1
	fileName := "testInput2"
	partOne(fileName)

	// Part 2

}

// copied from https://github.com/brianstarke/aoc-2019/blob/master/day2-2/main.go#L31
func formatInput(n string) []string {
	fileName := fmt.Sprintf("%s.txt", n)
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(b), ",")
}

// Part 1 Functions
func partOne(fileName string) {
	inputA := formatInput(fmt.Sprintf("%sA", fileName))
	inputB := formatInput(fmt.Sprintf("%sB", fileName))

	d := getDistance(inputA, inputB)
	fmt.Printf("Part 1 Answer: %v\n", d)

	if fileName == "testInput1" {
		if part1TestInput1Answer != d {
			fmt.Printf("Part 1 is incorrect, got: %v, want: %v\n", d, part1TestInput1Answer)
		}

		fmt.Printf("Part 1 is correct\n")
	}

	if fileName == "testInput2" {
		if part1TestInput2Answer != d {
			fmt.Printf("Part 1 is incorrect, got: %v, want: %v\n", d, part1TestInput2Answer)
		}

		fmt.Printf("Part 1 is correct\n")
	}
}

func getDistance(inputA []string, inputB []string) int {
	// get wire paths
	wirePathA := getPaths(inputA)
	wirePathB := getPaths(inputB)

	// figure out where they first cross
	i := findFirstIntersection(wirePathA, wirePathB)

	// return calculated manhattan distance
	return calculateManhattanDistance(i)
}

// this could be optimized
func findFirstIntersection(pathA WirePath, pathB WirePath) Port {
	var intersectingPorts []Port
	for _, portA := range pathA.Ports {
		for _, portB := range pathB.Ports {
			if portA == portB {
				intersectingPorts = append(intersectingPorts, portA)
			}
		}
	}

	var shortestDistance int
	var portOfShortestDistance Port
	for _, port := range intersectingPorts {
		d := calculateManhattanDistance(port)
		if port.Row == 0 && port.Column == 0 {
			continue
		}
		if shortestDistance == 0 {
			portOfShortestDistance = port
			shortestDistance = d
		}

		if d < shortestDistance {
			portOfShortestDistance = port
			shortestDistance = d
		}
	}

	return portOfShortestDistance
}

func calculateManhattanDistance(i Port) int {
	c := math.Abs(float64(i.Column))
	r := math.Abs(float64(i.Row))

	return int(c) + int(r)
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
		for i := 1; i < distance+2; i++ {
			n := Port{
				Row:    currentPort.Row,
				Column: currentColumn + i,
			}
			ports = append(ports, n)
		}

		nextPort.Row = currentPort.Row
		nextPort.Column = currentPort.Column + distance
	case "L":
		currentColumn := currentPort.Column
		for i := 1; i < distance+2; i++ {
			n := Port{
				Row:    currentPort.Row,
				Column: currentColumn - i,
			}
			ports = append(ports, n)
		}

		nextPort.Row = currentPort.Row
		nextPort.Column = currentPort.Column - distance
	case "U":
		currentRow := currentPort.Row
		for i := 1; i < distance+2; i++ {
			n := Port{
				Row:    currentRow - i,
				Column: currentPort.Column,
			}
			ports = append(ports, n)
		}

		nextPort.Row = currentPort.Row - distance
		nextPort.Column = currentPort.Column
	case "D":
		currentRow := currentPort.Row
		for i := 1; i < distance+2; i++ {
			n := Port{
				Row:    currentRow + i,
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
