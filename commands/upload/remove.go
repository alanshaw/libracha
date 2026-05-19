//go:build !codegen

package upload

import "github.com/fil-forge/libforge/commands"

const RemoveCommand = "/upload/remove"

type RemoveOK = commands.Unit

var Remove = commands.MustParse[*RemoveArguments](RemoveCommand)
