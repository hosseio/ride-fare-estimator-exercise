package bootstrap

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type Config struct {

}

func Run(ctx context.Context, cfg Config) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return nil
	})

	return g.Wait()
}
