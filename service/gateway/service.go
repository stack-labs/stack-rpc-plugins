package gateway

import (
	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc/util/log"

	"github.com/stack-labs/stack-rpc-plugins/service/gateway/api"
	"github.com/stack-labs/stack-rpc-plugins/service/gateway/plugin"
)

func Run(svc stack.Service) {
	// gateway options
	opts := api.Options()
	svc.Init(opts...)

	// after stack service start run api gateway
	svc.Init(stack.AfterStart(func() error {
		opts, err := api.Run(svc)
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
