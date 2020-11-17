module github.com/stack-labs/stack-rpc-plugins/service/gateway

go 1.14

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
	github.com/stack-labs/stack-rpc v1.0.0 => ../../../stack-rpc
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/mux v1.7.4
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/nats-io/nats-server/v2 v2.1.9 // indirect
	github.com/soheilhy/cmux v0.1.4 // indirect
	github.com/stack-labs/stack-rpc v1.0.0
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200427203606-3cfed13b9966 // indirect
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	go.etcd.io/bbolt v1.3.5 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)
