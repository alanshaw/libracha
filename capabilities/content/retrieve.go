package content

import (
	dm "github.com/fil-forge/libforge/capabilities/content/datamodel"
	cdm "github.com/fil-forge/libforge/capabilities/datamodel"
	"github.com/fil-forge/ucantone/validator/bindcap"
)

const RetrieveCommand = "/content/retrieve"

type (
	RetrieveArguments = dm.RetrieveArgumentsModel
	Blob              = dm.BlobModel
	Range             = dm.RangeModel
	RetrieveOK        = cdm.UnitModel
)

var Retrieve, _ = bindcap.New[*RetrieveArguments](RetrieveCommand)
