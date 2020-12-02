// Package api is an API Gateway
package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc-plugins/service/gateway/handler"
	"github.com/stack-labs/stack-rpc-plugins/service/gateway/helper"
	"github.com/stack-labs/stack-rpc-plugins/service/gateway/plugin"
	ahandler "github.com/stack-labs/stack-rpc/api/handler"
	aapi "github.com/stack-labs/stack-rpc/api/handler/api"
	"github.com/stack-labs/stack-rpc/api/handler/event"
	ahttp "github.com/stack-labs/stack-rpc/api/handler/http"
	arpc "github.com/stack-labs/stack-rpc/api/handler/rpc"
	"github.com/stack-labs/stack-rpc/api/handler/web"
	"github.com/stack-labs/stack-rpc/api/resolver"
	"github.com/stack-labs/stack-rpc/api/resolver/grpc"
	"github.com/stack-labs/stack-rpc/api/resolver/host"
	"github.com/stack-labs/stack-rpc/api/resolver/path"
	rrstack "github.com/stack-labs/stack-rpc/api/resolver/stack"
	"github.com/stack-labs/stack-rpc/api/router"
	regRouter "github.com/stack-labs/stack-rpc/api/router/registry"
	"github.com/stack-labs/stack-rpc/api/server"
	"github.com/stack-labs/stack-rpc/api/server/acme"
	"github.com/stack-labs/stack-rpc/api/server/acme/autocert"
	httpapi "github.com/stack-labs/stack-rpc/api/server/http"
	"github.com/stack-labs/stack-rpc/pkg/cli"
	"github.com/stack-labs/stack-rpc/util/log"
)

type config struct {
	Address      string      `json:"address"`
	Handler      string      `json:"handler"`
	Resolver     string      `json:"resolver"`
	RPCPath      string      `json:"rpc_path"`
	APIPath      string      `json:"api_path"`
	ProxyPath    string      `json:"proxy_path"`
	Namespace    string      `json:"namespace"`
	HeaderPrefix string      `json:"header_prefix"`
	EnableRPC    bool        `json:"enable_rpc"`
	EnableACME   bool        `json:"enable_acme"`
	EnableTLS    bool        `json:"enable_tls"`
	ACME         *acmeConfig `json:"acme"`
	TLS          *helper.TLS `json:"tls"`
}

type acmeConfig struct {
	Provider          string   `json:"provider"`
	ChallengeProvider string   `json:"challenge_provider"`
	CA                string   `json:"ca"`
	Hosts             []string `json:"hosts"`
}

func newDefaultConfig() *config {
	return &config{
		Address:      ":8080",
		Handler:      "meta",
		Resolver:     "stack",
		RPCPath:      "/rpc",
		APIPath:      "/",
		ProxyPath:    "/{service:[a-zA-Z0-9]+}",
		Namespace:    "stack.rpc.api",
		HeaderPrefix: "X-Stack-",
		EnableRPC:    false,
		ACME: &acmeConfig{
			Provider:          "autocert",
			ChallengeProvider: "cloudflare",
			CA:                acme.LetsEncryptProductionCA,
		},
	}
}

