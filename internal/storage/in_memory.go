package storage

import (
	"sync"

	"github.com/hosseio/ride-fare-estimator-exercise/internal"
)

type InMemoryRideRepository struct {
	sync.RWMutex
	rides map[int]internal.Ride
}

func NewInMemoryRideRepository() *InMemoryRideRepository {
	return &InMemoryRideRepository{
		rides: make(map[int]internal.Ride),
	}
}

func (r *InMemoryRideRepository) Get(id int) (internal.Ride, error) {
	r.RLock()
	defer r.RUnlock()
	ride := r.rides[id]

	if ride.ID() == 0 { // zero value (invalid)
		return ride, internal.ErrRideNotFound
	}

	return ride, nil
}

func (r *InMemoryRideRepository) Save(ride internal.Ride) error {
	r.RLock()
	defer r.RUnlock()

	r.rides[ride.ID()] = ride

	return nil
}
