package main

import (
	"net/http"

	"github.com/stack-labs/stack-rpc-plugins/service/gateway/plugin"
	"github.com/stack-labs/stack-rpc/config"
	"github.com/stack-labs/stack-rpc/pkg/cli"
	"github.com/stack-labs/stack-rpc/util/log"
)

func init() {
	err := plugin.Register(
		plugin.NewPlugin(
			plugin.WithName("example"),
			plugin.WithFlag(
				cli.StringFlag{
					Name:  "gateway_example_key",
					Usage: "gateway plugin flag",
					Value: "default",
				},
			),
			plugin.WithInit(func(cfg config.Config) error {
				conf := struct {
					Key string `json:"key"`
				}{}
				cfg.Get("gateway", "example").Scan(&conf)

				log.Infof("gateway plugin example init with key=%v", conf.Key)
				return nil
			}),
			plugin.WithHandler(func(h http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					log.Info("gateway plugin example")

					h.ServeHTTP(w, r)
				})
			}),
		),
	)
	if err != nil {
		log.Error(err)
	}
}
