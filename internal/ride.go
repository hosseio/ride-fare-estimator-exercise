package internal

import "errors"

type Ride struct {
	id              int
	segments        SegmentList
	currentPosition Position
}

var ErrInvalidRideID = errors.New("invalid ride ID")

func NewRide(id int) (Ride, error) {
	if id <= 0 {
		return Ride{}, ErrInvalidRideID
	}

	return Ride{id: id}, nil
}

func (r *Ride) AddPosition(position Position) {
	if r.currentPosition == (Position{}) {
		r.currentPosition = position
		return
	}

	segment, err := NewSegmentFromPositions(r.currentPosition, position)
	if err != nil {
		return
	}

	r.segments.Add(segment)
	r.currentPosition = position
}

func (r Ride) ID() int {
	return r.id
}

type Memento struct {
	ID              int
	Segments        SegmentList
	CurrentPosition Position
}

func RestoreState(state Memento) Ride {
	return Ride{state.ID, state.Segments, state.CurrentPosition}
}
