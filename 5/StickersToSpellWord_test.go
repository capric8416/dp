package main

import (
	"math/rand"
	"testing"
	"time"
)

type Parameter struct {
	target   string
	stickers []string
}

func GenRandomCases() []Parameter {
	rand.Seed(time.Now().UnixMicro())

	count := 0
	for count == 0 {
		count = rand.Intn(10)
	}

	cases := make([]Parameter, count)
	for i := 0; i < count; i++ {
		targetLength := 0
		for targetLength == 0 {
			targetLength = rand.Intn(10)
		}
		t := make([]rune, targetLength)
		for j := 0; j < targetLength; j++ {
			t[j] = 'a' + rune(rand.Intn(26))
		}
		target := string(t)

		stickersCount := 0
		for stickersCount == 0 {
			stickersCount = rand.Intn(10)
		}
		stickers := make([]string, stickersCount)
		for k := 0; k < stickersCount; k++ {
			stickerLength := 0
			for stickerLength == 0 {
				stickerLength = rand.Intn(10)
			}

			sticker := make([]rune, stickerLength)
			for m := 0; m < stickerLength; m++ {
				sticker[m] = 'a' + rune(rand.Intn(26))
			}
			stickers[k] = string(sticker)
		}

		cases[i] = Parameter{target: target, stickers: stickers}
	}

	return cases
}

func TestStickersToSpellWord(t *testing.T) {
	cases := []struct {
		Name       string
		ParamsList []Parameter
	}{
		{
			"ok", []Parameter{
				{target: "babac", stickers: []string{"ba", "c", "abcd"}},
			},
		},
		{
			"empty", []Parameter{
				{target: "", stickers: []string{}},
			},
		},
		{
			"random", GenRandomCases(),
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			for _, p := range c.ParamsList {
				expected, got := StikersToSpellWord(p.target, p.stickers)
				if expected != got {
					t.Errorf("expected: %v, got: %v, case: (target=%v, stickers=%v)", expected, got, p.target, p.stickers)
				} else {
					t.Logf("accepted: %v, case: (target=%v, stickers=%v)", expected, p.target, p.stickers)
				}
			}
		})
	}
}
