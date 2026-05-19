//go:build !codegen

package upload

import "github.com/fil-forge/libforge/commands"

const ListCommand = "/upload/list"

var List = commands.MustParse[*ListArguments](ListCommand)
