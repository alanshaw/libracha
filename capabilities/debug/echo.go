package debug

import (
	ddm "github.com/fil-forge/libforge/capabilities/debug/datamodel"
	"github.com/fil-forge/ucantone/ucan/delegation/policy"
	"github.com/fil-forge/ucantone/validator/bindcap"
	"github.com/fil-forge/ucantone/validator/capability"
)

const EchoCommand = "/debug/echo"

type (
	EchoArguments = ddm.EchoArgumentsModel
	EchoOK        = EchoArguments
)

var Echo, _ = bindcap.New[*EchoArguments](
	EchoCommand,
	capability.WithPolicyBuilder(policy.NotEqual(".message", "")),
)
