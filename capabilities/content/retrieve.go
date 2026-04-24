package content

import (
	dm "github.com/alanshaw/libracha/capabilities/content/datamodel"
	"github.com/alanshaw/ucantone/validator/bindcap"
)

const RetrieveCommand = "/content/retrieve"

type (
	RetrieveArguments = dm.RetrieveArgumentsModel
	Blob              = dm.BlobModel
	RetrieveOK        = dm.RetrieveOKModel
)

var Retrieve, _ = bindcap.New[*RetrieveArguments](RetrieveCommand)
