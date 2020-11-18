package gateway

import (
	"testing"

	"github.com/stack-labs/stack-rpc"
)

func TestRun(t *testing.T) {
	type args struct {
		svc stack.Service
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
