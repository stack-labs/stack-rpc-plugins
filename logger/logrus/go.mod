module github.com/stack-labs/stack-rpc-plugins/logger/logrus

go 1.14

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.3
	github.com/stack-labs/stack-rpc v1.0.0 => ../../../stack-rpc
)

require (
	github.com/sirupsen/logrus v1.4.2
	github.com/stack-labs/stack-rpc v1.0.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)
