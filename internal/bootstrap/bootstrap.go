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
	cCtx, cancelFn := context.WithCancel(ctx)
	g, gCtx := errgroup.WithContext(cCtx)

	controller, err := initController(gCtx, cfg)
	if err != nil {
		return err
	}
	g.Go(func() error {
		return controller.Start(gCtx)
	})

	csvReader, err := initCSVReader(gCtx, cfg)
	if err != nil {
		return err
	}
	g.Go(func() error {
		err := csvReader.Read(gCtx)
		cancelFn()
		return err
	})

	err = g.Wait()
	if err != nil {
		return err
	}

	// TODO write output

	return nil
}
