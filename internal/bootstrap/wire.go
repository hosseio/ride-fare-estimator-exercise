//+build wireinject

package bootstrap

import (
	"context"

	"github.com/hosseio/ride-fare-estimator-exercise/internal/reader"

	"github.com/hosseio/ride-fare-estimator-exercise/internal"

	cromberbus "github.com/chiguirez/cromberbus/v2"
	"github.com/google/wire"
	"github.com/hosseio/ride-fare-estimator-exercise/internal/creator"
	"github.com/hosseio/ride-fare-estimator-exercise/internal/storage"
	"github.com/hosseio/ride-fare-estimator-exercise/io"
)

var creatorSet = wire.NewSet(
	creator.NewCreatePositionCommandHandler,
)

var storageSet = wire.NewSet(
	getInMemoryStorage,
)

var inMemStorage *storage.InMemory

func getInMemoryStorage() *storage.InMemory {
	if inMemStorage != nil {
		return inMemStorage
	}

	inMemStorage = storage.NewInMemory()

	return inMemStorage
}

var ioSet = wire.NewSet(
	io.NewController,
	io.NewCSVReader,
	getDemuxer,
	getCSVFilepath,
	io.NewCSVWriter,
	getCSVOutputFilepath,
)

var readerSet = wire.NewSet(
	reader.NewFareRetriever,
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

func getCSVOutputFilepath(cfg Config) io.CSVOutFilepath {
	return io.CSVOutFilepath(cfg.CSV.OutputFilepath)
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
		wire.Bind(new(internal.RideRepository), new(*storage.InMemory)),
		getBus,
	)

	return io.Controller{}, nil
}

func initCSVWriter(ctx context.Context, cfg Config) (io.CSVWriter, error) {
	wire.Build(
		ioSet,
		readerSet,
		storageSet,
		wire.Bind(new(internal.RideView), new(*storage.InMemory)),
	)

	return io.CSVWriter{}, nil
}
