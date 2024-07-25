package keeper

import (
	"github.com/reiver/go-ethaddr"
)

type KeeperContract struct {
	Address         ethaddr.Address
	ChainID         uint64
	FromBlockNumber uint64
}
