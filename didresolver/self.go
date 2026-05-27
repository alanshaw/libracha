package didresolver

import (
	"context"
	"fmt"

	"github.com/fil-forge/ucantone/did"
	"github.com/fil-forge/ucantone/principal"
	"github.com/fil-forge/ucantone/ucan"
)

type SelfResolver struct {
	self     did.DID
	verifier ucan.Verifier
}

func (r *SelfResolver) Resolve(_ context.Context, input did.DID) (ucan.Verifier, error) {
	if input != r.self {
		return nil, fmt.Errorf("not the service's own DID")
	}
	return r.verifier, nil
}

// NewSelfResolver returns a DID resolver tier that satisfies requests for the
// service's own DID using the in-memory identity. Returns an error for any
// other DID so a [TieredResolver] falls through to the next tier.
func NewSelfResolver(id principal.Signer) *SelfResolver {
	return &SelfResolver{
		self:     id.DID(),
		verifier: id.Verifier(),
	}
}
