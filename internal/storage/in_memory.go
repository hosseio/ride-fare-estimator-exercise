package storage

import (
	"sync"

	"github.com/hosseio/ride-fare-estimator-exercise/internal"
)

type InMemory struct {
	sync.RWMutex
	rides map[int]internal.Ride
}

func NewInMemory() *InMemory {
	return &InMemory{
		rides: make(map[int]internal.Ride),
	}
}

func (r *InMemory) Get(id int) (internal.Ride, error) {
	r.Lock()
	defer r.Unlock()
	ride := r.rides[id]

	if ride.ID() == 0 { // zero value (invalid)
		return ride, internal.ErrRideNotFound
	}

	return ride, nil
}

func (r *InMemory) Save(ride internal.Ride) error {
	r.Lock()
	defer r.Unlock()

	r.rides[ride.ID()] = ride

	return nil
}

func (r *InMemory) All() ([]internal.Ride, error) {
	r.Lock()
	defer r.Unlock()

	var rides []internal.Ride
	for _, r := range r.rides {
		rides = append(rides, r)
	}

	return rides, nil
}
