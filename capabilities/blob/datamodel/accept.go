package datamodel

import (
	"github.com/alanshaw/ucantone/ucan"
	"github.com/alanshaw/ucantone/ucan/promise"
)

type AcceptArgumentsModel struct {
	Blob BlobModel       `cborgen:"blob" dagjsongen:"blob"`
	Put  promise.AwaitOK `cborgen:"_put" dagjsongen:"_put"`
}

type AcceptOKModel struct {
	Site ucan.Link `cborgen:"site" dagjsongen:"site"`
}
