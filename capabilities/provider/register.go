package provider

import (
	cdm "github.com/alanshaw/libracha/capabilities/datamodel"
	pdm "github.com/alanshaw/libracha/capabilities/provider/datamodel"
	"github.com/alanshaw/ucantone/validator/bindcap"
)

const RegisterCommand = "/provider/register"

type (
	RegisterArguments = pdm.RegisterArgumentsModel
	RegisterOK        = cdm.UnitModel
)

var Register, _ = bindcap.New[*RegisterArguments](RegisterCommand)
