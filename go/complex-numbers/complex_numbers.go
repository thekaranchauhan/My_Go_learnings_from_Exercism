package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
	re  float64
	img float64
}

func (n Number) Real() float64 {
	return n.re
}

func (n Number) Imaginary() float64 {
	return n.img
}

func (n1 Number) Add(n2 Number) Number {
	return Number{n1.re + n2.re, n1.img + n2.img}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{n1.re - n2.re, n1.img - n2.img}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{n1.re*n2.re - n1.img*n2.img, n1.img*n2.re + n1.re*n2.img}
}

func (n Number) Times(factor float64) Number {
	return Number{factor * n.re, factor * n.img} // WTF?
	/*
	    r := math.Pow(n.Abs(), factor)
	    arg := math.Atan(n.img / n.re) * factor
		return Number{ r * math.Cos(arg), r * math.Sin(arg)}
	*/
}

func (n1 Number) Divide(n2 Number) Number {
	den := n2.re*n2.re + n2.img*n2.img
	return Number{(n1.re*n2.re + n1.img*n2.img) / den, (n1.img*n2.re - n1.re*n2.img) / den}
}

func (n Number) Conjugate() Number {
	return Number{n.re, -n.img}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.re*n.re + n.img*n.img)
}

func (n Number) Exp() Number {
	return Number{math.Exp(n.re) * math.Cos(n.img), math.Exp(n.re) * math.Sin(n.img)}
}
