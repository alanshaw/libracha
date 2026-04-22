package testutil

import (
	crand "crypto/rand"
	"testing"

	"github.com/alanshaw/ucantone/did"
	"github.com/alanshaw/ucantone/principal/ed25519"
	"github.com/ipfs/go-cid"
	mh "github.com/multiformats/go-multihash"
)

func RandomBytes(t *testing.T, size int) []byte {
	bytes := make([]byte, size)
	_, _ = crand.Read(bytes)
	return bytes
}

func RandomCID(t *testing.T) cid.Cid {
	return cid.NewCidV1(cid.Raw, RandomMultihash(t))
}

func RandomDID(t *testing.T) did.DID {
	return RandomSigner(t).DID()
}

func RandomMultihash(t *testing.T) mh.Multihash {
	bytes := RandomBytes(t, 10)
	return Must(mh.Sum(bytes, mh.SHA2_256, -1))(t)
}

func RandomSigner(t *testing.T) ed25519.Signer {
	return Must(ed25519.Generate())(t)
}
