package claim

import (
	"github.com/alanshaw/libracha/capabilities/claim/datamodel"
	"github.com/alanshaw/ucantone/validator/bindcap"
)

type (
	CacheArguments = datamodel.CacheArgumentsModel
	CacheOK        = datamodel.CacheOKModel
	Provider       = datamodel.ProviderModel
)

const CacheCommand = "/claim/cache"

var Cache, _ = bindcap.New[*CacheArguments](CacheCommand)
