module github.com/stack-labs/stack-rpc-plugins/config/source/apollo

go 1.14

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
	github.com/stack-labs/stack-rpc v1.0.0 => ../../../../stack-rpc
)

require (
	github.com/magiconair/properties v1.8.4
	github.com/spf13/viper v1.7.1
	github.com/stack-labs/stack-rpc v1.0.0
	github.com/tevid/gohamcrest v1.1.1
)
