package replica

import (
	"github.com/alanshaw/libracha/capabilities/blob"
	rdm "github.com/alanshaw/libracha/capabilities/blob/replica/datamodel"
	"github.com/alanshaw/ucantone/ucan/delegation/policy"
	"github.com/alanshaw/ucantone/validator/bindcap"
	"github.com/alanshaw/ucantone/validator/capability"
)

type (
	AllocateArguments = rdm.AllocateArgumentsModel
	AllocateOK        = rdm.AllocateOKModel
	Blob              = rdm.BlobModel
)

const AllocateCommand = "/blob/replica/allocate"

var Allocate, _ = bindcap.New[*AllocateArguments](
	AllocateCommand,
	capability.WithPolicyBuilder(
		policy.GreaterThan(".blob.size", 0),
		policy.LessThanOrEqual(".blob.size", blob.MaxBlobSize),
	),
)
