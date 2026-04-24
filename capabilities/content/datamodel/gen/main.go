package main

import (
	jsg "github.com/alanshaw/dag-json-gen"
	dm "github.com/alanshaw/libracha/capabilities/content/datamodel"
	cbg "github.com/whyrusleeping/cbor-gen"
)

func main() {
	if err := cbg.WriteTupleEncodersToFile("../cbor_gen.tuples.go", "datamodel",
		dm.RangeModel{},
	); err != nil {
		panic(err)
	}

	if err := cbg.WriteMapEncodersToFile("../cbor_gen.maps.go", "datamodel",
		dm.BlobModel{},
		dm.RetrieveArgumentsModel{},
		dm.RetrieveOKModel{},
	); err != nil {
		panic(err)
	}

	if err := jsg.WriteTupleEncodersToFile("../json_gen.tuples.go", "datamodel",
		dm.RangeModel{},
	); err != nil {
		panic(err)
	}

	if err := jsg.WriteMapEncodersToFile("../json_gen.maps.go", "datamodel",
		dm.BlobModel{},
		dm.RetrieveArgumentsModel{},
		dm.RetrieveOKModel{},
	); err != nil {
		panic(err)
	}
}
