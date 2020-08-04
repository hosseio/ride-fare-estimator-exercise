package main

import (
	"context"

	"github.com/chiguirez/snout"
	"gitlab.emobg.tech/go/one-connected-fleet/Collision/internal/bootstrap"
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
