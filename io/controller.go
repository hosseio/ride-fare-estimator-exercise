package io

import (
	"context"
	cromberbus "github.com/chiguirez/cromberbus/v2"
	"gitlab.emobg.tech/go/one-connected-fleet/Collision/internal/creator"
	"golang.org/x/sync/errgroup"
	"sync/atomic"
	"time"
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
			keepOn := true
			for keepOn {
				select {
				case <-gCtx.Done():
					keepOn = false
					break
				case position := <-ch:
					err := c.bus.Dispatch(creator.CreatePositionCommand{
						RideID:    position.RideID,
						Lat:       position.Lat,
						Lon:       position.Lon,
						Timestamp: position.Timestamp,
					})
				}
			}
			return nil
		})
	}

	return g.Wait()
}
