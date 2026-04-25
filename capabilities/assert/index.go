package assert

import (
	"github.com/alanshaw/libracha/capabilities/assert/datamodel"
	"github.com/alanshaw/ucantone/validator/bindcap"
)

type (
	IndexArguments = datamodel.IndexArgumentsModel
	IndexOK        = datamodel.IndexOKModel
)

const IndexCommand = "/assert/index"

// Index claims that a content graph can be found in blob(s) that are identified
// and indexed in the given index CID.
var Index, _ = bindcap.New[*IndexArguments](IndexCommand)
