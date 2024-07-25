package keeper

import (
	"github.com/reiver/go-conclient"
)

// Client is a client for a StakeWise Keeper contract on some blockchain-network.
type Client struct {
	conclient.Client
}

func MakeClient(keepercontract KeeperContract, rpcurl string) Client {
	var contract = conclient.Contract{
		ABI:             abi,
		Address:         keepercontract.Address,
		ChainID:         keepercontract.ChainID,
		FromBlockNumber: keepercontract.FromBlockNumber,
	}

	return Client{
		Client:conclient.MakeClient(contract, rpcurl),
	}
}
