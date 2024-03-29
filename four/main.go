package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input = "193651-649729"

var testInput = "234545-234567" // 6 for part 1, 5 for part 2

func main() {
	i := input
	partOne(i)
	partTwo(i)
}

func partOne(i string) {
	fmt.Printf("Part 1 Answer: %v\n", findNumPossible(i))
}

func findNumPossible(i string) int {
	var possiblePasswords int
	rs := strings.Split(i, "-")
	l, u := getUpperAndLowerBounds(rs)

	r := u - l

	for i := 0; i < r; i++ {
		n := l + i

		ns := strconv.Itoa(n)

		if !hasDoubleDigits(ns) {
			continue
		}

		if !hasAscendingDigits(ns) {
			continue
		}

		possiblePasswords++
	}

	return possiblePasswords
}

func getUpperAndLowerBounds(r []string) (int, int) {
	ls := r[0]
	us := r[1]

	l, err := strconv.Atoi(ls)
	if err != nil {
		panic(err)
	}

	u, err := strconv.Atoi(us)
	if err != nil {
		panic(err)
	}

	return l, u
}

func hasDoubleDigits(s string) bool {
	h := false

	for i := 0; i < 5; i++ {
		if s[i] == s[i+1] {
			h = true
		}
	}

	return h
}

func hasAscendingDigits(s string) bool {
	h := true

	for i := 0; i < 5; i++ {
		if s[i] > s[i+1] {
			h = false
		}
	}

	return h
}

func partTwo(i string) {
	fmt.Printf("Part 2 Answer: %v\n", findMoreRestrictedNumPossible(i))
}

func findMoreRestrictedNumPossible(i string) int {
	var possiblePasswords int
	rs := strings.Split(i, "-")
	l, u := getUpperAndLowerBounds(rs)

	r := u - l

	for i := 0; i < r; i++ {
		n := l + i

		ns := strconv.Itoa(n)

		if !hasAscendingDigits(ns) {
			continue
		}

		if !hasIsolatedDoubleDigits(ns) {
			continue
		}

		possiblePasswords++
	}

	return possiblePasswords
}

func hasIsolatedDoubleDigits(s string) bool {
	h := false

	m := make(map[byte]int)

	for i := 0; i < 5; i++ {
		if s[i] == s[i+1] {
			h = true
			m[s[i]] += 2
		}
	}

	if h {
		h = false
		for _, v := range m {
			if v == 2 {
				h = true
			}
		}
	}

	return h
}
