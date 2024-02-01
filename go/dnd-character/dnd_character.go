package dndcharacter

import (
	"math"
	"math/rand"
	"time"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	// Seed the random number generator to ensure different results each time
	rand.New(rand.NewSource(time.Now().UnixNano()))
	// Generate a random number between 3 and 18 (inclusive)
	return rand.Intn(16) + 3
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	avatar := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability()}

	avatar.Hitpoints = 10 + Modifier(avatar.Constitution)
	return avatar
}
