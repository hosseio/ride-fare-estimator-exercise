package internal

import (
	"errors"
	"time"
)

const (
	minLat = -90
	maxLat = 90
	minLon = -180
	maxLon = 180
)

var ErrInvalidRideID = errors.New("invalid ride ID")

type Position struct {
	rideID      int
	coordinates Coordinates
	date        time.Time
}

func (p Position) Date() time.Time {
	return p.date
}

func (p Position) Coordinates() Coordinates {
	return p.coordinates
}

func (p Position) RideID() int {
	return p.rideID
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

func NewPosition(rideID int, lat float64, lon float64, timestamp int64) (Position, error) {
	if rideID <= 0 {
		return Position{}, ErrInvalidRideID
	}
	coordinates, err := NewCoordinates(lat, lon)
	if err != nil {
		return Position{}, err
	}
	date := time.Unix(timestamp, 0)

	return Position{rideID: rideID, coordinates: coordinates, date: date}, nil

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
