module github.com/stack-labs/stack-rpc-plugins/service/gateway

go 1.14

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
	github.com/stack-labs/stack-rpc v1.0.0 => ../../../stack-rpc
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.7.4
	github.com/stack-labs/stack-rpc v1.0.0-rc2
	github.com/stretchr/testify v1.4.0
)
