package allergies

var allergyMap = map[string]uint{"eggs": 1 << 0, "peanuts": 1 << 1, "shellfish": 1 << 2,
	"strawberries": 1 << 3, "tomatoes": 1 << 4, "chocolate": 1 << 5, "pollen": 1 << 6, "cats": 1 << 7}

func Allergies(score uint) []string {
	allergies := make([]string, 0, len(allergyMap))
	for desc, flag := range allergyMap {
		if score&flag != 0 {
			allergies = append(allergies, desc)
		}
	}
	return allergies
}

func AllergicTo(score uint, allergen string) bool { return score&allergyMap[allergen] != 0 }
