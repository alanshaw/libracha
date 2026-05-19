//go:build !codegen

package http

import "github.com/fil-forge/libforge/commands"

const PutCommand = "/http/put"

type PutOK = commands.Unit

var Put = commands.MustParse[*PutArguments](PutCommand)
