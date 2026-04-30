package blob

import (
	bdm "github.com/alanshaw/libracha/capabilities/blob/datamodel"
	cdm "github.com/alanshaw/libracha/capabilities/datamodel"
	"github.com/alanshaw/ucantone/validator/bindcap"
)

const RemoveCommand = "/blob/remove"

type (
	RemoveArguments = bdm.RemoveArgumentsModel
	RemoveOK        = cdm.UnitModel
)

var Remove, _ = bindcap.New[*RemoveArguments](RemoveCommand)
