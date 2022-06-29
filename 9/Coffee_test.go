package main

import (
	"math/rand"
	"testing"
	"time"
)

type Parameter struct {
	washMachineTime []int
	customers       int
	washTime        int
	dryTime         int
}

func GenRandomCases() []Parameter {
	rand.Seed(time.Now().UnixMicro())

	count := 0
	for count == 0 {
		count = rand.Intn(10)
	}

	cases := make([]Parameter, count)
	for i := 0; i < count; i++ {
		machines := 0
		for machines == 0 {
			machines = rand.Intn(5)
		}
		washMachineTime := make([]int, machines)
		for j := 0; j < machines; j++ {
			washMachineTime[j] = rand.Intn(8)
		}

		customers := 0
		for customers == 0 {
			customers = rand.Intn(20)
		}

		washTime := 0
		for washTime == 0 {
			washTime = rand.Intn(5)
		}

		dryTime := 0
		for dryTime == 0 {
			dryTime = rand.Intn(20)
		}

		cases[i] = Parameter{
			washMachineTime: washMachineTime,
			customers:       customers,
			washTime:        washTime,
			dryTime:         dryTime,
		}
	}

	return cases
}

func TestCoffee(t *testing.T) {
	cases := []struct {
		Name       string
		ParamsList []Parameter
	}{
		{
			"random", GenRandomCases(),
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			for _, p := range c.ParamsList {
				expected, got := Coffee(p.washMachineTime, p.customers, p.washTime, p.dryTime)
				if expected != got {
					t.Errorf("expected: %v, got: %v, case: (washMachineTime=%v, customers=%v, washTime=%v, dryTime=%v)", expected, got, p.washMachineTime, p.customers, p.washTime, p.dryTime)
				} else {
					t.Logf("accepted: %v, case: (washMachineTime=%v, customers=%v, washTime=%v, dryTime=%v)", expected, p.washMachineTime, p.customers, p.washTime, p.dryTime)
				}
			}
		})
	}
}
