package internal

import "errors"

var ErrRideNotFound = errors.New("ride not found")

//go:generate moq -out ride_repository_mock.go . RideRepository
type RideRepository interface {
	Save(Ride) error
	Get(id int) (Ride, error)
}
