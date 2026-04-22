package retrieval

import (
	"net/http"
	"net/url"

	"github.com/alanshaw/ucantone/client"
	"github.com/alanshaw/ucantone/execution"
)

type HTTPHeaderClient struct {
	*client.Client[*http.Request, *http.Response]
}

type httpHeaderClientConfig struct {
	client    *http.Client
	listeners []client.EventListener
}

type ClientOption func(*httpHeaderClientConfig)

func WithHTTPClient(client *http.Client) ClientOption {
	return func(cfg *httpHeaderClientConfig) {
		cfg.client = client
	}
}

// WithEventListener adds an event listener to the HTTP client for monitoring
// requests and responses.
func WithEventListener(listener client.EventListener) ClientOption {
	return func(cfg *httpHeaderClientConfig) {
		cfg.listeners = append(cfg.listeners, listener)
	}
}

func NewClient(serviceURL *url.URL, options ...ClientOption) (*HTTPHeaderClient, error) {
	codec := DefaultHTTPHeaderOutboundCodec
	cfg := httpHeaderClientConfig{
		client: http.DefaultClient,
	}
	for _, opt := range options {
		opt(&cfg)
	}
	c := client.New(&httpTransport{cfg.client, serviceURL}, codec)
	c.Listeners = cfg.listeners
	return &HTTPHeaderClient{Client: c}, nil
}

// Note: execution response metadata is a [HTTPHeaderResponseContainer] and it
// is the caller's responsibility to close the response body.
func (c *HTTPHeaderClient) Execute(execRequest execution.Request) (execution.Response, error) {
	return c.Client.Execute(execRequest)
}

type httpTransport struct {
	client *http.Client
	url    *url.URL
}

func (t *httpTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.URL = t.url
	return t.client.Do(r)
}
