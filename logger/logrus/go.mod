module github.com/stack-labs/stack-rpc-plugins/logger/logrus

go 1.14

replace github.com/stack-labs/stack-rpc v1.0.0 => ../../../stack-rpc

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/sirupsen/logrus v1.4.2
	github.com/stack-labs/stack-rpc v1.0.0
	gopkg.in/yaml.v2 v2.3.0
)
