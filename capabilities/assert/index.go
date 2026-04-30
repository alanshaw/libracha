package assert

import (
	adm "github.com/alanshaw/libracha/capabilities/assert/datamodel"
	cdm "github.com/alanshaw/libracha/capabilities/datamodel"
	"github.com/alanshaw/ucantone/validator/bindcap"
)

type (
	IndexArguments = adm.IndexArgumentsModel
	IndexOK        = cdm.UnitModel
)

const IndexCommand = "/assert/index"

// Index claims that a content graph can be found in blob(s) that are identified
// and indexed in the given index CID.
var Index, _ = bindcap.New[*IndexArguments](IndexCommand)
