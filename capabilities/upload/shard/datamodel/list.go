package datamodel

import "github.com/ipfs/go-cid"

type ListArgumentsModel struct {
	Cursor *string `cborgen:"cursor,omitempty" dagjsongen:"cursor,omitempty"`
	Size   *int64  `cborgen:"size,omitempty" dagjsongen:"size,omitempty"`
}

type ListOKModel struct {
	Cursor  *string   `cborgen:"cursor,omitempty" dagjsongen:"cursor,omitempty"`
	Size    int64     `cborgen:"size" dagjsongen:"size"`
	Results []cid.Cid `cborgen:"results" dagjsongen:"results"`
}
