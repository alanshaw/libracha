//go:build !codegen

package pdp

import "github.com/fil-forge/libforge/commands"

var Accept = commands.MustParse[*AcceptArguments]("/pdp/accept")
