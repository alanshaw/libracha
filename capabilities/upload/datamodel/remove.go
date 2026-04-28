package datamodel

import (
	"github.com/ipfs/go-cid"
)

type RemoveArgumentsModel struct {
	Root cid.Cid `cborgen:"root" dagjsongen:"root"`
}

type RemoveOKModel struct{}
