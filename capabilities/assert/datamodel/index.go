package datamodel

import "github.com/ipfs/go-cid"

type IndexArgumentsModel struct {
	Content cid.Cid `cborgen:"content" dagjsongen:"content"`
	Index   cid.Cid `cborgen:"index" dagjsongen:"index"`
}

type IndexOKModel struct{}
