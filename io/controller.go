package io

import (
	"context"
	"log"

	cromberbus "github.com/chiguirez/cromberbus/v2"
	"gitlab.emobg.tech/go/one-connected-fleet/Collision/internal/creator"
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
		g.Go(func() error {
			for {
				select {
				case <-gCtx.Done():
					break
				case position := <-ch:
					err := c.bus.Dispatch(creator.CreatePositionCommand{
						RideID:    position.RideID,
						Lat:       position.Lat,
						Lon:       position.Lon,
						Timestamp: position.Timestamp,
					})

					log.Printf("error creating position: %s", err.Error())
				}
			}

			return nil
		})
	}

	return g.Wait()
}
