package ucan

import (
	cdm "github.com/alanshaw/libracha/capabilities/datamodel"
	udm "github.com/alanshaw/libracha/capabilities/ucan/datamodel"
	"github.com/alanshaw/ucantone/validator/bindcap"
)

type (
	ConcludeArguments = udm.ConcludeArgumentsModel
	ConcludeOK        = cdm.UnitModel
)

const ConcludeCommand = "/ucan/conclude"

var Conclude, _ = bindcap.New[*ConcludeArguments](ConcludeCommand)
