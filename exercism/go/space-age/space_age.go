package space

type Planet string

const EarthOrbit float64 = 31557600

var PlanetOrbits = map[Planet]float64{
	"Earth":   1 * EarthOrbit,
	"Mercury": 0.2408467 * EarthOrbit,
	"Venus":   0.61519726 * EarthOrbit,
	"Mars":    1.8808158 * EarthOrbit,
	"Jupiter": 11.862615 * EarthOrbit,
	"Saturn":  29.447498 * EarthOrbit,
	"Uranus":  84.016846 * EarthOrbit,
	"Neptune": 164.79132 * EarthOrbit,
}

func Age(seconds float64, planet Planet) float64 {
	return seconds / PlanetOrbits[planet]
}
