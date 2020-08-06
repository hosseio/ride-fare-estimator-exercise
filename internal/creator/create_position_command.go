package creator

import (
	"errors"
	"gitlab.emobg.tech/go/one-connected-fleet/Collision/internal"
)

type CreatePositionCommand struct {
	RideID    int
	Lat       float64
	Lon       float64
	Timestamp int64
}

type CreatePositionCommandHandler struct {
	repository internal.RideRepository
}

func NewCreatePositionCommandHandler(repository internal.RideRepository) CreatePositionCommandHandler {
	return CreatePositionCommandHandler{repository: repository}
}

func (h CreatePositionCommandHandler) Handle(command CreatePositionCommand) error {
	ride, err := h.repository.Get(command.RideID)
	if err != nil {
		if !errors.Is(err, internal.ErrRideNotFound) {
			return err
		}
		ride, err = internal.NewRide(command.RideID)
		if err != nil {
			return err
		}
	}

	position, err := internal.NewPosition(command.Lat, command.Lon, command.Timestamp)
	if err != nil {
		return err
	}

	err = ride.AddPosition(position)
	if err != nil {
		if !errors.Is(err, internal.ErrTooMuchSpeed) {
			return err
		}

		return nil
	}

	return h.repository.Save(ride)
}
