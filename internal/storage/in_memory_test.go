//+build integration

package storage

import (
	"testing"
	"time"

	"github.com/hosseio/ride-fare-estimator-exercise/internal"
	"github.com/stretchr/testify/require"
)

func TestInMemoryRideRepository(t *testing.T) {
	assertThat := require.New(t)
	t.Run("Given an in memory ride repository", func(t *testing.T) {
		sut := NewInMemory()
		t.Run("When a ride with 2 segments is saved", func(t *testing.T) {
			ride, _ := internal.NewRide(42)
			now := time.Now()
			start := now.Add(-1 * time.Hour)
			p1, _ := internal.NewPosition(37.389098, -5.984379, start.Unix())
			_ = ride.AddPosition(p1)
			p3, _ := internal.NewPosition(37.389277, -5.984701, start.Add(5*time.Second).Unix())
			_ = ride.AddPosition(p3)
			p4, _ := internal.NewPosition(37.389391, -5.985022, start.Add(10*time.Second).Unix())
			_ = ride.AddPosition(p4)

			err := sut.Save(ride)
			t.Run("Then it can be retrieved by id", func(t *testing.T) {
				assertThat.NoError(err)

				retrieved, err := sut.Get(42)
				assertThat.NoError(err)
				assertThat.Equal(42, retrieved.ID())
				assertThat.Equal(2, retrieved.Segments().Len())
				assertThat.Equal(37.389391, retrieved.CurrentPosition().Coordinates().Lat())
			})
		})
		t.Run("When a non-existence ride is asked", func(t *testing.T) {
			_, err := sut.Get(1138)
			t.Run("Then a not found error is returned", func(t *testing.T) {
				assertThat.Error(err)
				assertThat.Equal(internal.ErrRideNotFound, err)
			})
		})
	})
}
