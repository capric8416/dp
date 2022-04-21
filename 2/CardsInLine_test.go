package main

import (
	"math/rand"
	"testing"
	"time"
)

type RandomCases struct {
	Name       string
	ParamsList [][]int
}

func GenRandomCases() RandomCases {
	rand.Seed(time.Now().UnixMicro())

	size := 0
	for size == 0 {
		size = rand.Intn(20)
	}

	count := 0
	for count == 0 {
		count = rand.Intn(100)
	}

	params := make([][]int, count)
	for r := 0; r < count; r++ {
		params[r] = make([]int, size)
		for c := 0; c < size; c++ {
			params[r][c] = rand.Intn(100)
		}
	}

	return RandomCases{"random", params}
}

func TestCardsInLine(t *testing.T) {
	cases := []struct {
		Name       string
		ParamsList [][]int
	}{
		{
			"ok", [][]int{
				{5, 20, 15, 10, 20},
				{5, 7, 4, 5, 8, 1, 6, 0, 3, 4, 6, 1, 7},
			},
		},
		{
			"nil", [][]int{
				nil,
			},
		},
		{
			"empty", [][]int{
				{},
			},
		},
		GenRandomCases(),
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			for _, p := range c.ParamsList {
				expected, got := CardsInLine(p)
				if got != expected {
					t.Errorf("expected: %v, got: %v, case: (arr=%v)", expected, got, p)
				} else {
					t.Logf("accepted: %v case: (arr=%v)", expected, p)
				}
			}
		})
	}
}
