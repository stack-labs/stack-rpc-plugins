module github.com/stack-labs/stack-rpc-plugins/registry/zookeeper

go 1.14

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/go-zookeeper/zk v1.0.2
	github.com/google/uuid v1.1.2
	github.com/miekg/dns v1.1.31 // indirect
	github.com/mitchellh/hashstructure v1.0.0
	github.com/smartystreets/assertions v0.0.0-20180927180507-b2de0cb4f26d
	github.com/smartystreets/goconvey v1.6.4
	github.com/stack-labs/stack-rpc v1.0.0-rc2
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208 // indirect
)
