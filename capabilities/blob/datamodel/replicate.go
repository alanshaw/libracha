package datamodel

import (
	"github.com/alanshaw/ucantone/ucan"
	"github.com/alanshaw/ucantone/ucan/promise"
)

type ReplicateArgumentsModel struct {
	// Blob is the blob that must be replicated.
	Blob BlobModel `cborgen:"blob"`
	// Replicas is the number of replicas to ensure.
	// e.g. Replicas: 3 will ensure 3 copies of the data exist in a network in total.
	Replicas uint64 `cborgen:"replicas"`
	// Site is a link to a location commitment indicating where the Blob must be
	// fetched from.
	Site ucan.Link `cborgen:"site"`
}

type ReplicateOKModel struct {
	// Site resolves to additional locations for the blob. They are links to
	// `/blob/replica/transfer` tasks.
	Site []promise.AwaitOK `cborgen:"site"`
}
