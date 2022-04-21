package main

import (
	"math/rand"
	"testing"
	"time"
)

type Parameter struct {
	weights     []int
	values      []int
	bagCapacity int
}

func GenRandomCases() []Parameter {
	rand.Seed(time.Now().UnixMicro())

	size := 0
	for size == 0 {
		size = rand.Intn(100)
	}

	count := 0
	for count == 0 {
		count = rand.Intn(100)
	}

	cases := make([]Parameter, size)
	for r := 0; r < size; r++ {
		capacity := 0
		for capacity == 0 {
			capacity = rand.Intn(10)
		}

		cases[r].bagCapacity = capacity
		cases[r].weights = make([]int, count)
		cases[r].values = make([]int, count)
		for c := 0; c < count; c++ {
			cases[r].weights[c] = rand.Intn(20)
			cases[r].values[c] = rand.Intn(20)
		}
	}

	return cases
}

func TestKnapsack(t *testing.T) {

	cases := []struct {
		Name       string
		ParamsList []Parameter
	}{
		{
			"ok", []Parameter{
				{
					weights:     []int{10, 20, 30},
					values:      []int{60, 100, 120},
					bagCapacity: 51,
				},
			},
		},
		{
			"random", GenRandomCases(),
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			for _, p := range c.ParamsList {
				expected, got := Knapsack(p.weights, p.values, p.bagCapacity)
				if expected != got {
					t.Errorf("expected: %v, got: %v, case: (weights=%v, values=%v, bagCapacity=%v)", expected, got, p.weights, p.values, p.bagCapacity)
				} else {
					t.Logf("accepted: %v, case: (weights=%v, values=%v, bagCapacity=%v)", expected, p.weights, p.values, p.bagCapacity)
				}
			}
		})
	}
}
