package config

type bindings []string

// WithBinding configures ingress for an endpoint
//
// The requestedBindings argument specifies the type of ingress for the endpoint.
func WithBindings(requestedBindings ...string) interface {
	HTTPEndpointOption
	TLSEndpointOption
	TCPEndpointOption
	LabeledTunnelOption
} {
	ret := bindings{}
	for _, binding := range requestedBindings {
		ret = append(ret, binding)
	}
	return ret
}

func (b bindings) ApplyTLS(cfg *tlsOptions) {
	cfg.Bindings = []string(b)
}

func (b bindings) ApplyTCP(cfg *tcpOptions) {
	cfg.Bindings = []string(b)
}

func (b bindings) ApplyHTTP(cfg *httpOptions) {
	cfg.Bindings = []string(b)
}

func (b bindings) ApplyLabeled(cfg *labeledOptions) {
	cfg.Bindings = []string(b)
}
