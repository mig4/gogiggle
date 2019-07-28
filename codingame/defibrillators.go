package codingame

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const EARTH_RADIUS = 6371
const RADIAN_MULTIPLIER = math.Pi / 180

/*
AngleDec represents an angle in Decimal Degrees

https://en.wikipedia.org/wiki/Decimal_degrees

Angle North (+0°-90°) or South (-0°-90°) of Equator (0°) = Latitude,
or East (+0°-180°) or West (-0°-180°) of Prime Meridian (0°) = Longitude.
*/
type AngleDec float64

/*
FromString constructs a new AngleDec from string.

It cleans up the string (separators are normalised to `.`) and parses it as
a float and converts to AngleDec.
*/
func FromString(s string) (AngleDec, error) {
	f, err := strconv.ParseFloat(
		strings.Replace(strings.TrimSpace(s), ",", ".", 1), 64,
	)
	if err != nil {
		return 0, err
	}
	if math.IsNaN(f) {
		return 0, fmt.Errorf("%v isn't a valid value for degree of an angle", f)
	}
	return AngleDec(f), nil
}

/*
Radians converts this angle from degrees to radians.

https://en.wikipedia.org/wiki/Radian

Formula:

    AngleRad = AngleDec * (π/180)
*/
func (this AngleDec) Radians() float64 {
	return float64(this * RADIAN_MULTIPLIER)
}

/*
Position encapsulates geographic coordinates of a point.
*/
type Position struct {
	latitude  AngleDec
	longitude AngleDec
}

/*
NewPosition validates given coordinates and creates a new Position.

The angles are given as strings representing decimal degrees, they are cleaned
up (separators normalised to `.`), converted into AngleDec (float) and
validated that they are in range.

Return an error if that fails, otherwise return the new Position.
*/
func NewPosition(latitude string, longitude string) (*Position, error) {
	latDec, err := FromString(latitude)
	if err != nil {
		return nil, fmt.Errorf("Could not parse %v as latitude: %s", latitude, err)
	}
	if -90 > latDec || latDec > 90 {
		return nil, fmt.Errorf("Latitude out of bounds (-90°..90°): %v", latDec)
	}
	lonDec, err := FromString(longitude)
	if err != nil {
		return nil, fmt.Errorf("Could not parse %v as longitude: %s", longitude, err)
	}
	if -180 > lonDec || lonDec > 180 {
		return nil, fmt.Errorf("Longitude out of bounds (-180°..180°): %v", lonDec)
	}
	return &Position{latDec, lonDec}, nil
}

/*
DistanceTo calculates distance from this to other position.

Formula:

	x = (otherLongitude - thisLongitude) * cos((thisLatitude + otherLatitude) / 2)
	y = (otherLatitude - thisLatitude)
	distance = sqrt(x² + y²) * EARTH_RADIUS
*/
func (this *Position) DistanceTo(other Position) float64 {
	x := float64(other.longitude-this.longitude) *
		math.Cos(float64(this.latitude+other.latitude)/2)
	y := float64(other.latitude - this.latitude)
	return math.Sqrt(math.Pow(x, 2)+math.Pow(y, 2)) * EARTH_RADIUS
}

/*
Defibrillator encapsulates details and position of a defibrillator.
*/
type Defibrillator struct {
	id       int
	name     string
	address  string
	phone    string
	location Position
}

/*
DefibrillatorFromFields parses the given fields and returns a Defibrillator.

The fields are a result of reading in the `;` separated input lines and they
correspond to fields of the Defibrillator structure (with in-line position
fields) in order. Note that position fields are in reverse order than in
Position struct, i.e. longitude, then latitude.

The last two fields are converted into Position using NewPosition and that,
along with other fields is used as values for the Defibrillator.
*/
func DefibrillatorFromFields(fields []string) (*Defibrillator, error) {
	if len(fields) != 6 {
		return nil, fmt.Errorf("Expected 6 fields, got %d: %v", len(fields), fields)
	}
	this := Defibrillator{}

	id, err := strconv.Atoi(fields[0])
	if err != nil {
		return &this, err
	}
	this.id = id

	this.name = strings.TrimSpace(fields[1])
	this.address = strings.TrimSpace(fields[2])
	this.phone = strings.TrimSpace(fields[3])

	if this.name == "" {
		return &this, fmt.Errorf("name cannot be empty")
	}

	location, err := NewPosition(fields[5], fields[4])
	if err != nil {
		return &this, err
	}
	this.location = *location
	return &this, nil
}

/*
MapToDefibrillators maps each line in `lines` to a new Defibrillator.

It does this by calling DefibrillatorFromFields for each line. The given lines
should be slices of fields as read from CSV (`;` separated) input.
*/
func MapToDefibrillators(lines [][]string) (defibrillators []Defibrillator, err error) {
	defibrillators = make([]Defibrillator, len(lines))
	for i, line := range lines {
		defib, err := DefibrillatorFromFields(line)
		if err != nil {
			break
		}
		defibrillators[i] = *defib
	}
	return
}

/*
FindDefibrillator finds and returns a pointer to a Defibrillator in the given
slice that matches the given predicate (i.e. for which the given predicate
function returns true).
*/
func FindDefibrillator(
	defibrillators []Defibrillator,
	predicate func(*Defibrillator) bool,
) (*Defibrillator, error) {
	for _, defib := range defibrillators {
		if predicate(&defib) {
			return &defib, nil
		}
	}
	return nil, fmt.Errorf(
		"predicate didn't match any of the %d defibrillators",
		len(defibrillators),
	)
}

/*
FindNearestDefibrillator finds a Defibrillator from the given list that is
nearest (according to Position.DistanceTo) to the `user`.
*/
func FindNearestDefibrillator(user *Position, defibrillators []Defibrillator) (*Defibrillator, error) {
	var nearest Defibrillator
	distance := math.Inf(1)
	for _, defibrillator := range defibrillators {
		// fmt.Fprintf(os.Stderr, "user=%v, defibrillator=%v\n", user, defibrillator)
		if thisDist := user.DistanceTo(defibrillator.location); thisDist < distance {
			nearest = defibrillator
			distance = thisDist
		}
	}
	return &nearest, nil
}
