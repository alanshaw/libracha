package testutil

import (
	"net/url"

	"github.com/alanshaw/ucantone/did"
	"github.com/alanshaw/ucantone/principal/ed25519"
	"github.com/alanshaw/ucantone/principal/signer"
)

// did:key:z6Mkk89bC3JrVqKie71YEcc5M1SMVxuCgNx6zLZ8SYJsxALi
var Alice, _ = ed25519.Parse("MgCZT5vOnYZoVAeyjnzuJIVY9J4LNtJ+f8Js0cTPuKUpFnQ==")

// did:key:z6MkffDZCkCTWreg8868fG1FGFogcJj5X6PY93pPcWDn9bob
var Bob, _ = ed25519.Parse("MgCYbj5AJfVvdrjkjNCxB3iAUwx7RQHVQ7H1sKyHy46IosQ==")

// did:key:z6MkwYkD48SUrPhQ5Sf8qk5L8FW2L32Ze4guLnZXY4DrDCAR
var Carol, _ = ed25519.Parse("MgCa5pEVgZbqGILBFD3/TAd1a1OOJMuPsVz/uxS9ceU5jeg==")

// did:key:z6MktafZTREjJkvV5mfJxcLpNBoVPwDLhTuMg9ng7dY4zMAL
var Mallory, _ = ed25519.Parse("MgCYtH0AvYxiQwBG6+ZXcwlXywq9tI50G2mCAUJbwrrahkA==")

// did:key:z6Mkk3mDiu74xxyYEff5X1p568fVqEMczj5keYPT8qVMNsVC
var Service, _ = ed25519.Parse("MgCZyxtpD6SFBcXCXUKPTkLrc2+RlmaBjL/tMgWCT3+MUlw==")

var webServiceDID, _ = did.Parse("did:web:test.storacha.network")

// did:web:test.storacha.network
var WebService, _ = signer.Wrap(Service, webServiceDID)

var TestURL, _ = url.Parse("https://test.storacha.network")
