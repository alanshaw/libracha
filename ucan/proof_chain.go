package ucanlib

import "github.com/alanshaw/ucantone/ucan"

type DelegationQuerier interface {
	// Query finds delegations matching the given audience, command, and subject.
	// Note: subject MUST not be nil. Matching delegations MAY include powerline
	// delegations (with nil subject) and delegations where command is a matching
	// parent of the passed command.
	Query(aud ucan.Principal, cmd ucan.Command, sub ucan.Subject) ([]ucan.Delegation, error)
}

// ProofChain recursively builds a proof chain of delegations from the given
// audience to the given subject for the specified command. It returns the list
// of delegations and their corresponding links.
func ProofChain(store DelegationQuerier, aud ucan.Principal, cmd ucan.Command, sub ucan.Principal) ([]ucan.Delegation, []ucan.Link, error) {
	proofs := []ucan.Delegation{}
	links := []ucan.Link{}

	matches, err := store.Query(aud, cmd, sub)
	if err != nil {
		return nil, nil, err
	}

	for _, d := range matches {
		if d.Subject() != nil && d.Subject().DID() == d.Issuer().DID() {
			proofs = append(proofs, d)
			links = append(links, d.Link())
			break
		}
		// if subject is nil, or subject != issuer, we need more proof
		ps, ls, err := ProofChain(store, d.Issuer(), d.Command(), sub)
		if err != nil {
			return nil, nil, err
		}
		if len(ps) == 0 {
			continue // try a different path
		}
		proofs = append(proofs, d)
		proofs = append(proofs, ps...)
		links = append(links, d.Link())
		links = append(links, ls...)
		break
	}

	return proofs, links, nil
}
