package location

type LocationExtractor interface {
	GetLocation() (Location, Location, error)
}

type Location struct {
	Degrees   int
	Minutes   int
	Seconds   float64
	Direction string
}
