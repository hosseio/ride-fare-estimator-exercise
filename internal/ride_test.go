package internal

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestRide_AddPosition(t *testing.T) {
	assertThat := require.New(t)
	t.Run("Given a ride", func(t *testing.T) {
		ride, _ := NewRide(1)
		t.Run("When adding 5 positions with only 3 valid segments", func(t *testing.T) {
			now := time.Now()
			start := now.Add(-1 * time.Hour)
			p1, _ := NewPosition(37.389098, -5.984379, start.Unix())
			ride.AddPosition(p1)
			p2, _ := NewPosition(37.888339, -4.779336, now.Unix())
			ride.AddPosition(p2) // about 108km make in one hour: this one should be avoided
			p3, _ := NewPosition(37.389277, -5.984701, start.Add(5*time.Second).Unix())
			ride.AddPosition(p3)
			p4, _ := NewPosition(37.389391, -5.985022, start.Add(10*time.Second).Unix())
			ride.AddPosition(p4)
			p5, _ := NewPosition(37.356303, -5.981766, start.Add(11*time.Second).Unix())
			ride.AddPosition(p5) // about 4km in 1 second: this one should be avoided

			t.Run("Then the final ride has 2 segments", func(t *testing.T) {
				assertThat.Equal(2, ride.segments.Len())
			})
			t.Run("And the current position is the last one accepted", func(t *testing.T) {
				assertThat.Equal(37.389391, ride.currentPosition.Coordinates().Lat()) // p4
				assertThat.Equal(-5.985022, ride.currentPosition.Coordinates().Lon())
			})
		})
	})
}
