package index

import (
	cdm "github.com/fil-forge/libforge/capabilities/datamodel"
	dm "github.com/fil-forge/libforge/capabilities/index/datamodel"
	"github.com/fil-forge/ucantone/validator/bindcap"
)

const AddCommand = "/index/add"

type (
	AddArguments = dm.AddArgumentsModel
	AddOK        = cdm.UnitModel
)

var Add, _ = bindcap.New[*AddArguments](AddCommand)
