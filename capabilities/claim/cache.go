package claim

import (
	"github.com/fil-forge/libforge/capabilities/claim/datamodel"
	cdm "github.com/fil-forge/libforge/capabilities/datamodel"
	"github.com/fil-forge/ucantone/validator/bindcap"
)

type (
	CacheArguments = datamodel.CacheArgumentsModel
	CacheOK        = cdm.UnitModel
	Provider       = datamodel.ProviderModel
)

const CacheCommand = "/claim/cache"

var Cache, _ = bindcap.New[*CacheArguments](CacheCommand)
