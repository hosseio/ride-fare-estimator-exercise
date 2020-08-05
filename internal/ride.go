package internal

import "errors"

type Ride struct {
	id              int
	segments        SegmentList
	currentPosition Position
}

var ErrInvalidRideID = errors.New("invalid ride ID")

func NewRide(id int) (Ride, error) {
	if id <= 1 {
		return Ride{}, ErrInvalidRideID
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
