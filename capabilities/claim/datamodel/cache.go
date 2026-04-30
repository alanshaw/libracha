package datamodel

import (
	"github.com/ipfs/go-cid"
)

type CacheArgumentsModel struct {
	Claim    cid.Cid       `cborgen:"claim" dagjsongen:"claim"`
	Provider ProviderModel `cborgen:"provider" dagjsongen:"provider"`
}

type ProviderModel struct {
	Addresses [][]byte `cborgen:"addresses" dagjsongen:"addresses"`
}
