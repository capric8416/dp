package main

import (
	"math/rand"
	"testing"
	"time"
)

func GenRandomCases() []string {
	rand.Seed(time.Now().UnixMicro())

	count := 0
	for count == 0 {
		count = rand.Intn(50)
	}

	cases := make([]string, count)
	for i := 0; i < count; i++ {
		length := 0
		for length == 0 {
			length = rand.Intn(20)
		}

		digits := make([]rune, length)
		for j := 0; j < length; j++ {
			digits[j] = '0' + rand.Int31n(10)
		}

		cases[i] = string(digits)
	}

	return cases
}

func TestConvertToLetterString(t *testing.T) {
	cases := []struct {
		Name       string
		ParamsList []string
	}{
		{
			"empty", []string{
				"",
			},
		},
		{
			"one", []string{
				"0",
				"1",
			},
		},
		{
			"two", []string{
				"10",
				"20",
				"30",
			},
		},
		{
			"three", []string{
				"101",
				"111",
				"305",
			},
		},
		{
			"random",
			GenRandomCases(),
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			for _, digits := range c.ParamsList {
				expected, got := ConvertToLetterString(digits)
				if expected != got {
					t.Errorf("expected: %v, got: %v, case: %#v", expected, got, digits)
				} else {
					t.Logf("accepted: %v, case: %#v", expected, digits)
				}
			}
		})
	}
}
