package reader

import "github.com/hosseio/ride-fare-estimator-exercise/internal"

type Fare struct {
	RideID int
	Amount float64
}

type FareRetriever struct {
	rideView internal.RideView
}

func NewFareRetriever(rideView internal.RideView) FareRetriever {
	return FareRetriever{rideView: rideView}
}

func (r FareRetriever) Get() ([]Fare, error) {
	rides, err := r.rideView.All()
	if err != nil {
		return nil, err
	}

	var fares []Fare
	for _, ride := range rides {
		fare := Fare{
			RideID: ride.ID(),
			Amount: ride.FareEstimation(),
		}

		fares = append(fares, fare)
	}

	return fares, nil
}
