package api

import (
	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc/pkg/cli"
)

// api stackway options
func Options() (options []stack.Option) {
	flags := []cli.Flag{
		cli.StringFlag{
			Name:   "stackway_name",
			Usage:  "Stackway name",
			EnvVar: "STACK_STACKWAY_NAME",
		},
		cli.StringFlag{
			Name:   "stackway_address",
			Usage:  "Set the stackway address e.g 0.0.0.0:8080",
			EnvVar: "STACK_STACKWAY_ADDRESS",
		},
		cli.StringFlag{
			Name:   "stackway_handler",
			Usage:  "Specify the request handler to be used for mapping HTTP requests to services; {api, event, http, rpc}",
			EnvVar: "STACK_STACKWAY_HANDLER",
		},
		cli.StringFlag{
			Name:   "stackway_namespace",
			Usage:  "Set the namespace used by the stackway e.g. stack.rpc.api",
			EnvVar: "STACK_STACKWAY_NAMESPACE",
		},
		cli.StringFlag{
			Name:   "stackway_resolver",
			Usage:  "Set the hostname resolver used by the stackway {host, path, grpc}",
			EnvVar: "STACK_STACKWAY_RESOLVER",
		},
		cli.BoolFlag{
			Name:   "stackway_enable_rpc",
			Usage:  "Enable call the backend directly via /rpc",
			EnvVar: "STACK_STACKWAY_ENABLE_RPC",
		},
	}

	options = append(options, stack.Flags(flags...))

	return
}
