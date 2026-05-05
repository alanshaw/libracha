package blob

import (
	bdm "github.com/fil-forge/libforge/capabilities/blob/datamodel"
	cdm "github.com/fil-forge/libforge/capabilities/datamodel"
	"github.com/fil-forge/ucantone/validator/bindcap"
)

const RemoveCommand = "/blob/remove"

type (
	RemoveArguments = bdm.RemoveArgumentsModel
	RemoveOK        = cdm.UnitModel
)

var Remove, _ = bindcap.New[*RemoveArguments](RemoveCommand)
