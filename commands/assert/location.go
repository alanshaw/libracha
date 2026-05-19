//go:build !codegen

package assert

import "github.com/fil-forge/libforge/commands"

const LocationCommand = "/assert/location"

type LocationOK = commands.Unit

var Location = commands.MustParse[*LocationArguments](LocationCommand)
