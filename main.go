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
	//fmt.Println(score)
	//fmt.Println(s)
	if s[0] == 10 { // strike
		if len(s) == 3 { // tenth frame strike
			return score + Sum(s)
		}
		return Bowl(s[1:len(s)], score+Sum(s[0:3]))
	} else if Sum(s[0:2]) == 10 { // spare
		if len(s) == 3 { // tenth frame spare
			return score + Sum(s)
		}
		return Bowl(s[2:len(s)], score+Sum(s[0:3]))
	} else {
		if len(s) == 2 { // tenth open frame
			return score + Sum(s)
		}
		return Bowl(s[2:len(s)], score+Sum(s[0:2]))
	}
}
