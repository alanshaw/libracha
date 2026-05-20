//go:build !codegen

package space

import "github.com/fil-forge/libforge/commands"

type InfoArguments = commands.Unit

var Info = commands.MustParse[*InfoArguments, *InfoOK]("/space/info")

const UnknownSpaceErrorName = "UnknownSpace"
