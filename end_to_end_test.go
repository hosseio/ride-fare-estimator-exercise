//+build e2e

package ride_fare_estimator_exercise

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestE2E(t *testing.T) {
	assertThat := require.New(t)
	t.Run("Given a csv file with paths to be read", func(t *testing.T) {
		t.Run("When the system is running", func(t *testing.T) {
			t.Run("Then ", func(t *testing.T) {
				assertThat.True(false)
			})
		})
	})
}
