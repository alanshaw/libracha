package upload

import (
	cdm "github.com/alanshaw/libracha/capabilities/datamodel"
	udm "github.com/alanshaw/libracha/capabilities/upload/datamodel"
	"github.com/alanshaw/ucantone/validator/bindcap"
)

const AddCommand = "/upload/add"

type (
	AddArguments = udm.AddArgumentsModel
	AddOK        = cdm.UnitModel
)

var Add, _ = bindcap.New[*AddArguments](AddCommand)
