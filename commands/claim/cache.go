//go:build !codegen

package claim

import "github.com/fil-forge/libforge/commands"

const CacheCommand = "/claim/cache"

type CacheOK = commands.Unit

var Cache = commands.MustParse[*CacheArguments](CacheCommand)
