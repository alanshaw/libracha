package upload

import (
	cdm "github.com/fil-forge/libforge/capabilities/datamodel"
	udm "github.com/fil-forge/libforge/capabilities/upload/datamodel"
	"github.com/fil-forge/ucantone/validator/bindcap"
)

const RemoveCommand = "/upload/remove"

type (
	RemoveArguments = udm.RemoveArgumentsModel
	RemoveOK        = cdm.UnitModel
)

var Remove, _ = bindcap.New[*RemoveArguments](RemoveCommand)
