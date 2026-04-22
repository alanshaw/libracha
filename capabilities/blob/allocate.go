package blob

import (
	bdm "github.com/alanshaw/libracha/capabilities/blob/datamodel"
	"github.com/alanshaw/ucantone/ucan/delegation/policy"
	"github.com/alanshaw/ucantone/validator/bindcap"
	"github.com/alanshaw/ucantone/validator/capability"
)

const MaxBlobSize = 268_435_456

type (
	AllocateArguments = bdm.AllocateArgumentsModel
	AllocateOK        = bdm.AllocateOKModel
	BlobAddress       = bdm.BlobAddressModel
)

const AllocateCommand = "/blob/allocate"

var Allocate, _ = bindcap.New[*AllocateArguments](
	AllocateCommand,
	capability.WithPolicyBuilder(
		policy.GreaterThan(".blob.size", 0),
		policy.LessThanOrEqual(".blob.size", MaxBlobSize),
	),
)
