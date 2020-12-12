// Package api is an API Gateway
package api

import (
	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc/pkg/cli"
)

// api gateway options
func Options() (options []stack.Option) {
	flags := []cli.Flag{
		cli.StringFlag{
			Name:   "gateway_name",
			Usage:  "Gateway name",
			EnvVar: "STACK_GATEWAY_NAME",
		},
		cli.StringFlag{
			Name:   "gateway_address",
			Usage:  "Set the gateway address e.g 0.0.0.0:8080",
			EnvVar: "STACK_GATEWAY_ADDRESS",
		},
		cli.StringFlag{
			Name:   "gateway_handler",
			Usage:  "Specify the request handler to be used for mapping HTTP requests to services; {api, event, http, rpc}",
			EnvVar: "STACK_GATEWAY_HANDLER",
		},
		cli.StringFlag{
			Name:   "gateway_namespace",
			Usage:  "Set the namespace used by the gateway e.g. com.example.gateway",
			EnvVar: "STACK_GATEWAY_NAMESPACE",
		},
		cli.StringFlag{
			Name:   "gateway_resolver",
			Usage:  "Set the hostname resolver used by the gateway {host, path, grpc}",
			EnvVar: "STACK_GATEWAY_RESOLVER",
		},
		cli.BoolFlag{
			Name:   "gateway_enable_rpc",
			Usage:  "Enable call the backend directly via /rpc",
			EnvVar: "STACK_GATEWAY_ENABLE_RPC",
		},
	}

	options = append(options, stack.Flags(flags...))

	return
}
