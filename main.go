package main

import (
	"golang.org/x/exp/constraints"
)

// TODO: add tests,
// Add govet, golint, go ?? good practice code hygiene

type rolls []int

type Number interface {
	// Using "constraints" concisely enables all numeric types.
	constraints.Integer | constraints.Float
}

func Sum[T Number](s []T) T {
	var sum T
	for _, num := range s {
		sum = sum + num
	}
	return sum
}

func Bowl(s rolls, score int) int {
	if len(s) <= 3 {
		return score + Sum(s)
	}

	if s[0] == 10 { // strike
		return Bowl(s[1:len(s)], score+Sum(s[0:3]))
	} else if s[0]+s[1] == 10 { // spare
		return Bowl(s[2:len(s)], score+Sum(s[0:3]))
	} else { // open frame
		return Bowl(s[2:len(s)], score+Sum(s[0:2]))
	}
}
