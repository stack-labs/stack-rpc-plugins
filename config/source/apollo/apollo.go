package apollo

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/magiconair/properties"
	apo "github.com/zouyx/agollo/v4"
	apoC "github.com/zouyx/agollo/v4/env/config"
	"github.com/stack-labs/stack-rpc/config/source"
	"github.com/stack-labs/stack-rpc/util/log"
)

type apolloSource struct {
	appID      string
	client     *apo.Client
	namespaces []string
	opts       source.Options
}

var (
	DefaultAppID            = "micro"
	DefaultAddr             = "http://127.0.0.1:8080"
	DefaultCluster          = "dev"
	DefaultIsBackupConfig   = true
	DefaultNamespaces       = "application"
	DefaultSecret           = ""
	DefaultBackupConfigPath = "/app/"
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

	return
}

func NewSource(opts ...source.Option) source.Source {
	var options source.Options
	for _, o := range opts {
		o(&options)
	}

	appID := "micro"
	addr := DefaultAddr
	cluster := DefaultCluster
	isBackupConfig := DefaultIsBackupConfig
	namespaces := DefaultNamespaces
	secret := DefaultSecret
	backupConfigPath := DefaultBackupConfigPath

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
		isBackupConfigTemp, ok := options.Context.Value(isBackupConfigKey{}).(bool)
		if ok {
			isBackupConfig = isBackupConfigTemp
		}
		secretTemp, ok := options.Context.Value(secretKey{}).(string)
		if ok {
			secret = secretTemp
		}
		backupConfigPathTemp, ok := options.Context.Value(backupConfigPathKey{}).(string)
		if ok {
			backupConfigPath = backupConfigPathTemp
		}
	}

	c := &apoC.AppConfig{
		AppID:            appID,
		Cluster:          cluster,
		IP:               addr,
		NamespaceName:    namespaces,
		IsBackupConfig:   isBackupConfig,
		Secret:           secret,
		BackupConfigPath: backupConfigPath,
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

type DefaultLogger struct {
}

func (d *DefaultLogger) Debugf(format string, params ...interface{}) {
	log.Debugf(format, params)
}

func (d *DefaultLogger) Infof(format string, params ...interface{}) {
	log.Infof(format, params)
}

func (d *DefaultLogger) Warnf(format string, params ...interface{}) {
	log.Warnf(format, params)
}

func (d *DefaultLogger) Errorf(format string, params ...interface{}) {
	log.Errorf(format, params)
}

func (d *DefaultLogger) Debug(v ...interface{}) {
	log.Debug(v)
}
func (d *DefaultLogger) Info(v ...interface{}) {
	log.Info(v)
}

func (d *DefaultLogger) Warn(v ...interface{}) {
	log.Warn(v)
}

func (d *DefaultLogger) Error(v ...interface{}) {
	log.Error(v)
}
