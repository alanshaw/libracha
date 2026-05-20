//go:build !codegen

package upload

import "github.com/fil-forge/libforge/commands"

type AddOK = commands.Unit

var Add = commands.MustParse[*AddArguments, *AddOK]("/upload/add")
