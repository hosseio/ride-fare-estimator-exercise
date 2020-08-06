package internal

//go:generate moq -out ride_view_mock.go . RideView
type RideView interface {
	All() ([]Ride, error)
}
