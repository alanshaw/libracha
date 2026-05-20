//go:build !codegen

package pdp

import "github.com/fil-forge/libforge/commands"

const InfoCommand = "/pdp/info"

var Info = commands.MustParse[*InfoArguments, *InfoOK]("/pdp/info")
