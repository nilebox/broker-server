package server

import (
	"context"

	"github.com/nilebox/broker-server/example/controller"
	"github.com/nilebox/broker-server/pkg/server"
	"go.uber.org/zap"
)

type ExampleServer struct {
	Addr string
}

func (b *ExampleServer) Run(ctx context.Context) (returnErr error) {
	log := ctx.Value("log").(*zap.Logger)
	_ = log

	c, err := controller.NewController(ctx)
	if err != nil {
		return err
	}

	return server.Run(ctx, b.Addr, c)
}
