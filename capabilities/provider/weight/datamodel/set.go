package datamodel

import "github.com/alanshaw/ucantone/did"

type SetArgumentsModel struct {
	Provider did.DID `cborgen:"provider" dagjsongen:"provider"`
	Weight   uint64  `cborgen:"weight" dagjsongen:"weight"`
}
