//go:build !codegen

package blob

import "github.com/fil-forge/libforge/commands"

const AcceptCommand = "/blob/accept"

var Accept = commands.MustParse[*AcceptArguments](AcceptCommand)
