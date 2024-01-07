package twelve

import (
	"fmt"
	"strings"
)

func Verse(i int) string {

	numerals := []string{"", "first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
	gifts := []string{"twelve Drummers Drumming", "eleven Pipers Piping", "ten Lords-a-Leaping", "nine Ladies Dancing", "eight Maids-a-Milking", "seven Swans-a-Swimming", "six Geese-a-Laying", "five Gold Rings", "four Calling Birds", "three French Hens", "two Turtle Doves", "a Partridge in a Pear Tree"}
	if i > 1 {
		gifts[len(gifts)-1] = "and " + gifts[len(gifts)-1]
	}
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", numerals[i],
		strings.Join(gifts[len(gifts)-i:], ", "))
}

func Song() string {
	var verses = make([]string, 12)
	for i := 0; i < 12; i++ {
		verses[i] = Verse(i + 1)
	}
	return strings.Join(verses, "\n")
}
