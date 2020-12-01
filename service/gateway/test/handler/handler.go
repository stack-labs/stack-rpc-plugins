package handler

import (
	"context"

	test "github.com/stack-labs/stack-rpc-plugins/service/gateway/test"
	proto "github.com/stack-labs/stack-rpc/api/proto"
)

var _ test.TestHandler = &Handler{}

type Handler struct {
}

// rpc模式handler
func (*Handler) Rpc(ctx context.Context, req *test.Request, resp *test.Response) error {
	resp.Msg = req.Msg

	return nil
}

// api模式handler
func (*Handler) Api(ctx context.Context, req *proto.Request, resp *proto.Response) error {
	if values := req.Get["msg"].GetValues(); len(values) > 0 {
		resp.Body = values[0]
	}
	return nil
}