// run api gateway
func Run(svc stack.Service) ([]stack.Option, error) {
	cfg := svc.Options().Config
	conf := newDefaultConfig()
	if cfg != nil {
		if c := cfg.Get("gateway"); c != nil {
			c.Scan(conf)
		}
	}

	// Init plugins
	for _, p := range plugin.Plugins() {
		p.Init(cfg)
	}

	// Init API
	var opts []server.Option

	if conf.EnableACME {
		opts = append(opts, server.EnableACME(true))
		opts = append(opts, server.ACMEHosts(conf.ACME.Hosts...))
		switch conf.ACME.Provider {
		case "autocert":
			opts = append(opts, server.ACMEProvider(autocert.New()))
		default:
			log.Fatalf("%s is not a valid ACME provider\n", conf.ACME.Provider)
		}
	} else if conf.EnableTLS {
		config, err := helper.TLSConfig(conf.TLS)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		opts = append(opts, server.EnableTLS(true))
		opts = append(opts, server.TLSConfig(config))
	}

	// create the router
	var h http.Handler
	r := mux.NewRouter()
	h = r

	// return version and list of services
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		helper.ServeCORS(w, r)

		if r.Method == "OPTIONS" {
			return
		}

		// TODO index custom
		response := fmt.Sprintf(`{"version": "%s"}`, svc.Server().Options().Version)
		w.Write([]byte(response))
	})

	// strip favicon.ico
	r.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})

	// srvOpts = append(srvOpts, stack.Name(Name))
	// if i := time.Duration(ctx.GlobalInt("register_ttl")); i > 0 {
	// 	srvOpts = append(srvOpts, stack.RegisterTTL(i*time.Second))
	// }
	// if i := time.Duration(ctx.GlobalInt("register_interval")); i > 0 {
	// 	srvOpts = append(srvOpts, stack.RegisterInterval(i*time.Second))
	// }

	// initialise svc
	// svc := stack.NewService(srvOpts...)
	// register rpc handler
	if conf.EnableRPC {
		log.Logf("Registering RPC Handler at %s", conf.RPCPath)
		r.Handle(conf.RPCPath, handler.NewRPCHandlerFunc(svc.Options()))
	}

	// resolver options
	ropts := []resolver.Option{
		resolver.WithNamespace(conf.Namespace),
		resolver.WithHandler(conf.Handler),
	}

	// default resolver
	rr := rrstack.NewResolver(ropts...)

	switch conf.Resolver {
	case "host":
		rr = host.NewResolver(ropts...)
	case "path":
		rr = path.NewResolver(ropts...)
	case "grpc":
		rr = grpc.NewResolver(ropts...)
	}

	switch conf.Handler {
	case "rpc":
		log.Logf("Registering API RPC Handler at %s", conf.APIPath)
		rt := regRouter.NewRouter(
			router.WithNamespace(conf.Namespace),
			router.WithHandler(arpc.Handler),
			router.WithResolver(rr),
			router.WithRegistry(svc.Options().Registry),
		)
		rp := arpc.NewHandler(
			ahandler.WithNamespace(conf.Namespace),
			ahandler.WithRouter(rt),
			ahandler.WithService(svc),
		)
		r.PathPrefix(conf.APIPath).Handler(rp)
	case "api":
		log.Logf("Registering API Request Handler at %s", conf.APIPath)
		rt := regRouter.NewRouter(
			router.WithNamespace(conf.Namespace),
			router.WithHandler(aapi.Handler),
			router.WithResolver(rr),
			router.WithRegistry(svc.Options().Registry),
		)
		ap := aapi.NewHandler(
			ahandler.WithNamespace(conf.Namespace),
			ahandler.WithRouter(rt),
			ahandler.WithService(svc),
		)
		r.PathPrefix(conf.APIPath).Handler(ap)
	case "event":
		log.Logf("Registering API Event Handler at %s", conf.APIPath)
		rt := regRouter.NewRouter(
			router.WithNamespace(conf.Namespace),
			router.WithHandler(event.Handler),
			router.WithResolver(rr),
			router.WithRegistry(svc.Options().Registry),
		)
		ev := event.NewHandler(
			ahandler.WithNamespace(conf.Namespace),
			ahandler.WithRouter(rt),
			ahandler.WithService(svc),
		)
		r.PathPrefix(conf.APIPath).Handler(ev)
	case "http", "proxy":
		log.Logf("Registering API HTTP Handler at %s", conf.ProxyPath)
		rt := regRouter.NewRouter(
			router.WithNamespace(conf.Namespace),
			router.WithHandler(ahttp.Handler),
			router.WithResolver(rr),
			router.WithRegistry(svc.Options().Registry),
		)
		ht := ahttp.NewHandler(
			ahandler.WithNamespace(conf.Namespace),
			ahandler.WithRouter(rt),
			ahandler.WithService(svc),
		)
		r.PathPrefix(conf.ProxyPath).Handler(ht)
	case "web":
		log.Logf("Registering API Web Handler at %s", conf.APIPath)
		rt := regRouter.NewRouter(
			router.WithNamespace(conf.Namespace),
			router.WithHandler(web.Handler),
			router.WithResolver(rr),
			router.WithRegistry(svc.Options().Registry),
		)
		w := web.NewHandler(
			ahandler.WithNamespace(conf.Namespace),
			ahandler.WithRouter(rt),
			ahandler.WithService(svc),
		)
		r.PathPrefix(conf.APIPath).Handler(w)
	default:
		log.Logf("Registering API Default Handler at %s", conf.APIPath)
		rt := regRouter.NewRouter(
			router.WithNamespace(conf.Namespace),
			router.WithResolver(rr),
			router.WithRegistry(svc.Options().Registry),
		)
		r.PathPrefix(conf.APIPath).Handler(handler.Meta(svc, rt))
	}

	// reverse wrap handler
	plugins := append(plugin.Plugins(), plugin.Plugins()...)
	for i := len(plugins); i > 0; i-- {
		h = plugins[i-1].Handler()(h)
	}

	// create the server
	api := httpapi.NewServer(conf.Address)
	api.Init(opts...)
	api.Handle("/", h)

	// Start API
	if err := api.Start(); err != nil {
		log.Error(err)
		return nil, err
	}

	options := []stack.Option{
		stack.AfterStop(func() error {
			log.Infof("api stop")
			return api.Stop()
		}),
	}
	return options, nil
}

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
