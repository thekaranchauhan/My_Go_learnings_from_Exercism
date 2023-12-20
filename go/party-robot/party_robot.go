package partyrobot

import (
	"fmt"
	"strings"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var builder strings.Builder

	builder.WriteString(Welcome(name) + "\n")
	builder.WriteString("You have been assigned to table " + fmt.Sprintf("%03d", table) + ". ")
	builder.WriteString("Your table is " + direction + ", exactly " + fmt.Sprintf("%.1f", distance) + " meters from here.\n")
	builder.WriteString("You will be sitting next to " + neighbor + ".")

	defer builder.Reset()

	return builder.String()
}
