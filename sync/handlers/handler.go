// (c) 2024-2029, GB Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package handlers

import (
	"github.com/geeblock/coreth/core/state/snapshot"
	"github.com/geeblock/coreth/core/types"
	"github.com/ethereum/go-ethereum/common"
)

type BlockProvider interface {
	GetBlock(common.Hash, uint64) *types.Block
}

type SnapshotProvider interface {
	Snapshots() *snapshot.Tree
}

type SyncDataProvider interface {
	BlockProvider
	SnapshotProvider
}
