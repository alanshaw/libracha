package commands

import (
	"github.com/fil-forge/ucantone/bind"
	"github.com/fil-forge/ucantone/ucan"
)

// MustParse is like [bind.Parse] but panics if the command cannot be
// constructed. It exists for package-level command declarations where the
// command and options are static — any error indicates a programming bug
// (malformed command string, invalid option) and should fail loudly at init
// rather than be silently dropped with `, _`.
func MustParse[Args, OK ucan.CBORValue](cmd string) bind.Binding[Args, OK] {
	c, err := bind.Parse[Args, OK](cmd)
	if err != nil {
		panic(err)
	}
	return c
}
