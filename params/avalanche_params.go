// (c) 2024-2029, GB Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package params

import (
	"math/big"

	"github.com/geeblock/geego/utils/units"
)

// Minimum Gas Price
const (
	// MinGasPrice is the number of nAVAX required per gas unit for a
	// transaction to be valid, measured in wei
	LaunchMinGasPrice        int64 = 470 * GWei
	ApricotPhase1MinGasPrice int64 = 225 * GWei

	AvalancheAtomicTxFee = units.MilliGee

	ApricotPhase1GasLimit uint64 = 8_000_000
	CortinaGasLimit       uint64 = 15_000_000

	ApricotPhase3MinBaseFee               int64  = 750 * GWei
	ApricotPhase3MaxBaseFee               int64  = 2250 * GWei
	ApricotPhase3InitialBaseFee           int64  = 2250 * GWei
	ApricotPhase3TargetGas                uint64 = 10_000_000
	ApricotPhase4MinBaseFee               int64  = 250 * GWei
	ApricotPhase4MaxBaseFee               int64  = 10_000 * GWei
	ApricotPhase4BaseFeeChangeDenominator uint64 = 120
	ApricotPhase5TargetGas                uint64 = 15_000_000
	ApricotPhase5BaseFeeChangeDenominator uint64 = 360
	EtnaMinBaseFee                        int64  = GWei

	DynamicFeeExtraDataSize        = 80
	RollupWindow            uint64 = 10

	// The base cost to charge per atomic transaction. Added in Apricot Phase 5.
	AtomicTxBaseCost uint64 = 10_000
)

// The atomic gas limit specifies the maximum amount of gas that can be consumed by the atomic
// transactions included in a block and is enforced as of ApricotPhase5. Prior to ApricotPhase5,
// a block included a single atomic transaction. As of ApricotPhase5, each block can include a set
// of atomic transactions where the cumulative atomic gas consumed is capped by the atomic gas limit,
// similar to the block gas limit.
//
// This value must always remain <= MaxUint64.
var AtomicGasLimit *big.Int = big.NewInt(100_000)
