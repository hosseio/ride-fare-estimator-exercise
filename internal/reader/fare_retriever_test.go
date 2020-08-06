package reader

import (
	"testing"
	"time"

	"github.com/hosseio/ride-fare-estimator-exercise/internal"
	"github.com/stretchr/testify/require"
)

func TestFareRetriever_Get(t *testing.T) {
	assertThat := require.New(t)
	t.Run("Given a fare retriever", func(t *testing.T) {
		viewMock := &internal.RideViewMock{}
		sut := NewFareRetriever(viewMock)
		t.Run("When it is executed having a ride in the system", func(t *testing.T) {
			viewMock.AllFunc = func() ([]internal.Ride, error) {
				ride, _ := internal.NewRide(42)
				now := time.Now()
				start := now.Add(-1 * time.Hour)

				p1, _ := internal.NewPosition(37.389098, -5.984379, start.Unix())
				err := ride.AddPosition(p1)
				assertThat.NoError(err)

				p3, _ := internal.NewPosition(37.389277, -5.984701, start.Add(5*time.Second).Unix())
				err = ride.AddPosition(p3)
				assertThat.NoError(err)

				return []internal.Ride{ride}, nil
			}

			fares, err := sut.Get()
			t.Run("Then the corresponding fares are returned", func(t *testing.T) {
				assertThat.NoError(err)
				assertThat.Len(fares, 1)
			})
		})
	})
}
