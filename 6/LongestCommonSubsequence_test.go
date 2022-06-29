package main

import (
	"math/rand"
	"testing"
	"time"
)

type Parameter struct {
	text1 string
	text2 string
}

func GenRandomCases() []Parameter {
	rand.Seed(time.Now().UnixMicro())

	count := 0
	for count == 0 {
		count = rand.Intn(50)
	}
	cases := make([]Parameter, count)

	for i := 0; i < count; i++ {
		text1Length := 0
		for text1Length == 0 {
			text1Length = rand.Intn(20)
		}
		rune1 := make([]rune, text1Length)
		for j := 0; j < text1Length; j++ {
			rune1[j] = 'a' + rune(rand.Intn(26))
		}
		cases[i].text1 = string(rune1)

		text2Length := 0
		for text2Length == 0 {
			text2Length = rand.Intn(20)
		}
		rune2 := make([]rune, text2Length)
		for j := 0; j < text2Length; j++ {
			rune2[j] = 'a' + rune(rand.Intn(26))
		}
		cases[i].text2 = string(rune2)
	}

	return cases
}

func TestLongestCommonSubsequence(t *testing.T) {
	cases := []struct {
		Name       string
		ParamsList []Parameter
	}{
		{
			"ok", []Parameter{
				{text1: "abcde", text2: "ace"},
				{text1: "abc", text2: "abc"},
				{text1: "abc", text2: "def"},
				{text1: "abcfffffffffdm", text2: "agcyfxdom"},
			},
		},
		{
			"empty", []Parameter{
				{text1: "", text2: "def"},
				{text1: "abc", text2: ""},
				{text1: "", text2: ""},
			},
		},
		{
			"random", GenRandomCases(),
		},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			for _, p := range c.ParamsList {
				expected, got := LongestCommonSubsequence(p.text1, p.text2)
				if expected != got {
					t.Errorf("expected: %#v, got: %#v, case: (text1=%#v, text2=%#v)", expected, got, p.text1, p.text2)
				} else {
					t.Logf("accepted: %#v, case: (text1=%#v, text2=%#v)", expected, p.text1, p.text2)
				}
			}
		})
	}
}
