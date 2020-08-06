package creator

import (
	"testing"
	"time"

	"github.com/hosseio/ride-fare-estimator-exercise/internal"
	"github.com/stretchr/testify/require"
)

func TestCreatePositionCommandHandler(t *testing.T) {
	assertThat := require.New(t)

	t.Run("Given a create position command handler", func(t *testing.T) {
		repository := &internal.RideRepositoryMock{}
		sut := NewCreatePositionCommandHandler(repository)
		var saveCalls int

		t.Run("When adding a position of a non existence ride", func(t *testing.T) {
			command := CreatePositionCommand{
				RideID:    42,
				Lat:       1,
				Lon:       1,
				Timestamp: time.Now().Unix(),
			}

			repository.GetFunc = func(id int) (internal.Ride, error) {
				return internal.Ride{}, internal.ErrRideNotFound
			}
			repository.SaveFunc = func(ride internal.Ride) error {
				saveCalls++
				assertThat.Equal(42, ride.ID())
				return nil
			}

			err := sut.Handle(command)
			t.Run("Then the new ride is saved", func(t *testing.T) {
				assertThat.NoError(err)
				assertThat.Equal(1, saveCalls)
				saveCalls = 0
			})
		})

		t.Run("When adding a position creating a segment of an existence ride", func(t *testing.T) {
			ride, _ := internal.NewRide(42)
			p1, _ := internal.NewPosition(37.389098, -5.984379, time.Now().Add(5*time.Second).Unix())
			_ = ride.AddPosition(p1)

			command := CreatePositionCommand{
				RideID: 42,
				// the following is a few meters in 5 seconds
				Lat:       37.389277,
				Lon:       -5.984701,
				Timestamp: time.Now().Unix(),
			}

			repository.GetFunc = func(id int) (internal.Ride, error) {
				return ride, nil
			}
			repository.SaveFunc = func(ride internal.Ride) error {
				saveCalls++
				assertThat.Equal(42, ride.ID())
				// a new segment should exist
				assertThat.Equal(1, ride.Segments().Len())
				return nil
			}

			err := sut.Handle(command)
			t.Run("Then the new ride is saved", func(t *testing.T) {
				assertThat.NoError(err)
				assertThat.Equal(1, saveCalls)
				saveCalls = 0
			})
		})
		t.Run("When adding a position to be ignored", func(t *testing.T) {
			ride, _ := internal.NewRide(42)
			p1, _ := internal.NewPosition(37.389098, -5.984379, time.Now().Add(5*time.Second).Unix())
			_ = ride.AddPosition(p1)

			command := CreatePositionCommand{
				RideID: 42,
				// the following is a few kms in 5 seconds: too much speed => ignore new position
				Lat:       37.888339,
				Lon:       -4.779336,
				Timestamp: time.Now().Unix(),
			}

			repository.GetFunc = func(id int) (internal.Ride, error) {
				return ride, nil
			}
			repository.SaveFunc = func(ride internal.Ride) error {
				return nil
			}

			err := sut.Handle(command)
			t.Run("Then the new ride is NOT saved", func(t *testing.T) {
				assertThat.NoError(err)
				assertThat.Equal(0, saveCalls)
			})
		})
	})
}
