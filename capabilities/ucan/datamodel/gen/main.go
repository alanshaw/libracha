package main

import (
	udm "github.com/alanshaw/libracha/capabilities/ucan/datamodel"
	cbg "github.com/whyrusleeping/cbor-gen"
)

func main() {
	if err := cbg.WriteMapEncodersToFile("../cbor_gen.go", "datamodel",
		udm.ConcludeArgumentsModel{},
	); err != nil {
		panic(err)
	}
}
