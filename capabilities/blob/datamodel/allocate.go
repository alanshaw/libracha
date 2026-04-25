package datamodel

import (
	"github.com/alanshaw/libracha/capabilities"
	"github.com/alanshaw/ucantone/ucan"
)

type AllocateArgumentsModel struct {
	Blob  BlobModel `cborgen:"blob" dagjsongen:"blob"`
	Cause ucan.Link `cborgen:"cause" dagjsongen:"cause"`
}

type AllocateOKModel struct {
	Size    uint64            `cborgen:"size" dagjsongen:"size"`
	Address *BlobAddressModel `cborgen:"address,omitempty" dagjsongen:"address,omitempty"`
}

type BlobAddressModel struct {
	URL     capabilities.CborURL  `cborgen:"url" dagjsongen:"url"`
	Headers map[string]string     `cborgen:"headers" dagjsongen:"headers"`
	Expires capabilities.CborTime `cborgen:"expires" dagjsongen:"expires"`
}
