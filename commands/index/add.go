//go:build !codegen

package index

import (
	"github.com/fil-forge/libforge/commands"
	"github.com/fil-forge/ucantone/errors"
)

const AddCommand = "/index/add"

type AddOK = commands.Unit

var Add = commands.MustParse[*AddArguments](AddCommand)

const IndexNotFoundErrorName = "IndexNotFound"

var ErrIndexNotFound = errors.New(IndexNotFoundErrorName, "index not found in space")
