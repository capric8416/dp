package main

import "fmt"

func StikersToSpellWord(target string, stickers []string) (expected int, got int) {
	expected = StikersToSpellWordRecurse(target, stickers)
	got = StikersToSpellWordDp(target, stickers)
	return
}

func main() {
	fmt.Println(StikersToSpellWord("babac", []string{"ba", "c", "abcd"}))
}
