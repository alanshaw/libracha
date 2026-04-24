package datamodel

import (
	"github.com/multiformats/go-multihash"
)

type BlobModel struct {
	Digest multihash.Multihash `cborgen:"digest"`
}

type RangeModel struct {
	Start uint64 `cborgen:"start"`
	End   uint64 `cborgen:"end"`
}

type RetrieveArgumentsModel struct {
	Blob  BlobModel  `cborgen:"blob"`
	Range RangeModel `cborgen:"range"`
}

type RetrieveOKModel struct{}
