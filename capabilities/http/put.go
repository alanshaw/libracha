package http

import (
	cdm "github.com/fil-forge/libforge/capabilities/datamodel"
	hdm "github.com/fil-forge/libforge/capabilities/http/datamodel"
	"github.com/fil-forge/ucantone/validator/bindcap"
)

const PutCommand = "/http/put"

type (
	PutArguments = hdm.PutArgumentsModel
	PutOK        = cdm.UnitModel
)

var Put, _ = bindcap.New[*PutArguments](PutCommand)
