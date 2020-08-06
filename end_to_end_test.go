//+build e2e

package ride_fare_estimator_exercise

import (
	"bufio"
	"context"
	"encoding/csv"
	"os"
	"testing"

	"github.com/hosseio/ride-fare-estimator-exercise/internal/bootstrap"

	"github.com/stretchr/testify/require"
)

func TestE2E(t *testing.T) {
	assertThat := require.New(t)
	var cfg bootstrap.Config
	ctx := context.Background()
	setup := func() {
		cfg.CSV.InputFilepath = "e2e_file.csv"
		cfg.CSV.OutputFilepath = "e2e_output.csv"

		bootstrap.Run(ctx, cfg)
	}
	tearDown := func() {
		_ = os.Remove(cfg.CSV.OutputFilepath)
	}
	t.Run("Given a csv file with positions of 9 rides (id 1-9) to be read", func(t *testing.T) {
		// e2e_file.csv
		t.Run("When the system is running", func(t *testing.T) {
			setup()
			defer tearDown()
			t.Run("Then an output file with the fares of all rides is written", func(t *testing.T) {
				file, err := os.Open(cfg.CSV.OutputFilepath)
				assertThat.NoError(err)
				reader := csv.NewReader(bufio.NewReader(file))

				expected := make(map[string]string)
				expected["1"] = "10.039141"
				expected["2"] = "11.799228"
				expected["3"] = "32.544254"
				expected["4"] = "1.350361"
				expected["5"] = "21.475829"
				expected["6"] = "8.114287"
				expected["7"] = "28.710878"
				expected["8"] = "7.908595"
				expected["9"] = "5.047124"

				lines := []string{}
				for {
					record, err := reader.Read()
					if err != nil {
						break
					}
					if err != nil {
						t.Fatal(err)
					}

					lines = append(lines, record[0])
					assertThat.Equal(expected[record[0]], record[1])
				}
			})
		})
	})
}
