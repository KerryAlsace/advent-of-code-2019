package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	i := getInput()
	partOne(i)
}

func getInput() []string {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(b), "\n")
}

func partOne(i []string) {
	totalFuelRequired := 0

	for _, module := range i {
		totalFuelRequired += calculateModuleMass(module)
	}

	fmt.Printf("Part 1 Answer: %v\n", totalFuelRequired)
}

func calculateModuleMass(mass string) int {
	// convert to int
	m, err := strconv.Atoi(mass)
	if err != nil {
		panic(err)
	}

	// divide by 3
	n := m / 3

	// round down to nearest whole number
	// Note: dividing by an int in Go makes this happen automatically, don't need to code it

	// subtract 2
	return n - 2
}
