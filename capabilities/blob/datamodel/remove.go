package datamodel

import "github.com/multiformats/go-multihash"

type RemoveArgumentsModel struct {
	Digest multihash.Multihash `cborgen:"digest" dagjsongen:"digest"`
}
