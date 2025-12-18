package blob

import (
	bdm "github.com/alanshaw/libracha/capabilities/blob/datamodel"
	"github.com/alanshaw/ucantone/validator/bindcap"
)

type (
	AcceptArguments = bdm.AcceptArgumentsModel
	AcceptOK        = bdm.AcceptOKModel
)

const AcceptCommand = "/blob/accept"

var Accept, _ = bindcap.New[*AcceptArguments](AcceptCommand)
