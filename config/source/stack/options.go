package stack

import (
	"context"

	"github.com/stack-labs/stack-rpc/pkg/config/source"
)

type serviceNameKey struct{}
type pathKey struct{}

func ServiceName(a string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, serviceNameKey{}, a)
	}
}

// Path sets the key prefix to use
func Path(p string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, pathKey{}, p)
	}
}
