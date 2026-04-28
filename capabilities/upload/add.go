package upload

import (
	dm "github.com/alanshaw/libracha/capabilities/upload/datamodel"
	"github.com/alanshaw/ucantone/validator/bindcap"
)

const AddCommand = "/upload/add"

type (
	AddArguments = dm.AddArgumentsModel
	AddOK        = dm.AddOKModel
)

var Add, _ = bindcap.New[*AddArguments](AddCommand)
