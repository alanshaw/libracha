package retrieval

import (
	"github.com/fil-forge/ucantone/principal"
	"github.com/fil-forge/ucantone/server"
)

func NewServer(id principal.Signer, options ...server.HTTPOption) *server.HTTPServer {
	options = append(options, server.WithHTTPCodec(DefaultHTTPHeaderInboundCodec))
	return server.NewHTTP(id, options...)
}
