//go:build !codegen

package space

import "github.com/fil-forge/libforge/commands"

const InfoCommand = "/space/info"

type InfoArguments = commands.Unit

var Info = commands.MustParse[*InfoArguments](InfoCommand)

const UnknownSpaceErrorName = "UnknownSpace"
