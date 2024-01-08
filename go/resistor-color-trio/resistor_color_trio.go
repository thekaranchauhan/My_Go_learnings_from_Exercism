package resistorcolortrio

import "fmt"

var m = map[string]int{"black": 0, "brown": 1, "red": 2, "orange": 3, "yellow": 4, "green": 5, "blue": 6, "violet": 7, "grey": 8, "white": 9}
var units = []string{"", "kilo", "mega", "giga"}

// Decode and format the resistance.
func Label(colors []string) string {
	val, unit := m[colors[0]]*10+m[colors[1]], 0
	if val > 0 {

		// Multiply val by 10^colors[2]
		for pow := m[colors[2]]; pow > 0; pow, val = pow-1, val*10 {
		}

		// Select unit, and divide val by unit's val
		for ; val%1000 == 0; unit, val = unit+1, val/1000 {
		}
	}
	return fmt.Sprintf("%d %sohms", val, units[unit])
}
