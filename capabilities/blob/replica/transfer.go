package replica

import (
	"github.com/fil-forge/libforge/capabilities/blob"
	rdm "github.com/fil-forge/libforge/capabilities/blob/replica/datamodel"
	"github.com/fil-forge/ucantone/ucan/delegation/policy"
	"github.com/fil-forge/ucantone/validator/bindcap"
	"github.com/fil-forge/ucantone/validator/capability"
)

type (
	TransferArguments = rdm.TransferArgumentsModel
	TransferOK        = rdm.TransferOKModel
)

const TransferCommand = "/blob/replica/transfer"

var Transfer, _ = bindcap.New[*TransferArguments](
	TransferCommand,
	capability.WithPolicyBuilder(
		policy.GreaterThan(".blob.size", 0),
		policy.LessThanOrEqual(".blob.size", blob.MaxBlobSize),
	),
)
