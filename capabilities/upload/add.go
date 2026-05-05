package upload

import (
	cdm "github.com/fil-forge/libforge/capabilities/datamodel"
	udm "github.com/fil-forge/libforge/capabilities/upload/datamodel"
	"github.com/fil-forge/ucantone/validator/bindcap"
)

const AddCommand = "/upload/add"

type (
	AddArguments = udm.AddArgumentsModel
	AddOK        = cdm.UnitModel
)

var Add, _ = bindcap.New[*AddArguments](AddCommand)
