package internal

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewPosition(t *testing.T) {
	assertThat := require.New(t)

	t.Run("Given a negative ride ID", func(t *testing.T) {
		rideID := -1
		lat := 0.0
		lon := 0.0
		timestamp := time.Now().Unix()
		t.Run("When the position is created", func(t *testing.T) {
			_, err := NewPosition(rideID, lat, lon, timestamp)
			t.Run("Then an error is returned", func(t *testing.T) {
				assertThat.Error(err)
			})
		})
	})
	t.Run("Given an invalid latitude", func(t *testing.T) {
		rideID := 1
		lat := -100.0
		lon := 0.0
		timestamp := time.Now().Unix()
		t.Run("When the position is created", func(t *testing.T) {
			_, err := NewPosition(rideID, lat, lon, timestamp)
			t.Run("Then an error is returned", func(t *testing.T) {
				assertThat.Error(err)
			})
		})
	})
	t.Run("Given an invalid longitude", func(t *testing.T) {
		rideID := 1
		lat := 0.0
		lon := 190.0
		timestamp := time.Now().Unix()
		t.Run("When the position is created", func(t *testing.T) {
			_, err := NewPosition(rideID, lat, lon, timestamp)
			t.Run("Then an error is returned", func(t *testing.T) {
				assertThat.Error(err)
			})
		})
	})
	t.Run("Given valid data", func(t *testing.T) {
		rideID := 1
		lat := 0.0
		lon := 0.0
		now := time.Now()
		timestamp := now.Unix()
		t.Run("When the position is created", func(t *testing.T) {
			position, err := NewPosition(rideID, lat, lon, timestamp)
			t.Run("Then it is properly created", func(t *testing.T) {
				assertThat.NoError(err)
				assertThat.Equal(rideID, position.RideID())
				assertThat.Equal(lat, position.Coordinates().Lat())
				assertThat.Equal(lon, position.Coordinates().Lon())
				assertThat.Equal(now.Format(time.RFC3339), position.Date().Format(time.RFC3339))
			})
		})
	})
}
