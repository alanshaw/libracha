package claim

import (
	"github.com/alanshaw/libracha/capabilities/claim/datamodel"
	cdm "github.com/alanshaw/libracha/capabilities/datamodel"
	"github.com/alanshaw/ucantone/validator/bindcap"
)

type (
	CacheArguments = datamodel.CacheArgumentsModel
	CacheOK        = cdm.UnitModel
	Provider       = datamodel.ProviderModel
)

const CacheCommand = "/claim/cache"

var Cache, _ = bindcap.New[*CacheArguments](CacheCommand)
