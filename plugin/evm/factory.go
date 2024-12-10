// (c) 2024-2029, GB Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"github.com/geeblock/geeblockgo/ids"
	"github.com/geeblock/geeblockgo/utils/logging"
	"github.com/geeblock/geeblockgo/vms"
)

var (
	// ID this VM should be referenced by
	ID = ids.ID{'e', 'v', 'm'}

	_ vms.Factory = &Factory{}
)

type Factory struct{}

func (*Factory) New(logging.Logger) (interface{}, error) {
	return &VM{}, nil
}
