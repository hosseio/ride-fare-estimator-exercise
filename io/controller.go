package io

import (
	"context"
	"log"

	cromberbus "github.com/chiguirez/cromberbus/v2"
	"github.com/hosseio/ride-fare-estimator-exercise/internal/creator"
	"golang.org/x/sync/errgroup"
)

type Controller struct {
	demuxer *Demuxer
	bus     cromberbus.CommandBus
}

func NewController(demuxer *Demuxer, bus cromberbus.CommandBus) Controller {
	return Controller{demuxer: demuxer, bus: bus}
}

func (c Controller) Start(ctx context.Context) error {
	g, gCtx := errgroup.WithContext(ctx)

	for _, ch := range c.demuxer.inputs {
		g.Go(c.listen(gCtx, ch))
	}

	return g.Wait()
}

func (c Controller) listen(ctx context.Context, ch chan PositionDTO) func() error {
	return func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case position, ok := <-ch:
				if !ok {
					break
				}
				err := c.bus.Dispatch(creator.CreatePositionCommand{
					RideID:    position.RideID,
					Lat:       position.Lat,
					Lon:       position.Lon,
					Timestamp: position.Timestamp,
				})

				if err != nil {
					log.Printf("error creating position: %s", err.Error())
				}
			}
		}
	}
}
