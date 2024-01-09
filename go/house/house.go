package house

import (
	"strings"
)

var subject = []string{
	"", "house that Jack built.", "malt", "rat", "cat", "dog", "cow with the crumpled horn",
	"maiden all forlorn", "man all tattered and torn",
	"priest all shaven and shorn", "rooster that crowed in the morn",
	"farmer sowing his corn", "horse and the hound and the horn",
}
var action = []string{
	"", "", "lay in", "ate", "killed", "worried", "tossed",
	"milked", "kissed", "married", "woke", "kept",
	"belonged to",
}

func Verse(v int) string {
	verse := "This is the " + subject[v]
	for i := v; i > 1; i-- {
		verse += "\nthat " + action[i] + " the " + subject[i-1]
	}
	return verse
}

func Song() string {
	verses := make([]string, 12)
	for i := 1; i <= 12; i++ {
		verses[i-1] = Verse(i)
	}
	return strings.Join(verses, "\n\n")
}
