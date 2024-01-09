package foodchain

import (
	"fmt"
	"strings"
)

type item struct {
	name, action string
}

func Verse(v int) string {
	items := []item{
		item{name: "fly", action: ""},
		item{name: "spider", action: "It wriggled and jiggled and tickled inside her."},
		item{name: "bird", action: "How absurd to swallow a bird!"},
		item{name: "cat", action: "Imagine that, to swallow a cat!"},
		item{name: "dog", action: "What a hog, to swallow a dog!"},
		item{name: "goat", action: "Just opened her throat and swallowed a goat!"},
		item{name: "cow", action: "I don't know how she swallowed a cow!"},
		item{name: "horse", action: "She's dead, of course!"},
	}

	stanzas := []string{
		fmt.Sprintf("I know an old lady who swallowed a %s.", items[v-1].name),
	}

	if v > 1 {
		stanzas = append(stanzas, items[v-1].action)
	}

	if v == 8 {
		return strings.Join(stanzas, "\n")
	}

	for s := v - 1; s > 0; s-- {
		chain := fmt.Sprintf("She swallowed the %s to catch the %s", items[s].name, items[s-1].name)

		if items[s-1].name == "spider" {
			chain = chain + " that wriggled and jiggled and tickled inside her."
		} else {
			chain = chain + "."
		}

		stanzas = append(stanzas, chain)
	}

	stanzas = append(stanzas, "I don't know why she swallowed the fly. Perhaps she'll die.")
	return strings.Join(stanzas, "\n")
}

func Verses(start, end int) string {
	var verse []string
	for v := start; v <= end; v++ {
		verse = append(verse, Verse(v))
	}

	return strings.Join(verse, "\n\n")
}

func Song() string {
	return Verses(1, 8)
}
