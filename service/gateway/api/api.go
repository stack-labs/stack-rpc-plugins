package api

import (
	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc/pkg/cli"
)

type Server interface {
	Start() error
	Stop() error
}

// gateway api options
func Options() (options []stack.Option) {
	flags := []cli.Flag{
		cli.StringFlag{
			Name:   "gateway_handler",
			Usage:  "Specify the request handler to be used for mapping HTTP requests to services; {api, event, http, rpc}",
			EnvVar: "GATEWAY_HANDLER",
		},
		cli.StringFlag{
			Name:   "gateway_namespace",
			Usage:  "Set the namespace used by the gateway e.g. com.example.gateway",
			EnvVar: "GATEWAY_NAMESPACE",
		},
		cli.StringFlag{
			Name:   "gateway_resolver",
			Usage:  "Set the hostname resolver used by the gateway {host, path, grpc}",
			EnvVar: "GATEWAY_RESOLVER",
		},
		cli.BoolFlag{
			Name:   "gateway_enable_rpc",
			Usage:  "Enable call the backend directly via /rpc",
			EnvVar: "GATEWAY_ENABLE_RPC",
		},
	}

	options = append(options, stack.Flags(flags...))

	return
}
