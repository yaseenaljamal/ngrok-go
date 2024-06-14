package config

type bindings []string

// WithBinding configures bindings for an endpoint
//
// The requestedBindings argument is a
func WithBindings(requestedBindings ...string) interface {
	HTTPEndpointOption
	TLSEndpointOption
	TCPEndpointOption
} {
	return (*bindings)(&requestedBindings)
}

func (b *bindings) ApplyTLS(opts *tlsOptions) {
	opts.Bindings = b
}

func (b *bindings) ApplyTCP(opts *tcpOptions) {
	opts.Bindings = b
}

func (b *bindings) ApplyHTTP(opts *httpOptions) {
	opts.Bindings = b
}
