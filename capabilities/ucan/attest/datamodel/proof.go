package datamodel

import "github.com/ipfs/go-cid"

type ProofArgumentsModel struct {
	Proof cid.Cid `cborgen:"proof" dagjsongen:"proof"`
}
