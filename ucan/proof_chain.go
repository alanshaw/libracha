package ucanlib

import (
	"context"
	"iter"

	"github.com/alanshaw/ucantone/ucan"
	"github.com/alanshaw/ucantone/ucan/command"
)

type DelegationMatcher interface {
	// Match finds delegations matching the given audience, command, and subject.
	// Note: subject MUST not be nil. Matching delegations MAY include powerline
	// delegations (with nil subject) and delegations where command is a matching
	// parent of the passed command.
	Match(ctx context.Context, aud ucan.Principal, cmd ucan.Command, sub ucan.Subject) iter.Seq2[ucan.Delegation, error]
}

type DelegationFinder interface {
	// FindByAudienceCommandSubject retrieves delegations for the given audience,
	// command, and subject. Note: subject MAY be nil to indicate powerline.
	FindByAudienceCommandSubject(ctx context.Context, aud ucan.Principal, cmd ucan.Command, sub ucan.Subject) iter.Seq2[ucan.Delegation, error]
}

type FinderDelegationMatcher struct {
	finder DelegationFinder
}

// NewDelegationMatcher creates a simple delegation matcher that queries the
// passed finder to retrieve delegations matching the given audience, command,
// and subject.
func NewDelegationMatcher(finder DelegationFinder) *FinderDelegationMatcher {
	return &FinderDelegationMatcher{finder: finder}
}

func (gm *FinderDelegationMatcher) Match(ctx context.Context, aud ucan.Principal, cmd ucan.Command, sub ucan.Principal) iter.Seq2[ucan.Delegation, error] {
	return func(yield func(ucan.Delegation, error) bool) {
		cmdVariations := []ucan.Command{}
		segs := cmd.Segments()
		for i := len(segs) - 1; i >= 0; i-- {
			cmd := command.Top().Join(segs[0 : i+1]...)
			cmdVariations = append(cmdVariations, cmd)
		}
		cmdVariations = append(cmdVariations, command.Top())

		for _, cmd := range cmdVariations {
			for dlg, err := range gm.finder.FindByAudienceCommandSubject(ctx, aud, cmd, sub) {
				if err != nil {
					yield(nil, err)
					return
				}
				if !yield(dlg, nil) {
					return
				}
			}
			// try powerline
			// TODO: stop earily if we already found delegations?
			for dlg, err := range gm.finder.FindByAudienceCommandSubject(ctx, aud, cmd, nil) {
				if err != nil {
					yield(nil, err)
					return
				}
				if !yield(dlg, nil) {
					return
				}
			}
		}
	}
}

// ProofChain recursively builds a proof chain of delegations from the given
// audience to the given subject for the specified command. It returns the list
// of delegations and their corresponding links.
func ProofChain(ctx context.Context, matcher DelegationMatcher, aud ucan.Principal, cmd ucan.Command, sub ucan.Principal) ([]ucan.Delegation, []ucan.Link, error) {
	proofs := []ucan.Delegation{}
	links := []ucan.Link{}

	for d, err := range matcher.Match(ctx, aud, cmd, sub) {
		if err != nil {
			return nil, nil, err
		}
		if d.Subject() != nil && d.Subject().DID() == d.Issuer().DID() {
			proofs = append(proofs, d)
			links = append(links, d.Link())
			break
		}
		// if subject is nil, or subject != issuer, we need more proof
		ps, ls, err := ProofChain(ctx, matcher, d.Issuer(), d.Command(), sub)
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
