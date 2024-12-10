// (c) 2024-2029, GB Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"github.com/geeblock/geeblockgo/api/metrics"
	"github.com/geeblock/geeblockgo/ids"
	"github.com/geeblock/geeblockgo/snow"
	"github.com/geeblock/geeblockgo/snow/validators/validatorstest"
	"github.com/geeblock/geeblockgo/utils/crypto/bls"
	"github.com/geeblock/geeblockgo/utils/logging"
)

func TestSnowContext() *snow.Context {
	sk, err := bls.NewSecretKey()
	if err != nil {
		panic(err)
	}
	pk := bls.PublicFromSecretKey(sk)
	return &snow.Context{
		NetworkID:      0,
		SubnetID:       ids.Empty,
		ChainID:        ids.Empty,
		NodeID:         ids.EmptyNodeID,
		PublicKey:      pk,
		Log:            logging.NoLog{},
		BCLookup:       ids.NewAliaser(),
		Metrics:        metrics.NewPrefixGatherer(),
		ChainDataDir:   "",
		ValidatorState: &validatorstest.State{},
	}
}
