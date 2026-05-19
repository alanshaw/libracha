//go:build !codegen

package upload

import "github.com/fil-forge/libforge/commands"

const AddCommand = "/upload/add"

type AddOK = commands.Unit

var Add = commands.MustParse[*AddArguments](AddCommand)
