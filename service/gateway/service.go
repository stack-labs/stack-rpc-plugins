package gateway

import (
	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc-plugins/service/gateway/api"
	"github.com/stack-labs/stack-rpc-plugins/service/gateway/plugin"
	"github.com/stack-labs/stack-rpc/util/log"
)

func Hook(svc stack.Service) {
	apiServer := api.NewServer(svc)

	// gateway options
	_ = svc.Init(api.Options()...)

	// gateway hook
	_ = svc.Init(
		stack.AfterStart(apiServer.Start),
		stack.AfterStop(apiServer.Stop),
	)

	// plugin tags
	plugins := plugin.Plugins()
	for _, p := range plugins {
		log.Debugf("plugin: %s", p.String())
		if flags := p.Flags(); len(flags) > 0 {
			log.Debugf("flags: %+#s", flags)
			_ = svc.Init(stack.Flags(flags...))
		}
	}

	return
}
