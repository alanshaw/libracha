package datamodel

import "github.com/ipfs/go-cid"

type ConcludeArgumentsModel struct {
	Receipt cid.Cid `cborgen:"receipt"`
}
