package main

import (
	ddm "github.com/alanshaw/libracha/capabilities/debug/datamodel"
	cbg "github.com/whyrusleeping/cbor-gen"
)

func main() {
	if err := cbg.WriteMapEncodersToFile("../cbor_gen.go", "datamodel",
		ddm.EchoArgumentsModel{},
	); err != nil {
		panic(err)
	}
}
