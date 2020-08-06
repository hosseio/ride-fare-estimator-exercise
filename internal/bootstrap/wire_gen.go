// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package bootstrap

import (
	"context"
	"github.com/chiguirez/cromberbus/v2"
	"github.com/google/wire"
	"gitlab.emobg.tech/go/one-connected-fleet/Collision/internal/creator"
	"gitlab.emobg.tech/go/one-connected-fleet/Collision/internal/storage"
	"gitlab.emobg.tech/go/one-connected-fleet/Collision/io"
)

// Injectors from wire.go:

func initCSVReader(ctx context.Context, cfg Config) (io.CSVReader, error) {
	ioDemuxer := getDemuxer(ctx, cfg)
	csvFilepath := getCSVFilepath(cfg)
	csvReader := io.NewCSVReader(ioDemuxer, csvFilepath)
	return csvReader, nil
}

func initController(ctx context.Context, cfg Config) (io.Controller, error) {
	ioDemuxer := getDemuxer(ctx, cfg)
	inMemoryRideRepository := storage.NewInMemoryRideRepository()
	createPositionCommandHandler := creator.NewCreatePositionCommandHandler(inMemoryRideRepository)
	commandBus, err := getBus(createPositionCommandHandler)
	if err != nil {
		return io.Controller{}, err
	}
	controller := io.NewController(ioDemuxer, commandBus)
	return controller, nil
}

// wire.go:

var creatorSet = wire.NewSet(creator.NewCreatePositionCommandHandler)

var storageSet = wire.NewSet(storage.NewInMemoryRideRepository)

var ioSet = wire.NewSet(io.NewController, io.NewCSVReader, getDemuxer,
	getCSVFilepath,
)

var demuxer *io.Demuxer

func getDemuxer(ctx context.Context, cfg Config) *io.Demuxer {
	if demuxer != nil {
		return demuxer
	}

	demuxer = io.NewDemuxer()

	return demuxer
}

func getCSVFilepath(cfg Config) io.CSVFilepath {
	return io.CSVFilepath(cfg.CSV.InputFilepath)
}

var bus cromberbus.CommandBus

func getBus(creator2 creator.CreatePositionCommandHandler) (cromberbus.CommandBus, error) {
	if bus != nil {
		return bus, nil
	}
	handlerResolver := cromberbus.NewMapHandlerResolver()
	handlerResolver.AddHandler(creator2.Handle)

	bus = cromberbus.NewCromberBus(handlerResolver)

	return bus, nil
}
