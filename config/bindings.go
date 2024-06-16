package config

import "fmt"

type bindings []string

// WithBinding configures bindings for an endpoint
//
// The requestedBindings argument is a
func WithBindings(requestedBindings ...string) interface {
	HTTPEndpointOption
	TLSEndpointOption
	TCPEndpointOption
} {
	fmt.Printf("\n\nDebug Log: what is requestedBindings %v\n\n", requestedBindings)
	ret := make(bindings, len(requestedBindings))
	for _, binding := range requestedBindings {
		ret = append(ret, binding)
	}
	return ret
}

func (b bindings) ApplyTLS(opts *tlsOptions) {
	fmt.Printf("\n\nDebug Log TLS: what are bindings %v\n\n", b)
	opts.Bindings = b
	fmt.Printf("\n\nDebug Log TLS: what are opts.Bindings %v\n\n", opts.Bindings)
}

func (b bindings) ApplyTCP(opts *tcpOptions) {
	fmt.Printf("\n\nDebug Log TCP : what are bindings %v\n\n", b)
	opts.Bindings = b
	fmt.Printf("\n\nDebug Log TCP: what are opts.Bindings %v\n\n", opts.Bindings)
}

func (b bindings) ApplyHTTP(opts *httpOptions) {
	fmt.Printf("\n\nDebug Log HTTP: what are bindings %v\n\n", b)
	opts.Bindings = b
	fmt.Printf("\n\nDebug Log HTTP: what are opts.Bindings %v\n\n", opts.Bindings)
}
