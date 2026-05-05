package datamodel

import (
	"github.com/fil-forge/ucantone/ucan"
	"github.com/fil-forge/ucantone/ucan/promise"
)

// +libracha:map-encoders
type AcceptArgumentsModel struct {
	Blob BlobModel       `cborgen:"blob" dagjsongen:"blob"`
	Put  promise.AwaitOK `cborgen:"_put" dagjsongen:"_put"`
}

type AcceptOKModel struct {
	Site ucan.Link `cborgen:"site" dagjsongen:"site"`
}
