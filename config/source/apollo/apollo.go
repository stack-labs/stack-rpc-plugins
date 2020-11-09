package apollo

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/magiconair/properties"
	"github.com/stack-labs/stack-rpc/config/source"
	log "github.com/stack-labs/stack-rpc/logger"
	apo "github.com/stack-labs/stack-rpc/plugins/config/source/apollo/agollo"
	apoC "github.com/stack-labs/stack-rpc/plugins/config/source/apollo/agollo/env/config"
)

type apolloSource struct {
	appID      string
	client     *apo.Client
	namespaces []string
	opts       source.Options
}

var (
	DefaultAppID          = "stack"
	DefaultAddr           = "http://127.0.0.1:8080"
	DefaultCluster        = "dev"
	DefaultIsBackupConfig = false
	DefaultNamespaces     = "application"
	DefaultSecret         = ""
)

func (a *apolloSource) Read() (set *source.ChangeSet, err error) {
	return read(a.namespaces, a.client)
}

func (a *apolloSource) Watch() (source.Watcher, error) {
	return newWatcher(a.namespaces, a.client)
}

// Write is unsupported
func (a *apolloSource) Write(cs *source.ChangeSet) error {
	return nil
}

func (a *apolloSource) String() string {
	return "apollo"
}

func read(ns []string, client *apo.Client) (set *source.ChangeSet, err error) {
	s := map[string]string{}
	set = &source.ChangeSet{}
	for _, namespace := range ns {
		cache := client.GetConfigCache(namespace)
		cache.Range(func(key, value interface{}) bool {
			s[fmt.Sprintf("%v", key)] = fmt.Sprintf("%v", value)
			return true
		})
	}

	p := properties.LoadMap(s)
	set.Data, _ = json.Marshal(p.Map())
	set.Checksum = set.Sum()
	set.Format = "json"
	set.Source = "file"

	if p == nil || p.Len() == 0 {
		err = fmt.Errorf("apollo data is nill, check the apollo error logs")
		log.Warn(err)
	}

	return
}

func NewSource(opts ...source.Option) source.Source {
	var options source.Options
	for _, o := range opts {
		o(&options)
	}

	appID := "stack"
	addr := DefaultAddr
	cluster := DefaultCluster
	namespaces := DefaultNamespaces
	secret := DefaultSecret

	if options.Context != nil {
		appIDTemp, ok := options.Context.Value(appIDKey{}).(string)
		if !ok {
			log.Errorf("appId is necessary")
		} else {
			appID = appIDTemp
		}
		clusterTemp, ok := options.Context.Value(clusterKey{}).(string)
		if ok {
			cluster = clusterTemp
		}
		addrTemp, ok := options.Context.Value(addrKey{}).(string)
		if ok {
			addr = addrTemp
		}
		namespaceTemp, ok := options.Context.Value(namespacesKey{}).(string)
		if ok {
			namespaces = namespaceTemp
		}

		secretTemp, ok := options.Context.Value(secretKey{}).(string)
		if ok {
			secret = secretTemp
		}
	}

	c := &apoC.AppConfig{
		AppID:          appID,
		Cluster:        cluster,
		IP:             addr,
		NamespaceName:  namespaces,
		IsBackupConfig: false,
		Secret:         secret,
	}

	client, err := apo.StartWithConfig(func() (*apoC.AppConfig, error) {
		return c, nil
	})

	if err != nil {
		log.Errorf("apollo client init error: %s", err)
	}

	return &apolloSource{
		appID:      appID,
		client:     client,
		namespaces: strings.Split(namespaces, ","),
		opts:       options,
	}
}
