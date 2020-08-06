//+build wireinject

package bootstrap

import (
	"context"

	"gitlab.emobg.tech/go/one-connected-fleet/Collision/internal"

	cromberbus "github.com/chiguirez/cromberbus/v2"
	"github.com/google/wire"
	"gitlab.emobg.tech/go/one-connected-fleet/Collision/internal/creator"
	"gitlab.emobg.tech/go/one-connected-fleet/Collision/internal/storage"
	"gitlab.emobg.tech/go/one-connected-fleet/Collision/io"
)

var creatorSet = wire.NewSet(
	creator.NewCreatePositionCommandHandler,
)

var storageSet = wire.NewSet(
	storage.NewInMemoryRideRepository,
)

var ioSet = wire.NewSet(
	io.NewController,
	io.NewCSVReader,
	getDemuxer,
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

func getBus(creator creator.CreatePositionCommandHandler) (cromberbus.CommandBus, error) {
	if bus != nil {
		return bus, nil
	}
	handlerResolver := cromberbus.NewMapHandlerResolver()
	handlerResolver.AddHandler(creator.Handle)

	bus = cromberbus.NewCromberBus(handlerResolver)

	return bus, nil
}

func initCSVReader(ctx context.Context, cfg Config) (io.CSVReader, error) {
	wire.Build(
		ioSet,
	)

	return io.CSVReader{}, nil
}

func initController(ctx context.Context, cfg Config) (io.Controller, error) {
	wire.Build(
		ioSet,
		creatorSet,
		storageSet,
		wire.Bind(new(internal.RideRepository), new(*storage.InMemoryRideRepository)),
		getBus,
	)

	return io.Controller{}, nil
}
