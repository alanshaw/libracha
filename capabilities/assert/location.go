package assert

import (
	adm "github.com/fil-forge/libforge/capabilities/assert/datamodel"
	cdm "github.com/fil-forge/libforge/capabilities/datamodel"
	"github.com/fil-forge/ucantone/validator/bindcap"
)

type (
	LocationArguments = adm.LocationArgumentsModel
	LocationOK        = cdm.UnitModel
	Range             = adm.RangeModel
)

const LocationCommand = "/assert/location"

var Location, _ = bindcap.New[*LocationArguments](LocationCommand)
