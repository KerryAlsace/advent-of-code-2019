package main

import (
	"fmt"
	"io/ioutil"
	"math"
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
func formatInput(n string) []string {
	fileName := fmt.Sprintf("%s.txt", n)
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

	// fmt.Printf("WirePathA: %v\n", wirePathA)
	// fmt.Printf("WirePathB: %v\n", wirePathB)

	// figure out where they first cross
	i := findFirstIntersection(wirePathA, wirePathB)

	// calculate manhattan distance
	d := calculateManhattanDistance(i)

	fmt.Printf("Part 1 Answer: %v\n", d)
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

	fmt.Printf("Intersecting Ports: %v\n", intersectingPorts)

	var shortestDistance int
	var portOfShortestDistance Port
	for i, port := range intersectingPorts {
		d := calculateManhattanDistance(port)
		if i == 0 {
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
	fmt.Printf("Instructions: %v\n", instructions)
	var ports []Port
	var nextPort Port

	split := strings.SplitN(instructions, "", 2)
	direction := split[0]
	distance, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("distance: %v\n", distance)

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

	fmt.Printf("Ports: %v\n", ports)
	fmt.Printf("NextPort: %v\n", nextPort)

	return ports, nextPort
}
