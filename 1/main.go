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
	partTwo(i)
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

	for _, m := range i {
		// convert to int
		mass, err := strconv.Atoi(m)
		if err != nil {
			panic(err)
		}

		totalFuelRequired += calculateModuleMass(mass)
	}

	fmt.Printf("Part 1 Answer: %v\n", totalFuelRequired)
}

func calculateModuleMass(mass int) int {
	// divide by 3
	n := mass / 3

	// round down to nearest whole number
	// Note: dividing by an int in Go makes this happen automatically, don't need to code it

	// subtract 2
	return n - 2
}

func partTwo(i []string) {
	totalFuelRequired := 0

	for _, m := range i {
		// convert to int
		mass, err := strconv.Atoi(m)
		if err != nil {
			panic(err)
		}

		totalFuelRequired += calculateTotalModuleMass(mass)
	}

	fmt.Printf("Part 2 Answer: %v\n", totalFuelRequired)
}

func calculateTotalModuleMass(mass int) int {
	fuel := calculateModuleMass(mass)

	if fuel < 1 {
		return 0
	}

	return fuel + calculateTotalModuleMass(fuel)
}
