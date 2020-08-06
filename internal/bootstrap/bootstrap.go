package bootstrap

import (
	"context"
	"errors"

	"golang.org/x/sync/errgroup"
)

type Config struct {
	CSV struct {
		InputFilepath  string `mapstructure:"input_filepath"`
		OutputFilepath string `mapstructure:"output_filepath"`
	} `mapstructure:"csv"`
}

func (c Config) Validate() error {
	if c.CSV.InputFilepath == "" {
		return errors.New("missing input filepath variable")
	}
	if c.CSV.OutputFilepath == "" {
		return errors.New("missing output filepath variable")
	}

	return nil
}

func Run(ctx context.Context, cfg Config) error {
	err := cfg.Validate()
	if err != nil {
		return err
	}
	writer, err := initCSVWriter(ctx, cfg)
	if err != nil {
		return err
	}

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

	return writer.Write(ctx)
}
