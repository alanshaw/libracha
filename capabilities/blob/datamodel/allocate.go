package datamodel

import (
	"github.com/alanshaw/libracha/capabilities"
	"github.com/alanshaw/ucantone/ucan"
)

type AllocateArgumentsModel struct {
	Blob  BlobModel `cborgen:"blob"`
	Cause ucan.Link `cborgen:"cause"`
}

type AllocateOKModel struct {
	Size    uint64            `cborgen:"size"`
	Address *BlobAddressModel `cborgen:"address,omitempty"`
}

type BlobAddressModel struct {
	URL     capabilities.CborURL  `cborgen:"url"`
	Headers map[string]string     `cborgen:"headers"`
	Expires capabilities.CborTime `cborgen:"expires"`
}
