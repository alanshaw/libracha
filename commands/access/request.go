//go:build !codegen

package access

import "github.com/fil-forge/libforge/commands"

// RequestFactKey is the key in metadata in any delegation created by a
// successful access request. The value is a link back to the `/access/request`
// invocation.
const RequestMetaKey = "accessRequest"

const RequestCommand = "/access/request"

// Request can be invoked by an agent to request set of capabilities from the
// account.
var Request = commands.MustParse[*RequestArguments](RequestCommand)

const (
	InvalidAuthorizationAccountErrorName  = "InvalidAuthorizationAccount"
	InvalidAuthorizationAudienceErrorName = "InvalidAuthorizationAudience"
)
