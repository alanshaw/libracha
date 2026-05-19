//go:build !codegen

package ucan

import (
	"github.com/fil-forge/libforge/commands"
	"github.com/fil-forge/ucantone/errors"
)

const ConcludeCommand = "/ucan/conclude"

type ConcludeOK = commands.Unit

var Conclude = commands.MustParse[*ConcludeArguments](ConcludeCommand)

const ConclusionReceiptNotFoundErrorName = "ConclusionReceiptNotFound"

var ErrConclusionReceiptNotFound = errors.New(ConclusionReceiptNotFoundErrorName, "conclusion receipt not found")
