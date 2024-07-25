package keeper

import (
	"github.com/reiver/go-conclient"
)

// Client is a client for a StakeWise Keeper contract on some blockchain-network.
type Client struct {
	conclient.Client
}

func MakeClient(contract Contract, rpcurl string) Client {
	return Client{
		Client:conclient.MakeClient(contract, rpcurl),
	}
}
