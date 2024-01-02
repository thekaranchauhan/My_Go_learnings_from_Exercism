package triangle

import "math"

// Kind is the triangle kind
type Kind int

// Kind constants
const (
	NaT Kind = iota // Not a triangle
	Equ             // Equilateral
	Iso             // Isosceles
	Sca             // Scalene
)

// KindFromSides tells if a triangle is equilateral, isosceles or scalene
func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a+b+c) || a+b <= c || a+c <= b || b+c <= a {
		return NaT
	}
	if a == b && a == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}
