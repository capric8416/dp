package main

import (
	"math/rand"
	"testing"
	"time"
)

type Parameter struct {
	targetXPos int
	targetYPos int
	steps      int
}

func GenRandomCases() []Parameter {
	rand.Seed(time.Now().UnixMicro())

	count := 0
	for count == 0 {
		count = rand.Intn(20)
	}

	cases := make([]Parameter, count)
	for i := 0; i < count; i++ {
		cases[i] = Parameter{
			targetXPos: rand.Intn(9),
			targetYPos: rand.Intn(10),
			steps:      rand.Intn(6),
		}
	}

	return cases
}

func TestHourseJump(t *testing.T) {
	cases := []struct {
		Name       string
		ParamsList []Parameter
	}{
		{
			"random",
			GenRandomCases(),
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			for _, p := range c.ParamsList {
				expected, got := HourseJump(p.targetXPos, p.targetYPos, p.steps)
				if expected != got {
					t.Errorf("expected: %v, got: %v, case: (targetXPos=%v, targetYPos=%v, steps=%v)", expected, got, p.targetXPos, p.targetYPos, p.steps)
				} else {
					t.Logf("accepted: %v, case: (targetXPos=%v, targetYPos=%v, steps=%v)", expected, p.targetXPos, p.targetYPos, p.steps)
				}
			}
		})
	}
}
