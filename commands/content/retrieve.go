//go:build !codegen

package content

import "github.com/fil-forge/libforge/commands"

const RetrieveCommand = "/content/retrieve"

type RetrieveOK = commands.Unit

var Retrieve = commands.MustParse[*RetrieveArguments](RetrieveCommand)
