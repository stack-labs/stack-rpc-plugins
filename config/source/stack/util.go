package stack

import (
	"time"

	proto "github.com/stack-labs/stack-rpc-plugins/config/source/stack/proto"
	"github.com/stack-labs/stack-rpc/pkg/config/source"
)

func toChangeSet(c *proto.ChangeSet) *source.ChangeSet {
	return &source.ChangeSet{
		Data:      c.Data,
		Checksum:  c.Checksum,
		Format:    c.Format,
		Timestamp: time.Unix(c.Timestamp, 0),
		Source:    c.Source,
	}
}
