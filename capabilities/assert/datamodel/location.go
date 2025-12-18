package datamodel

import (
	"github.com/alanshaw/libracha/capabilities"
	"github.com/alanshaw/ucantone/did"
	"github.com/multiformats/go-multihash"
)

type LocationArgumentsModel struct {
	Space    did.DID                `cborgen:"space"`
	Content  multihash.Multihash    `cborgen:"content"`
	Location []capabilities.CborURL `cborgen:"location"`
	Range    *RangeModel            `cborgen:"range,omitempty"`
}

type RangeModel struct {
	Offset uint64  `cborgen:"offset"`
	Length *uint64 `cborgen:"length,omitempty"`
}
