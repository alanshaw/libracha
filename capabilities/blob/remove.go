package blob

import (
	bdm "github.com/alanshaw/libracha/capabilities/blob/datamodel"
	"github.com/alanshaw/ucantone/validator/bindcap"
)

const RemoveCommand = "/blob/remove"

type (
	RemoveArguments = bdm.RemoveArgumentsModel
	RemoveOK        = bdm.RemoveOKModel
)

var Remove, _ = bindcap.New[*RemoveArguments](RemoveCommand)
