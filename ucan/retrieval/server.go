package retrieval

import (
	"github.com/alanshaw/ucantone/principal"
	"github.com/alanshaw/ucantone/server"
)

func NewServer(id principal.Signer, options ...server.HTTPOption) *server.HTTPServer {
	options = append(options, server.WithHTTPCodec(DefaultHTTPHeaderInboundCodec))
	return server.NewHTTP(id, options...)
}
