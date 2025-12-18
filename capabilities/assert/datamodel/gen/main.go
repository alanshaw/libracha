package main

import (
	"github.com/alanshaw/libracha/capabilities/assert/datamodel"
	cbg "github.com/whyrusleeping/cbor-gen"
)

func main() {
	if err := cbg.WriteMapEncodersToFile("../cbor_gen.go", "datamodel",
		datamodel.LocationArgumentsModel{},
		datamodel.RangeModel{},
	); err != nil {
		panic(err)
	}
}
