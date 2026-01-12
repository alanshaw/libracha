package receipt

import (
	"context"
	"fmt"
	"iter"
	"net/http"

	"github.com/alanshaw/ucantone/ipld/codec/dagcbor"
	"github.com/alanshaw/ucantone/ucan"
	"github.com/alanshaw/ucantone/ucan/container"
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("receipt")

type Finder interface {
	// FindByTask returns all UCAN containers that have tokens related to the
	// given task i.e. invocations or receipts for the task.
	FindByTask(ctx context.Context, task cid.Cid) iter.Seq2[ucan.Container, error]
}

func NewHandler(tokens Finder) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		taskLink, err := cid.Parse(r.PathValue("task"))
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid task CID: %v", err), http.StatusBadRequest)
			return
		}

		var invocations []ucan.Invocation
		var delegations []ucan.Delegation
		var receipts []ucan.Receipt
		for c, err := range tokens.FindByTask(r.Context(), taskLink) {
			if err != nil {
				http.Error(w, fmt.Sprintf("failed to find receipt token: %v", err), http.StatusInternalServerError)
				return
			}
			invocations = append(invocations, c.Invocations()...)
			delegations = append(delegations, c.Delegations()...)
			receipts = append(receipts, c.Receipts()...)
		}

		out := container.New(
			container.WithInvocations(invocations...),
			container.WithDelegations(delegations...),
			container.WithReceipts(receipts...),
		)

		w.Header().Set("Content-Type", dagcbor.ContentType)
		err = out.MarshalCBOR(w)
		if err != nil {
			log.Errorw("marshaling receipt container", "error", err)
		}
	})
}
