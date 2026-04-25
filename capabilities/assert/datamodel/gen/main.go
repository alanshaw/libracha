package main

import (
	jsg "github.com/alanshaw/dag-json-gen"
	"github.com/alanshaw/libracha/capabilities/assert/datamodel"
	cbg "github.com/whyrusleeping/cbor-gen"
)

func main() {
	models := []any{
		datamodel.IndexArgumentsModel{},
		datamodel.IndexOKModel{},
		datamodel.LocationArgumentsModel{},
		datamodel.LocationOKModel{},
		datamodel.RangeModel{},
	}

	if err := cbg.WriteMapEncodersToFile("../cbor_gen.go", "datamodel", models...); err != nil {
		panic(err)
	}

	if err := jsg.WriteMapEncodersToFile("../json_gen.go", "datamodel", models...); err != nil {
		panic(err)
	}
}
