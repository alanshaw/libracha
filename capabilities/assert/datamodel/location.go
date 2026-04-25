package datamodel

import (
	"github.com/alanshaw/libracha/capabilities"
	"github.com/alanshaw/ucantone/did"
	"github.com/multiformats/go-multihash"
)

type LocationArgumentsModel struct {
	Space    did.DID                `cborgen:"space" dagjsongen:"space"`
	Content  multihash.Multihash    `cborgen:"content" dagjsongen:"content"`
	Location []capabilities.CborURL `cborgen:"location" dagjsongen:"location"`
	Range    *RangeModel            `cborgen:"range,omitempty" dagjsongen:"range,omitempty"`
}

type RangeModel struct {
	Offset uint64  `cborgen:"offset" dagjsongen:"offset"`
	Length *uint64 `cborgen:"length,omitempty" dagjsongen:"length,omitempty"`
}

type LocationOKModel struct{}
