package internal

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_distanceInMeters(t *testing.T) {
	type args struct {
		position  Position
		position2 Position
	}
	tests := []struct {
		args args
		want float64
	}{
		{
			args{
				Position{
					coordinates: Coordinates{lat: 22.55, lon: 43.12}, // Rio de Janeiro, Brazil
				},
				Position{
					coordinates: Coordinates{lat: 13.45, lon: 100.28}, // Bangkok, Thailand
				},
			},
			6094544.408786774,
		},
		{
			args{Position{
				coordinates: Coordinates{lat: 20.10, lon: 57.30}, // Port Louis, Mauritius
			},
				Position{
					coordinates: Coordinates{lat: 0.57, lon: 100.21}, // Padang, Indonesia
				},
			},
			5145525.771394785,
		},
		{
			args{Position{
				coordinates: Coordinates{lat: 51.45, lon: 1.15}, // Oxford, United Kingdom
			},
				Position{
					coordinates: Coordinates{lat: 41.54, lon: 12.27}, // Vatican, City Vatican City
				},
			},
			1389179.3118293067,
		},
		{
			args{Position{
				coordinates: Coordinates{lat: 22.34, lon: 17.05}, // Windhoek, Namibia
			},
				Position{
					coordinates: Coordinates{lat: 51.56, lon: 4.29}, // Rotterdam, Netherlands
				},
			},
			3429893.10043882,
		},
		{
			args{Position{
				coordinates: Coordinates{lat: 63.24, lon: 56.59}, // Esperanza, Argentina
			},
				Position{
					coordinates: Coordinates{lat: 8.50, lon: 13.14}, // Luanda, Angola
				},
			},
			6996185.95539861,
		},
		{
			args{Position{
				coordinates: Coordinates{lat: 90.00, lon: 0.00}, // North/South Poles
			},
				Position{
					coordinates: Coordinates{lat: 48.51, lon: 2.21}, // Paris,  France
				},
			},
			4613477.506482742,
		},
		{
			args{Position{
				coordinates: Coordinates{lat: 45.04, lon: 7.42}, // Turin, Italy
			},
				Position{
					coordinates: Coordinates{lat: 3.09, lon: 101.42}, // Kuala Lumpur, Malaysia
				},
			},
			10078111.954385415,
		},
	}

	for _, tt := range tests {
		if got := distanceInMeters(tt.args.position, tt.args.position2); got != tt.want {
			t.Errorf("distanceInMeters() = %v, want %v", got, tt.want)
		}
	}
}

func TestSegment(t *testing.T) {
	assertThat := require.New(t)
	t.Run("Given and ending position before the initial one", func(t *testing.T) {
		now := time.Now()
		oneHourAgo := now.Add(-1 * time.Hour)
		p1, _ := NewPosition(37.5702221, -5.9412794, now.Unix())
		p2, _ := NewPosition(37.888339, -4.779336, oneHourAgo.Unix())
		t.Run("When the segment is created", func(t *testing.T) {
			_, err := NewSegmentFromPositions(p1, p2)
			t.Run("Then an error is returned", func(t *testing.T) {
				assertThat.Error(err)
				assertThat.Equal(ErrInvalidPositionTimes, err)
			})
		})
	})
	t.Run("Given 2 positions in two moments indicating a speed greater than 100 km per hour", func(t *testing.T) {
		now := time.Now()
		oneHourAgo := now.Add(-1 * time.Hour)
		// about 108 km made in one hour
		p1, _ := NewPosition(37.5702221, -5.9412794, oneHourAgo.Unix())
		p2, _ := NewPosition(37.888339, -4.779336, now.Unix())
		t.Run("When the segment is created", func(t *testing.T) {
			_, err := NewSegmentFromPositions(p1, p2)
			t.Run("Then an error is returned", func(t *testing.T) {
				assertThat.Error(err)
				assertThat.Equal(ErrTooMuchSpeed, err)
			})
		})
	})
	t.Run("Given 2 positions in two moments indicating a speed lower than 100 km per hour", func(t *testing.T) {
		now := time.Now()
		twoHoursAgo := now.Add(-2 * time.Hour)
		// about 108 km made in two hours
		p1, _ := NewPosition(37.5702221, -5.9412794, twoHoursAgo.Unix())
		p2, _ := NewPosition(37.888339, -4.779336, now.Unix())
		t.Run("When the segment is created", func(t *testing.T) {
			_, err := NewSegmentFromPositions(p1, p2)
			t.Run("Then an error is returned", func(t *testing.T) {
				assertThat.NoError(err)
			})
		})
	})
	t.Run("Given 2 positions for a moving segment about 1km and daytime", func(t *testing.T) {
		now := time.Now()
		initial := now.Add(-60 * time.Second)
		// about 108 km made in two hours
		p1, _ := NewPosition(52.363474, 4.875790, initial.Unix())
		p2, _ := NewPosition(52.359490, 4.862186, now.Unix())
		t.Run("When the segment is created", func(t *testing.T) {
			segment, err := NewSegmentFromPositions(p1, p2)
			t.Run("Then it is created", func(t *testing.T) {
				assertThat.NoError(err)
			})
			t.Run("And the fare is properly calculated", func(t *testing.T) {
				assertThat.Equal(0.7581297761001123, segment.fare)
			})
		})
	})
}
