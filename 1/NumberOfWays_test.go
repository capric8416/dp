package main

import (
	"testing"
)

func TestNumberOfWays(t *testing.T) {
	type Parameters struct {
		maxPos     int
		currentPos int
		moveSteps  int
		targetPos  int
	}

	cases := []struct {
		Name       string
		ParamsList []Parameters
	}{
		{
			"ok", []Parameters{
				{maxPos: 7, currentPos: 2, moveSteps: 4, targetPos: 4},
				{maxPos: 4, currentPos: 2, moveSteps: 4, targetPos: 4},
				{maxPos: 5, currentPos: 2, moveSteps: 4, targetPos: 6},
				{maxPos: 7, currentPos: 4, moveSteps: 9, targetPos: 5},
				{maxPos: 20, currentPos: 9, moveSteps: 10, targetPos: 13},
				{maxPos: 20, currentPos: 9, moveSteps: 10, targetPos: 15},
				{maxPos: 20, currentPos: 9, moveSteps: 10, targetPos: 16},
			},
		},
		{
			"moveSteps < 0", []Parameters{
				{maxPos: 7, currentPos: 2, moveSteps: -1, targetPos: 4},
				{maxPos: 7, currentPos: 2, moveSteps: -4, targetPos: 4},
			},
		},
		{
			"maxPos < 1", []Parameters{
				{maxPos: 0, currentPos: 2, moveSteps: 4, targetPos: 4},
				{maxPos: -1, currentPos: 2, moveSteps: 4, targetPos: 4},
			},
		},
		{
			"currentPos < 1", []Parameters{
				{maxPos: 7, currentPos: 0, moveSteps: 4, targetPos: 4},
				{maxPos: 7, currentPos: -1, moveSteps: 4, targetPos: 4},
			},
		},
		{
			"currentPos > maxPos", []Parameters{
				{maxPos: 7, currentPos: 8, moveSteps: 4, targetPos: 4},
				{maxPos: 7, currentPos: 9, moveSteps: 4, targetPos: 4},
			},
		},
		{
			"targetPos < 1", []Parameters{
				{maxPos: 7, currentPos: 2, moveSteps: 4, targetPos: 0},
				{maxPos: 7, currentPos: 2, moveSteps: 4, targetPos: -1},
			},
		},
		{
			"targetPos > maxPos", []Parameters{
				{maxPos: 7, currentPos: 2, moveSteps: 4, targetPos: 8},
				{maxPos: 7, currentPos: 2, moveSteps: 4, targetPos: 9},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			for _, p := range c.ParamsList {
				expected, got := NumberOfWays(p.maxPos, p.currentPos, p.moveSteps, p.targetPos)
				if expected != got {
					t.Errorf("expected: %v, got: %v, case: (maxPos=%v, currentPos=%v, moveSteps=%v, targetPos=%v)",
						expected, got, p.maxPos, p.currentPos, p.moveSteps, p.targetPos)
				} else {
					t.Logf("accepted: %v", got)
				}
			}
		})
	}
}
