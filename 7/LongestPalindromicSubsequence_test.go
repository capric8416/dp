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
		strLength := 0
		for strLength == 0 {
			strLength = rand.Intn(20)
		}
		runeList := make([]rune, strLength)
		for j := 0; j < strLength; j++ {
			runeList[j] = 'a' + rune(rand.Intn(26))
		}
		cases[i] = string(runeList)
	}

	return cases
}

func TestLongestPalindromicSubsequence(t *testing.T) {
	cases := []struct {
		Name       string
		ParamsList []string
	}{
		{
			"ok", []string{
				"bbbab",
				"a12b3c43def2ghi1kpm",
			},
		},
		{
			"empty", []string{
				"",
			},
		},
		{
			"random", GenRandomCases(),
		},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			for _, s := range c.ParamsList {
				expected, got := LongestPalindromicSubsequence(s)
				if expected != got {
					t.Errorf("expected: %v, got: %v, case: (s=%#v)", expected, got, s)
				} else {
					t.Logf("accepted: %v, case: (s=%#v)", expected, s)
				}
			}
		})
	}
}
