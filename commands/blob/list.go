//go:build !codegen

package blob

import "github.com/fil-forge/libforge/commands"

const ListCommand = "/blob/list"

var List = commands.MustParse[*ListArguments](ListCommand)
