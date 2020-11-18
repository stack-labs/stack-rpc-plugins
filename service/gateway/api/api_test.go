package api

import (
	"reflect"
	"testing"

	"github.com/stack-labs/stack-rpc"
)

func TestRun(t *testing.T) {
	type args struct {
		svc stack.Service
	}
	tests := []struct {
		name    string
		args    args
		want    []stack.Option
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Run(tt.args.svc)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() got = %v, want %v", got, tt.want)
			}
		})
	}
}
