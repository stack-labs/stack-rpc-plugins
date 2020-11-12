package api

import (
	"reflect"
	"testing"

	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc/cli"
)

func TestCommands(t *testing.T) {
	type args struct {
		options []stack.Option
	}
	tests := []struct {
		name string
		args args
		want []cli.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Commands(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Commands() = %v, want %v", got, tt.want)
			}
		})
	}
}
