package ucan

import (
	"github.com/alanshaw/libracha/capabilities/ucan/datamodel"
	"github.com/alanshaw/ucantone/validator/bindcap"
)

type ConcludeArguments = datamodel.ConcludeArgumentsModel

const ConcludeCommand = "/ucan/conclude"

var Conclude, _ = bindcap.New[*ConcludeArguments](ConcludeCommand)
