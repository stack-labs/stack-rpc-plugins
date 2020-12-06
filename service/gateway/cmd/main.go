package main

import (
	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc-plugins/service/gateway/api"
	gwServer "github.com/stack-labs/stack-rpc-plugins/service/gateway/server"
	"github.com/stack-labs/stack-rpc/plugin"
	"github.com/stack-labs/stack-rpc/server/mock"
	"github.com/stack-labs/stack-rpc/util/log"
)

func init() {
	plugin.DefaultServers["mock"] = mock.NewServer
}

func main() {
	svc := stack.NewService(stack.Name("stack.rpc.greeter"))

	// gateway server
	_ = svc.Init(api.Options()...)
	apiServer := api.NewServer(svc)
	svc.Init(
		stack.Server(
			gwServer.NewServer(gwServer.APIServer(apiServer)),
		),
	)

	// run service
	if err := svc.Run(); err != nil {
		log.Fatal(err)
	}
}
