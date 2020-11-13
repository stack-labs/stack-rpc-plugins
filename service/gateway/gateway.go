package gateway

import (
	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc-plugins/service/gateway/api"
	"github.com/stack-labs/stack-rpc-plugins/service/gateway/plugin"
	"github.com/stack-labs/stack-rpc/pkg/cli"
	"github.com/stack-labs/stack-rpc/util/log"
)

func Run(svc stack.Service) {
	c := svc.Options().Cmd
	app := c.App()

	// gateway options
	opts := api.Options()
	svc.Init(opts...)

	before := app.Before

	var ctx *cli.Context

	// run gateway
	app.Before = func(c *cli.Context) error {
		ctx = c
		return before(c)
	}

	// after stack service start run api gateway
	svc.Init(stack.AfterStart(func() error {
		opts, err := api.Run(ctx, svc)
		if err != nil {
			return err
		}

		svc.Init(opts...)
		return nil
	}))

	// plugin tags
	plugins := plugin.Plugins()
	for _, p := range plugins {
		log.Infof("plugin: %s", p.String())
		if flags := p.Flags(); len(flags) > 0 {
			log.Infof("flags: %+#s", flags)
			svc.Init(stack.Flags(flags...))
		}
	}

	return
}
