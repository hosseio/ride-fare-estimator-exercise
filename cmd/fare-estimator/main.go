package main

import (
	"context"

	"github.com/hosseio/ride-fare-estimator-exercise/internal/bootstrap"

	"github.com/chiguirez/snout"
)

func main() {
	kernel := snout.Kernel{
		RunE: bootstrap.Run,
	}

	kernelBootstrap := kernel.Bootstrap(
		"fare-estimator",
		&bootstrap.Config{},
	)

	if err := kernelBootstrap.Initialize(); err != nil {
		if err != context.Canceled {
			panic(err)
		}
	}
}
