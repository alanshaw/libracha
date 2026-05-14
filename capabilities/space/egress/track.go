package egress

import (
	edm "github.com/fil-forge/libforge/capabilities/space/egress/datamodel"
	"github.com/fil-forge/ucantone/validator/bindcap"
)

const TrackCommand = "/space/egress/track"

type (
	TrackArguments = edm.TrackArgumentsModel
	TrackOK        = edm.TrackOKModel
)

// Track is the capability a storage node invokes to ask the egress
// tracking service to record egress accounted for in a batch of
// `/content/retrieve` receipts. The tracking service responds by forking
// a `/space/egress/consolidate` sub-invocation onto the receipt's
// effects; the typed OK return is empty.
var Track, _ = bindcap.New[*TrackArguments](TrackCommand)
