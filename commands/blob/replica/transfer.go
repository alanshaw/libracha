//go:build !codegen

package replica

import "github.com/fil-forge/libforge/commands"

var Transfer = commands.MustParse[*TransferArguments, *TransferOK]("/blob/replica/transfer")
