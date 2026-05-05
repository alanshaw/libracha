package blob

import (
	bdm "github.com/fil-forge/libforge/capabilities/blob/datamodel"
	"github.com/fil-forge/ucantone/validator/bindcap"
)

type (
	AcceptArguments = bdm.AcceptArgumentsModel
	AcceptOK        = bdm.AcceptOKModel
)

const AcceptCommand = "/blob/accept"

var Accept, _ = bindcap.New[*AcceptArguments](AcceptCommand)
