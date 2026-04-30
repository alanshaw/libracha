package upload

import (
	cdm "github.com/alanshaw/libracha/capabilities/datamodel"
	udm "github.com/alanshaw/libracha/capabilities/upload/datamodel"
	"github.com/alanshaw/ucantone/validator/bindcap"
)

const RemoveCommand = "/upload/remove"

type (
	RemoveArguments = udm.RemoveArgumentsModel
	RemoveOK        = cdm.UnitModel
)

var Remove, _ = bindcap.New[*RemoveArguments](RemoveCommand)
