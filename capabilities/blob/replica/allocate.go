package replica

import (
	"github.com/fil-forge/libforge/capabilities/blob"
	rdm "github.com/fil-forge/libforge/capabilities/blob/replica/datamodel"
	"github.com/fil-forge/ucantone/ucan/delegation/policy"
	"github.com/fil-forge/ucantone/validator/bindcap"
	"github.com/fil-forge/ucantone/validator/capability"
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
