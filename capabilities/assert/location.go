package assert

import (
	"github.com/alanshaw/libracha/capabilities/assert/datamodel"
	"github.com/alanshaw/ucantone/validator/bindcap"
)

type (
	LocationArguments = datamodel.LocationArgumentsModel
	LocationOK        = datamodel.LocationOKModel
	Range             = datamodel.RangeModel
)

const LocationCommand = "/assert/location"

var Location, _ = bindcap.New[*LocationArguments](LocationCommand)
