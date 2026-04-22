package replica

import (
	"github.com/alanshaw/libracha/capabilities/blob"
	rdm "github.com/alanshaw/libracha/capabilities/blob/replica/datamodel"
	"github.com/alanshaw/ucantone/ucan/delegation/policy"
	"github.com/alanshaw/ucantone/validator/bindcap"
	"github.com/alanshaw/ucantone/validator/capability"
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
