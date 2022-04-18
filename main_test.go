package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBowl(t *testing.T) {
	type bowlingScore struct {
		input []int
		want  int
	}

	scores := []bowlingScore{
		// {[]int{0, 0, 0, 0, 0, 0, 0, 0, 0}, 0},
		{[]int{0, 1, 7, 2, 1, 2, 1, 2, 1, 0, 1, 2, 4, 3, 5, 4, 3, 2, 1, 3}, 45},
		{[]int{1, 5, 3, 6, 7, 2, 3, 6, 4, 4, 5, 3, 3, 3, 4, 5, 8, 1, 2, 6}, 81},
		{[]int{10, 3, 6, 7, 2, 3, 6, 4, 4, 5, 3, 3, 3, 4, 5, 8, 1, 2, 6}, 94},
		{[]int{1, 9, 3, 6, 7, 2, 3, 6, 4, 4, 5, 3, 3, 3, 4, 5, 8, 1, 2, 6}, 88},
		{[]int{10, 4, 6, 7, 2, 3, 6, 4, 4, 5, 3, 3, 3, 4, 5, 8, 1, 2, 6}, 103},
		{[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, 300},
	}

	for _, bs := range scores {
		assert.Equal(t, bs.want, Bowl(bs.input, 0))
	}
}
