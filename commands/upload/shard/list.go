//go:build !codegen

package shard

import "github.com/fil-forge/libforge/commands"

const ListCommand = "/upload/shard/list"

var List = commands.MustParse[*ListArguments](ListCommand)
