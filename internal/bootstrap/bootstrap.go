package bootstrap

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type Config struct {
	CSV struct {
		InputFilepath  string `mapstructure:"input_filepath"`
		OutputFilepath string `mapstructure:"output_filepath"`
	} `mapstructure:"csv"`
}

func Run(ctx context.Context, cfg Config) error {
	g, ctx := errgroup.WithContext(ctx)

	controller, err := initController(ctx, cfg)
	if err != nil {
		return err
	}
	g.Go(func() error {
		return controller.Start(ctx)
	})

	csvReader, err := initCSVReader(ctx, cfg)
	if err != nil {
		return err
	}
	g.Go(func() error {
		return csvReader.Read(ctx)
	})

	return g.Wait()
}
