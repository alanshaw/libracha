package blob

import (
	bdm "github.com/fil-forge/libforge/capabilities/blob/datamodel"
	"github.com/fil-forge/ucantone/ucan/delegation/policy"
	"github.com/fil-forge/ucantone/validator/bindcap"
	"github.com/fil-forge/ucantone/validator/capability"
)

const AddCommand = "/blob/add"

type (
	AddArguments = bdm.AddArgumentsModel
	Blob         = bdm.BlobModel
	AddOK        = bdm.AddOKModel
)

var Add, _ = bindcap.New[*AddArguments](
	AddCommand,
	capability.WithPolicyBuilder(
		policy.GreaterThan(".blob.size", 0),
		policy.LessThanOrEqual(".blob.size", MaxBlobSize),
	),
)
