package attest

import (
	cdm "github.com/fil-forge/libforge/capabilities/datamodel"
	adm "github.com/fil-forge/libforge/capabilities/ucan/attest/datamodel"
	"github.com/fil-forge/ucantone/validator/bindcap"
)

type (
	ProofArguments = adm.ProofArgumentsModel
	ProofOK        = cdm.UnitModel
)

const ProofCommand = "/ucan/attest/proof"

// Issued by a trusted authority (usually the one handling invocation) that
// attests a specific UCAN delegation has been considered authentic.
var Proof, _ = bindcap.New[*ProofArguments](ProofCommand)
