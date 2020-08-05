package internal

import (
	"errors"
	"github.com/umahmood/haversine"
	"time"
)

type Segment struct {
	initialPosition Position
	endPosition     Position
	speed           float64 // speed in km per hour
	elapsedTime     time.Duration
	distanceCovered float64 // distance covered in m
}

func (s Segment) InitialPosition() Position {
	return s.initialPosition
}

func (s Segment) EndPosition() Position {
	return s.endPosition
}

func (s Segment) Speed() float64 {
	return s.speed
}

func (s Segment) ElapsedTime() time.Duration {
	return s.elapsedTime
}

func (s Segment) DistanceCovered() float64 {
	return s.distanceCovered
}

const maxSpeed = 100

var ErrTooMuchSpeed = errors.New("to much speed to create the segment")

func NewSegmentFromPositions(initialPosition Position, endPosition Position) (Segment, error) {
	distance := distanceInMeters(initialPosition, endPosition)
	duration := initialPosition.Date().Sub(endPosition.Date())
	if endPosition.Date().After(initialPosition.Date()) {
		duration = endPosition.Date().Sub(initialPosition.Date())
	}

	speed := speedInKmH(distance, duration.Seconds())
	if speed > maxSpeed {
		return Segment{}, ErrTooMuchSpeed
	}

	return Segment{
		initialPosition: initialPosition,
		endPosition:     endPosition,
		speed:           speed,
		elapsedTime:     duration,
		distanceCovered: distance,
	}, nil
}

const msToKmh = 3.60

func speedInKmH(meters float64, seconds float64) float64 {
	metersPerSecond := meters / seconds

	return metersPerSecond * msToKmh
}

const kmToM = 1000

func distanceInMeters(position Position, position2 Position) float64 {
	_, km := haversine.Distance(
		haversine.Coord{
			Lat: position.Coordinates().Lat(),
			Lon: position.Coordinates().Lon(),
		},
		haversine.Coord{
			Lat: position2.Coordinates().Lat(),
			Lon: position2.Coordinates().Lon(),
		},
	)

	return km * kmToM
}

type SegmentList []Segment

func (l *SegmentList) Add(segment Segment) {
	*l = append(*l, segment)
}

func (l SegmentList) Len() int {
	return len([]Segment(l))
}
