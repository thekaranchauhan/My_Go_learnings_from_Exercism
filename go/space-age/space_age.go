package space

type Planet string

const secondsPerEarthYear float64 = 31557600.0

var planetOrbitPeriods = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	if period, found := planetOrbitPeriods[planet]; found {
		return seconds / (secondsPerEarthYear * period)
	}
	return -1.0
}
