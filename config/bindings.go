package config

type bindings []string

// WithBinding configures ingress for an endpoint
//
// The requestedBindings argument is a string of the type of ingress for the endpoint
func WithBindings(requestedBindings ...string) interface {
	HTTPEndpointOption
	TLSEndpointOption
	TCPEndpointOption
} {
	ret := bindings{}
	for _, binding := range requestedBindings {
		ret = append(ret, binding)
	}
	return ret
}

func (b bindings) ApplyTLS(opts *tlsOptions) {
	opts.Bindings = []string(b)
}

func (b bindings) ApplyTCP(opts *tcpOptions) {
	opts.Bindings = []string(b)
}

func (b bindings) ApplyHTTP(opts *httpOptions) {
	opts.Bindings = []string(b)
}
