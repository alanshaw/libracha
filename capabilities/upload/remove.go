package upload

import (
	dm "github.com/alanshaw/libracha/capabilities/upload/datamodel"
	"github.com/alanshaw/ucantone/validator/bindcap"
)

const RemoveCommand = "/upload/remove"

type (
	RemoveArguments = dm.RemoveArgumentsModel
	RemoveOK        = dm.RemoveOKModel
)

var Remove, _ = bindcap.New[*RemoveArguments](RemoveCommand)
