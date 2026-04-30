package assert

import (
	adm "github.com/alanshaw/libracha/capabilities/assert/datamodel"
	cdm "github.com/alanshaw/libracha/capabilities/datamodel"
	"github.com/alanshaw/ucantone/validator/bindcap"
)

type (
	LocationArguments = adm.LocationArgumentsModel
	LocationOK        = cdm.UnitModel
	Range             = adm.RangeModel
)

const LocationCommand = "/assert/location"

var Location, _ = bindcap.New[*LocationArguments](LocationCommand)
