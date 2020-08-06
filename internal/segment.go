package internal

import (
	"errors"
	"fmt"
	"time"

	"github.com/umahmood/haversine"
)

const (
	idleFareAmount            = 11.9    // per hour
	movingAtDaytimeFareAmount = 0.00074 // per m
	movingAtNightFareAmount   = 0.0013  // per m

	maxSpeed = 100

	msToKmh = 3.60
	kmToM   = 1000

	speedMoving = 10 // km/h

	limitDayHour = 5
)

type Segment struct {
	initialPosition Position
	endPosition     Position
	speed           float64 // speed in km per hour
	elapsedTime     time.Duration
	distanceCovered float64 // distance covered in m
	fare            float64
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

func (s Segment) Fare() float64 {
	return s.fare
}

var ErrTooMuchSpeed = errors.New("too much speed to create the segment")
var ErrInvalidPositionTimes = errors.New("the end position cannot be before the initial one. %v %v")

func NewSegmentFromPositions(initialPosition Position, endPosition Position) (Segment, error) {
	if initialPosition.Date().After(endPosition.Date()) {
		return Segment{}, fmt.Errorf("the end position cannot be before the initial one. %v %v", initialPosition.Date().Unix(), endPosition.Date().Unix())
	}
	distance := distanceInMeters(initialPosition, endPosition)
	duration := endPosition.Date().Sub(initialPosition.Date())

	speed := speedInKmH(distance, duration.Seconds())
	if speed > maxSpeed {
		return Segment{}, ErrTooMuchSpeed
	}

	segment := Segment{
		initialPosition: initialPosition,
		endPosition:     endPosition,
		speed:           speed,
		elapsedTime:     duration,
		distanceCovered: distance,
	}

	segment.calculateFare()

	return segment, nil
}

func speedInKmH(meters float64, seconds float64) float64 {
	metersPerSecond := meters / seconds

	return metersPerSecond * msToKmh
}

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

func (s *Segment) calculateFare() {
	if s.idle() {
		s.fare = idleFareAmount * s.elapsedTime.Hours()
		return
	}

	if s.atDay() {
		s.fare = movingAtDaytimeFareAmount * s.distanceCovered
		return
	}

	s.fare = movingAtNightFareAmount * s.distanceCovered
}

func (s Segment) idle() bool {
	return s.speed <= speedMoving
}

// assumption: the segment is at daytime or at night based only in the initial position timestamp
func (s Segment) atDay() bool {
	segmentAt := s.initialPosition.Date()
	y, m, d := segmentAt.Date()
	dayDate := time.Date(y, m, d, limitDayHour, 0, 0, 0, segmentAt.Location())

	return segmentAt.After(dayDate)
}

type SegmentList []Segment

func (l *SegmentList) Add(segment Segment) {
	*l = append(*l, segment)
}

func (l SegmentList) Len() int {
	return len([]Segment(l))
}
