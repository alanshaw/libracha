package didresolver_test

import (
	"testing"

	"github.com/fil-forge/libforge/didresolver"
	"github.com/fil-forge/ucantone/did"
	"github.com/fil-forge/ucantone/principal/ed25519"
	"github.com/fil-forge/ucantone/principal/signer"
	"github.com/stretchr/testify/require"
)

func TestSelfResolver(t *testing.T) {
	// A service identified by a did:web DID backed by an ed25519 key. Wrapping
	// the did:key signer makes it announce the did:web DID without changing how
	// it signs.
	didWeb, err := did.Parse("did:web:example.com")
	require.NoError(t, err)

	key, err := ed25519.Generate()
	require.NoError(t, err)

	self, err := signer.Wrap(key, didWeb)
	require.NoError(t, err)
	require.Equal(t, didWeb, self.DID())

	resolver := didresolver.NewSelfResolver(self)

	t.Run("resolves the service's own did:web without an HTTP request", func(t *testing.T) {
		verifier, err := resolver.Resolve(t.Context(), didWeb)
		require.NoError(t, err)

		// The resolved verifier announces the requested did:web — not the
		// underlying did:key — so token.VerifySignature's issuer-vs-verifier DID
		// equality check passes.
		require.Equal(t, didWeb, verifier.DID())

		// It is the service's real verifier: signatures from the signer verify.
		msg := []byte("hello")
		require.True(t, verifier.Verify(msg, self.Sign(msg)))
	})

	t.Run("does not resolve a different DID so a TieredResolver falls through", func(t *testing.T) {
		other, err := did.Parse("did:web:example.org")
		require.NoError(t, err)

		verifier, err := resolver.Resolve(t.Context(), other)
		require.Error(t, err)
		require.Nil(t, verifier)
		require.ErrorContains(t, err, "not the service's own DID")
	})

	t.Run("does not resolve the underlying did:key", func(t *testing.T) {
		// Only the wrapped did:web identity is served; the did:key the signer
		// wraps is a different DID and must not resolve.
		verifier, err := resolver.Resolve(t.Context(), key.DID())
		require.Error(t, err)
		require.Nil(t, verifier)
		require.ErrorContains(t, err, "not the service's own DID")
	})
}
