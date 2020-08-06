package internal

import (
	"errors"
	"fmt"
)

type Ride struct {
	id              int
	segments        SegmentList
	currentPosition Position
}

var ErrInvalidRideID = errors.New("invalid ride ID")

func NewRide(id int) (Ride, error) {
	if id <= 0 {
		return Ride{}, fmt.Errorf("invalid ride ID %v", id)
	}

	return Ride{id: id}, nil
}

func (r *Ride) AddPosition(position Position) error {
	if r.currentPosition == (Position{}) {
		r.currentPosition = position
		return nil
	}

	segment, err := NewSegmentFromPositions(r.currentPosition, position)
	if err != nil {
		return err
	}

	r.segments.Add(segment)
	r.currentPosition = position

	return nil
}

func (r Ride) ID() int {
	return r.id
}

func (r Ride) CurrentPosition() Position {
	return r.currentPosition
}

func (r Ride) Segments() SegmentList {
	return r.segments
}

func (r Ride) FareEstimation() float64 {
	var total float64
	for _, segment := range r.Segments() {
		total = total + segment.Fare()
	}

	return total
}
