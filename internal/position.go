package internal

import (
	"time"
)

const (
	minLat = -90
	maxLat = 90
	minLon = -180
	maxLon = 180
)

type Position struct {
	coordinates Coordinates
	date        time.Time
}

func (p Position) Date() time.Time {
	return p.date
}

func (p Position) Coordinates() Coordinates {
	return p.coordinates
}

type Coordinates struct {
	lat float64
	lon float64
}

func (c Coordinates) Lon() float64 {
	return c.lon
}

func (c Coordinates) Lat() float64 {
	return c.lat
}

func NewPosition(lat float64, lon float64, timestamp int64) (Position, error) {
	coordinates, err := NewCoordinates(lat, lon)
	if err != nil {
		return Position{}, err
	}
	date := time.Unix(timestamp, 0)

	return Position{coordinates: coordinates, date: date}, nil
}

func NewCoordinates(lat float64, lon float64) (Coordinates, error) {
	if lat < minLat || lat > maxLat {
		return Coordinates{}, ErrInvalidCoordinates{"invalid latitude"}
	}

	if lon < minLon || lon > maxLon {
		return Coordinates{}, ErrInvalidCoordinates{"invalid longitude"}
	}

	return Coordinates{lat: lat, lon: lon}, nil
}

type ErrInvalidCoordinates struct {
	message string
}

func (e ErrInvalidCoordinates) Error() string {
	return e.message
}
